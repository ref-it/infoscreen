package retrieve

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/apognu/gocal"
	"github.com/ref-it/infoscreen/server/common"
)

func getAndSaveCalendars() {
	common.InfoLogger.Println("Retrieving calendar data")

	// Open the config file with a list of all calendarn names and ID numbers
	calendarFilePath := "conf/calendars.json"
	calendarnDataFile, err := os.Open(calendarFilePath)
	if err != nil {
		common.ErrorLogger.Printf("Unable to open file %s.", calendarFilePath)
	}
	calendarnDataFileBytes, err := io.ReadAll(calendarnDataFile)
	if err != nil {
		common.ErrorLogger.Printf("Unable to read the data in file %s.", calendarFilePath)
	}

	// Unmarshal the JSON data
	var calendars []string
	err = json.Unmarshal(calendarnDataFileBytes, &calendars)
	if err != nil {
		common.ErrorLogger.Printf("Unable to unmarshal the content of %s.", calendarFilePath)
	}

	var events []gocal.Event

	// Iterate over all calendars
	for i := 0; i < len(calendars); i++ {
		reqest, _ := http.NewRequest("GET", calendars[i], nil)
		response, _ := http.DefaultClient.Do(reqest)

		start, end := time.Now(), time.Now().Add(4*24*time.Hour)

		c := gocal.NewParser(response.Body)
		c.Start, c.End = &start, &end
		c.Parse()

		for j := 0; j < len(c.Events); j++ {
			events = append(events, c.Events[j])
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].RawStart.Value < events[j].RawStart.Value
	})

	// Marshal the station's departures
	data, err := json.Marshal(events)
	if err != nil {
		common.ErrorLogger.Printf("Unable to marshal events")
	}

	// Write the departures into a file
	dataFilePath := "data/events/events.json"
	err = os.WriteFile(dataFilePath, data, 0644)
	if err != nil {
		common.ErrorLogger.Printf("Unable to write departure data to file %s", dataFilePath)
	}
}

func GetCalendars() {
	// Retrieve data every hour
	getAndSaveCalendars()
	ticker := time.NewTicker(time.Minute * 60)
	for range ticker.C {
		getAndSaveCalendars()
	}
}
