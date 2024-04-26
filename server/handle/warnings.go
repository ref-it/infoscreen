package handle

import (
	"io"
	"net/http"
	"os"
)

func WarningsHandler(w http.ResponseWriter, r *http.Request) {
	locationID := r.URL.Query().Get("location")
	warningsDataFile, err := os.Open("data/warnings/" + locationID + ".json")
	if err != nil {
		panic(err)
	}
	warningsDataFileBytes, _ := io.ReadAll(warningsDataFile)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(warningsDataFileBytes)
}
