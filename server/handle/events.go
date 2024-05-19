package handle

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/apognu/gocal"
	"github.com/ref-it/infoscreen/server/common"
)

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	eventsDataFile, err := os.Open("data/events/events.json")
	if err != nil {
		panic(err)
	}
	eventsDataFileBytes, _ := io.ReadAll(eventsDataFile)

	var eventsFull []gocal.Event
	err = json.Unmarshal(eventsDataFileBytes, &eventsFull)
	if err != nil {
		common.ErrorLogger.Printf("Unable to unmarshal events data")
	}

	var events [][]common.Event
	var eventsOfOneDay []common.Event
	for i := 0; i < len(eventsFull); i++ {
		var event common.Event
		event.Summary = eventsFull[i].Summary
		event.StartDate = eventsFull[i].Start.Format("2006-01-02")
		event.StartTime = eventsFull[i].Start.Format("15:04")
		event.EndDate = eventsFull[i].End.Format("2006-01-02")
		event.EndTime = eventsFull[i].End.Format("15:04")
		if i != 0 && eventsOfOneDay[0].StartDate != event.StartDate {
			events = append(events, eventsOfOneDay)
			eventsOfOneDay = []common.Event{}
		}
		eventsOfOneDay = append(eventsOfOneDay, event)
	}
	events = append(events, eventsOfOneDay)

	data, err := json.Marshal(events)
	if err != nil {
		common.ErrorLogger.Printf("Unable to unmarshal events data")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
