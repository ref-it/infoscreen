package parse

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ref-it/infoscreen/server/common"
)

func ParseWarnings(jsonResponse *http.Response) []common.Warning {
	var warningsDashboard []common.WarningsDashboard

	body, err := io.ReadAll(jsonResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &warningsDashboard)
	if err != nil {
		fmt.Println(1)
		panic(err)
	}

	warnings := []common.Warning{}
	for i := 0; i < len(warningsDashboard); i++ {
		warnings = append(warnings, getWarning(warningsDashboard[i].WarningID))
	}

	return warnings
}

func GetWarningsDashboard(locationId string) []common.Warning {
	var warnings []common.Warning
	url := "https://warnung.bund.de/api31/dashboard/" + locationId + ".json"
	reqest, _ := http.NewRequest("GET", url, nil)
	response, _ := http.DefaultClient.Do(reqest)
	warnings = ParseWarnings(response)

	return warnings
}

func getWarning(warningId string) common.Warning {
	url := "https://warnung.bund.de/api31/warnings/" + warningId + ".json"
	reqest, _ := http.NewRequest("GET", url, nil)
	response, _ := http.DefaultClient.Do(reqest)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var warning common.Warning
	err = json.Unmarshal(body, &warning)
	if err != nil {
		fmt.Println(1)
		panic(err)
	}

	return warning
}
