package retrieve

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/ref-it/infoscreen/server/common"
	"github.com/ref-it/infoscreen/server/parse"
)

func getAndSaveMenus() {
	// Get the current time as string of four digits and convert it to int
	currentTime, err := strconv.Atoi(time.Now().Format("1504"))
	if err != nil {
		common.ErrorLogger.Println("Unable to convert string to int.")
	}

	// Check if the current time is in the time range of interest
	if currentTime > 500 && currentTime < 2000 {
		common.InfoLogger.Println("Retrieving menu data")

		// Open the config file with a list of all canteen names, ID numbers and maps of abbrevations for additives and allergens
		canteenFilePath := "conf/canteens.json"
		canteenDataFile, err := os.Open(canteenFilePath)
		if err != nil {
			common.ErrorLogger.Printf("Unable to open file %s.", canteenFilePath)
		}
		canteenDataFileBytes, err := io.ReadAll(canteenDataFile)
		if err != nil {
			common.ErrorLogger.Printf("Unable to read the data in file %s.", canteenFilePath)
		}

		// Unmarshal the JSON data
		var canteenData common.CanteenData
		err = json.Unmarshal(canteenDataFileBytes, &canteenData)
		if err != nil {
			common.ErrorLogger.Printf("Unable to unmarshal the content of %s.", canteenFilePath)
		}

		// Iterate over all Canteens
		for i := 0; i < len(canteenData.Canteens); i++ {
			// Check if the current time is in the time range of interest for the specific Canteen
			if currentTime < 1400 || currentTime < 1430 && canteenData.Canteens[i].ID == 46 || canteenData.Canteens[i].ID == 53 {
				// Retrieve and parse the menu data
				mealsLunch, mealsDinner := parse.ParseMenu(canteenData.Canteens[i].ID, time.Now().Format("02.01.2006"), canteenData.Additives, canteenData.Allergens)

				// Marshal the menu data
				meals := common.Menu{CanteenName: canteenData.Canteens[i].Name, Lunch: mealsLunch, Dinner: mealsDinner, LastUpdated: time.Now().Unix()}
				data, err := json.Marshal(meals)
				if err != nil {
					common.ErrorLogger.Printf("Unable to marshal meals (Canteen ID: %s)", strconv.Itoa(canteenData.Canteens[i].ID))
				}

				// Write the menu data into a file
				dataFilePath := "data/menu/" + strconv.Itoa(canteenData.Canteens[i].ID) + ".json"
				err = os.WriteFile(dataFilePath, data, 0644)
				if err != nil {
					common.ErrorLogger.Printf("Unable to write canteen data to file %s", dataFilePath)
				}
			}
		}
	}
}

func GetMenu() {
	// Retrieve data every 15 minutes
	getAndSaveMenus()
	ticker := time.NewTicker(time.Minute * 15)
	for range ticker.C {
		getAndSaveMenus()
	}
}
