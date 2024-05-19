package parse

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ref-it/infoscreen/server/common"
)

func GetIlscEvent(url string) common.LaraEventList {
	var laraEvents common.LaraEventList

	reqest, _ := http.NewRequest("GET", url, nil)
	response, _ := http.DefaultClient.Do(reqest)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &laraEvents)
	if err != nil {
		panic(err)
	}

	return laraEvents
}
