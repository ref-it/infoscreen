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

func getAndSaveTrainDepartures() {
	common.InfoLogger.Println("Retrieving train departure data")

	// Open the config file with a list of all station names and ID numbers
	stationFilePath := "conf/stationsTrain.json"
	stationDataFile, err := os.Open(stationFilePath)
	if err != nil {
		common.ErrorLogger.Printf("Unable to open file %s.", stationFilePath)
	}
	stationDataFileBytes, err := io.ReadAll(stationDataFile)
	if err != nil {
		common.ErrorLogger.Printf("Unable to read the data in file %s.", stationFilePath)
	}

	// Unmarshal the JSON data
	var stations []common.TrainStationData
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
		departures := parse.GetTrainDepartures(stations[i].ID, config.DB.ClientID, config.DB.ApiKey)

		// Marshal the station's departures
		stationDepartures := common.TrainStationDepartures{StationName: stations[i].Name, Departures: departures}
		data, err := json.Marshal(stationDepartures)
		if err != nil {
			common.ErrorLogger.Printf("Unable to marshal station data (Station ID: %s)", strconv.Itoa(stations[i].ID))
		}

		// Write the departures into a file
		dataFilePath := "data/departures/" + strconv.Itoa(stations[i].ID) + ".json"
		err = os.WriteFile(dataFilePath, data, 0644)
		if err != nil {
			common.ErrorLogger.Printf("Unable to write departure data to file %s", dataFilePath)
		}
	}
}

func GetTrainDepartures() {
	// Retrieve data every 15 minutes
	getAndSaveTrainDepartures()
	ticker := time.NewTicker(time.Minute * 15)
	for range ticker.C {
		getAndSaveTrainDepartures()
	}
}
