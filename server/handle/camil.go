package handle

import (
	"io"
	"net/http"
	"os"
)

func CamilDearturesHandler(w http.ResponseWriter, r *http.Request) {
	departuresFile, err := os.Open("data/departures/camil.json")
	if err != nil {
		panic(err)
	}
	departuresFileBytes, _ := io.ReadAll(departuresFile)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(departuresFileBytes)
}
