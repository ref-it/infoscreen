package parse

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/ref-it/infoscreen/server/common"
)

func ParseTrainDepartures(xmlResponse *http.Response) []common.TrainDeparture {
	var trainDeparturesList []common.TrainDeparture

	/*bodyBytes, err := io.ReadAll(xmlResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)*/

	doc, err := goquery.NewDocumentFromReader(xmlResponse.Body)
	if err != nil {
		common.ErrorLogger.Println("Unable to read train departure response body.")
		return []common.TrainDeparture{}
	}

	doc.Find("s").Each(func(i int, s *goquery.Selection) {
		dp := s.Find("dp")
		if len(dp.Nodes) != 0 {
			tl := s.Find("tl")
			category, _ := tl.Attr("c")
			line, _ := dp.Attr("l")
			timestamp, _ := dp.Attr("pt")
			minutes := timestamp[8:10]
			hours := timestamp[6:8]
			routeString, _ := dp.Attr("ppth")
			route := strings.Split(routeString, "|")
			destination := route[len(route)-1]
			track, _ := dp.Attr("pp")
			trainDeparture := common.TrainDeparture{Category: category, Line: line, Hours: hours, Minutes: minutes, Destination: destination, Track: track}
			trainDeparturesList = append(trainDeparturesList, trainDeparture)
		}
	})

	sort.Slice(trainDeparturesList, func(i, j int) bool {
		return trainDeparturesList[i].Minutes < trainDeparturesList[j].Minutes
	})

	return trainDeparturesList
}

func GetTrainDepartures(stationId int, dbClientId string, dbApiKey string) []common.TrainDeparture {
	var trainDeparturesList []common.TrainDeparture
	baseUrl := "https://apis.deutschebahn.com/db-api-marketplace/apis/timetables/v1/plan/" + strconv.Itoa(stationId)

	for i := 0; i < 4; i++ {
		// Generate the URL for the API request
		trainDay := time.Now().Add(time.Hour * time.Duration(i)).Format("02")
		trainMonth := time.Now().Add(time.Hour * time.Duration(i)).Format("01")
		trainYear := time.Now().Add(time.Hour * time.Duration(i)).Format("06")
		trainHour := time.Now().Add(time.Hour * time.Duration(i)).Format("15")
		url := baseUrl + "/" + trainYear + trainMonth + trainDay + "/" + trainHour

		reqest, _ := http.NewRequest("GET", url, nil)

		// Set the headers necessary for the API request (Client ID and API Key)
		reqest.Header.Add("DB-Client-Id", dbClientId)
		reqest.Header.Add("DB-Api-Key", dbApiKey)
		reqest.Header.Add("accept", "application/xml")

		response, _ := http.DefaultClient.Do(reqest)

		trainDeparturesList = append(trainDeparturesList, ParseTrainDepartures(response)...)
	}

	return trainDeparturesList
}
