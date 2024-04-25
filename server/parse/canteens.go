package parse

import (
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ref-it/infoscreen/server/common"
)

func ParseMenu(canteenID int, menuDate string, additivesMap map[string]string, allergensMap map[string]string) ([]common.Meal, []common.Meal) {
	menuUrl := "https://www.stw-thueringen.de/xhr/loadspeiseplan.html"

	form := url.Values{}
	form.Add("resources_id", strconv.Itoa(canteenID))
	form.Add("date", menuDate)

	request, err := http.NewRequest("POST", menuUrl, strings.NewReader(form.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	if err != nil {
		common.ErrorLogger.Println("Unable to form menu request.")
		return []common.Meal{}, []common.Meal{}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		common.ErrorLogger.Println("Unable to retrieve menu data.")
		return []common.Meal{}, []common.Meal{}
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		common.ErrorLogger.Println("Unable to read menu response body.")
		return []common.Meal{}, []common.Meal{}
	}

	var mealsLunch []common.Meal
	var mealsDinner []common.Meal

	doc.Find(".splGroupWrapper").Each(func(i int, s *goquery.Selection) {
		s.Find(".rowMealInner").Each(func(j int, t *goquery.Selection) {
			mealName := strings.TrimSpace(t.Find(".mealText").Text())
			mealNameRegEx := regexp.MustCompile(`(mensaInternational|MensaInternational).*\:`)
			mealName = mealNameRegEx.ReplaceAllString(mealName, "")
			mealName = strings.TrimPrefix(mealName, "mensaVital:")
			mealName = strings.TrimPrefix(mealName, "MensaVital:")
			mealName = strings.TrimPrefix(mealName, "Burgertag")
			mealName = strings.Replace(mealName, "!", "", -1)
			mealName = strings.Replace(mealName, " ,", ",", -1)
			mealName = strings.TrimSpace(mealName)

			mealAdditives := strings.Split(strings.TrimPrefix(strings.TrimSpace(t.Find(".zusatzstoffe").Text()), "Zusatzstoffe: "), ",")
			if mealAdditives[0] == "" {
				mealAdditives = []string{}
			}
			for i := 0; i < len(mealAdditives); i++ {
				mealAdditives[i] = additivesMap[mealAdditives[i]]
			}

			mealAllergens := strings.Split(strings.TrimPrefix(strings.TrimSpace(t.Find(".allergene").Text()), "Allergene: "), ",")
			if mealAllergens[0] == "" {
				mealAllergens = []string{}
			}
			for i := 0; i < len(mealAllergens); i++ {
				mealAllergens[i] = allergensMap[mealAllergens[i]]
			}

			pricesString := strings.TrimSpace(t.Find(".mealPreise").Text())
			pricesSlice := strings.Split(pricesString, " / ")
			priceStudents, _ := strconv.ParseFloat(strings.Replace(pricesSlice[0], ",", ".", -1), 32)
			priceEmployees, _ := strconv.ParseFloat(strings.Replace(pricesSlice[1], ",", ".", -1), 32)
			priceGuests, _ := strconv.ParseFloat(strings.Replace(strings.TrimSuffix(pricesSlice[2], " €"), ",", ".", -1), 32)
			mealPrices := common.Prices{Students: float32(priceStudents), Employees: float32(priceEmployees), Guests: float32(priceGuests)}

			isVegetarian := false
			isVegan := false

			t.Find(".splIconMeal").Each(func(k int, u *goquery.Selection) {
				altText, _ := u.Attr("alt")
				if altText == "Vegetarische Speisen (V)" {
					isVegetarian = true
				}
				if altText == "Vegane Speisen (V*)" {
					isVegan = true
				}
			})

			newMeal := common.Meal{Name: mealName, Additives: mealAdditives, Allergens: mealAllergens, Prices: mealPrices, IsVegetarian: isVegetarian, IsVegan: isVegan}
			if i == 0 {
				mealsLunch = append(mealsLunch, newMeal)
			} else {
				mealsDinner = append(mealsLunch, newMeal)
			}
		})
	})

	return mealsLunch, mealsDinner
}
