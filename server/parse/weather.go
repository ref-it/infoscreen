package parse

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ref-it/infoscreen/server/common"
)

func ParseWeatherForecast(jsonResponse *http.Response) common.WeatherForecast {
	var weatherForecast common.WeatherForecast

	body, err := io.ReadAll(jsonResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &weatherForecast)
	if err != nil {
		panic(err)
	}

	return weatherForecast
}

func GetWeatherStationForecast(stationId string) common.WeatherForecast {
	var weatherForecast common.WeatherForecast
	url := "https://s3.eu-central-1.amazonaws.com/app-prod-static.warnwetter.de/v16/forecast_mosmix_" + stationId + ".json"
	reqest, _ := http.NewRequest("GET", url, nil)
	response, _ := http.DefaultClient.Do(reqest)
	weatherForecast = ParseWeatherForecast(response)

	return weatherForecast
}
