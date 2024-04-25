package handle

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ref-it/infoscreen/server/common"
)

func TrainDearturesHandler(w http.ResponseWriter, r *http.Request) {
	stationID, err := strconv.Atoi(r.URL.Query().Get("station"))
	if err != nil {
		panic(err)
	}

	departuresFile, err := os.Open("data/departures/" + strconv.Itoa(stationID) + ".json")
	if err != nil {
		panic(err)
	}
	departuresFileBytes, _ := io.ReadAll(departuresFile)
	var stationDepartures common.TrainStationDepartures
	err = json.Unmarshal(departuresFileBytes, &stationDepartures)
	if err != nil {
		panic(err)
	}

	var nextDepartures []common.TrainDeparture
	currentHours, _ := strconv.Atoi(time.Now().Format("15"))
	currentMinutes, _ := strconv.Atoi(time.Now().Format("04"))

	for i := 0; i < len(stationDepartures.Departures); i++ {
		departureHours, _ := strconv.Atoi(stationDepartures.Departures[i].Hours)
		departureMinutes, _ := strconv.Atoi(stationDepartures.Departures[i].Minutes)

		if departureHours == currentHours && departureMinutes >= currentMinutes || departureHours > currentHours && departureHours < currentHours+4 || departureHours == currentHours+4 && departureMinutes < currentMinutes {
			nextDepartures = append(nextDepartures, stationDepartures.Departures[i])
		}
	}

	nextStationDepartures := common.TrainStationDepartures{StationName: stationDepartures.StationName, Departures: nextDepartures}

	data, err := json.Marshal(nextStationDepartures)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
