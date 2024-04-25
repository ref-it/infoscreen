package retrieve

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/ref-it/infoscreen/server/common"
	"github.com/ref-it/infoscreen/server/parse"
)

func getAndSaveWeatherForecast() {
	common.InfoLogger.Println("Retrieving weather forecast data")

	// Open the config file with a list of all station names and ID numbers
	stationFilePath := "conf/stationsWeather.json"
	stationDataFile, err := os.Open(stationFilePath)
	if err != nil {
		common.ErrorLogger.Printf("Unable to open file %s.", stationFilePath)
	}
	stationDataFileBytes, err := io.ReadAll(stationDataFile)
	if err != nil {
		common.ErrorLogger.Printf("Unable to read the data in file %s.", stationFilePath)
	}

	// Unmarshal the JSON data
	var stations []common.WeatherStationData
	err = json.Unmarshal(stationDataFileBytes, &stations)
	if err != nil {
		common.ErrorLogger.Printf("Unable to unmarshal the content of %s.", stationFilePath)
	}

	// Open the configuration file with API credentials
	configFilePath := "conf/config.json"
	configFile, err := os.Open(configFilePath)
	if err != nil {
		common.ErrorLogger.Printf("Unable to open file %s.", configFilePath)
	}
	configFileBytes, _ := io.ReadAll(configFile)

	// Unmarshal the JSON data
	var config common.Config
	err = json.Unmarshal(configFileBytes, &config)
	if err != nil {
		common.ErrorLogger.Printf("Unable to unmarshal the content of %s.", configFilePath)
	}

	// Iterate over all stations
	for i := 0; i < len(stations); i++ {
		// Retrieve and parse the station's departure data
		weatherForecast := parse.GetWeatherStationForecast(stations[i].ID, config.DB.ClientID, config.DB.ApiKey)

		// Marshal the station's departures
		data, err := json.Marshal(weatherForecast)
		if err != nil {
			common.ErrorLogger.Printf("Unable to marshal station data (Station ID: %s)", stations[i].ID)
		}

		// Write the departures into a file
		dataFilePath := "data/weatherForecast/" + stations[i].ID + ".json"
		err = os.WriteFile(dataFilePath, data, 0644)
		if err != nil {
			common.ErrorLogger.Printf("Unable to write departure data to file %s", dataFilePath)
		}
	}
}

func GetWeatherForecast() {
	// Retrieve data every hour
	getAndSaveWeatherForecast()
	ticker := time.NewTicker(time.Minute * 60)
	for range ticker.C {
		getAndSaveWeatherForecast()
	}
}
