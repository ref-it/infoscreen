package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ref-it/infoscreen/server/common"
	"github.com/ref-it/infoscreen/server/handle"
	"github.com/ref-it/infoscreen/server/retrieve"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Join internally call path.Clean to prevent directory traversal
	path := filepath.Join(h.staticPath, r.URL.Path)

	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		// file does not exist or path is a directory, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func init() {
	common.InfoLogger = log.New(os.Stdout, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile)
	common.WarningLogger = log.New(os.Stdout, "[WARN]  ", log.Ldate|log.Ltime|log.Lshortfile)
	common.ErrorLogger = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	go retrieve.GetMenu()
	go retrieve.GetTrainDepartures()
	go retrieve.GetWeatherForecast()
	go retrieve.GetWarnings()
	go retrieve.GetCafeOpenings()
	go retrieve.GetCalendars()

	router := mux.NewRouter()

	router.HandleFunc("/api/menu", handle.MenuHandler)
	router.HandleFunc("/api/departures/train", handle.TrainDearturesHandler)
	router.HandleFunc("/api/departures/camil", handle.CamilDearturesHandler)
	router.HandleFunc("/api/weather/forecast", handle.WeatherForecastHandler)
	router.HandleFunc("/api/warnings", handle.WarningsHandler)
	router.HandleFunc("/api/bc-cafe", handle.CafeOpeningsHandler)
	router.HandleFunc("/api/events", handle.EventsHandler)
	//http.Handle("/", router)

	spa := spaHandler{staticPath: "../client/dist", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	fileServer1 := http.FileServer(http.Dir("../client/dist/"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fileServer1))

	// Handle static files
	fileServer := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{`*`})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	srv := &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(router),
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start webserver

	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		panic(err)
	}
}
