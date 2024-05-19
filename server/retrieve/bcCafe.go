package retrieve

import (
	"encoding/json"
	"os"
	"time"

	"github.com/ref-it/infoscreen/server/common"
	"github.com/ref-it/infoscreen/server/parse"
)

func getAndSaveCafeOpenings() {
	common.InfoLogger.Println("Retrieving cafe openings data")

	// Retrieve and parse the cafe openings data
	cafeOpenings := parse.GetIlscEvent("https://lara.il-sc.de/api/events/bc-cafe")

	// Marshal the cafe openings data
	data, err := json.Marshal(cafeOpenings)
	if err != nil {
		common.ErrorLogger.Printf("Unable to marshal cafe openings data")
	}

	// Write the departures into a file
	dataFilePath := "data/ilscEvents/bc-cafe.json"
	err = os.WriteFile(dataFilePath, data, 0644)
	if err != nil {
		common.ErrorLogger.Printf("Unable to write cafe openings data to file %s", dataFilePath)
	}
}

func GetCafeOpenings() {
	// Retrieve data every hour
	getAndSaveCafeOpenings()
	ticker := time.NewTicker(time.Minute * 60)
	for range ticker.C {
		getAndSaveCafeOpenings()
	}
}
