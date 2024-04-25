package handle

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	canteenID, err := strconv.Atoi(r.URL.Query().Get("canteen"))
	if err != nil {
		panic(err)
	}

	menuDataFile, err := os.Open("data/menu/" + strconv.Itoa(canteenID) + ".json")
	if err != nil {
		panic(err)
	}
	menuDataFileBytes, _ := io.ReadAll(menuDataFile)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(menuDataFileBytes)
}
