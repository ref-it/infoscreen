package handle

import (
	"io"
	"net/http"
	"os"
)

func CafeOpeningsHandler(w http.ResponseWriter, r *http.Request) {
	cafeOpeningsDataFile, err := os.Open("data/ilscEvents/bc-cafe.json")
	if err != nil {
		panic(err)
	}
	cafeOpeningsDataFileBytes, _ := io.ReadAll(cafeOpeningsDataFile)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(cafeOpeningsDataFileBytes)
}
