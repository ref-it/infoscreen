package handle

import (
	"io"
	"net/http"
	"os"
)

func WeatherForecastHandler(w http.ResponseWriter, r *http.Request) {
	stationID := r.URL.Query().Get("station")
	forecastDataFile, err := os.Open("data/weatherForecast/" + stationID + ".json")
	if err != nil {
		panic(err)
	}
	forecastDataFileBytes, _ := io.ReadAll(forecastDataFile)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(forecastDataFileBytes)
}
