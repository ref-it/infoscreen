package retrieve

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/ref-it/infoscreen/server/common"
	"github.com/ref-it/infoscreen/server/parse"
)

func getAndSaveWarnings() {
	common.InfoLogger.Println("Retrieving NINA warnings data")

	// Open the config file with a list of all station names and ID numbers
	stationFilePath := "conf/locationsWarnings.json"
	stationDataFile, err := os.Open(stationFilePath)
	if err != nil {
		common.ErrorLogger.Printf("Unable to open file %s.", stationFilePath)
	}
	stationDataFileBytes, err := io.ReadAll(stationDataFile)
	if err != nil {
		common.ErrorLogger.Printf("Unable to read the data in file %s.", stationFilePath)
	}

	// Unmarshal the JSON data
	var locations []common.WeatherStationData
	err = json.Unmarshal(stationDataFileBytes, &locations)
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

	// Iterate over all locations
	for i := 0; i < len(locations); i++ {
		// Retrieve and parse the station's departure data
		warnings := parse.GetWarningsDashboard(locations[i].ID)

		// Marshal the station's departures
		data, err := json.Marshal(warnings)
		if err != nil {
			common.ErrorLogger.Printf("Unable to marshal location data (Location ID: %s)", locations[i].ID)
		}

		// Write the departures into a file
		dataFilePath := "data/warnings/" + locations[i].ID + ".json"
		err = os.WriteFile(dataFilePath, data, 0644)
		if err != nil {
			common.ErrorLogger.Printf("Unable to write warnings data to file %s", dataFilePath)
		}
	}
}

func GetWarnings() {
	// Retrieve data 15 minutes
	getAndSaveWarnings()
	ticker := time.NewTicker(time.Minute * 15)
	for range ticker.C {
		getAndSaveWarnings()
	}
}
