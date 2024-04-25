package common

import (
	"log"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

type Canteen struct {
	Name string `json:"name"`
	City string `json:"city"`
	ID   int    `json:"id"`
}

type CanteenData struct {
	Canteens  []Canteen         `json:"canteens"`
	Additives map[string]string `json:"additives"`
	Allergens map[string]string `json:"allergens"`
}

type Config struct {
	DB DBApiCredentials `json:"db"`
}

type DBApiCredentials struct {
	ClientID string `json:"clientId"`
	ApiKey   string `json:"apiKey"`
}

type Meal struct {
	Name         string   `json:"name"`
	Additives    []string `json:"additives"`
	Allergens    []string `json:"allergens"`
	Prices       Prices   `json:"prices"`
	IsVegetarian bool     `json:"isVegetarian"`
	IsVegan      bool     `json:"isVegan"`
}

type Menu struct {
	CanteenName string `json:"canteenName"`
	Lunch       []Meal `json:"lunch"`
	Dinner      []Meal `json:"dinner"`
	LastUpdated int64  `json:"lastUpdated"`
}

type Prices struct {
	Students  float32 `json:"students"`
	Employees float32 `json:"employees"`
	Guests    float32 `json:"guests"`
}

type TrainStationData struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type TrainDeparture struct {
	Category    string `json:"category"`
	Line        string `json:"line"`
	Hours       string `json:"hours"`
	Minutes     string `json:"minutes"`
	Destination string `json:"destination"`
	Track       string `json:"track"`
}

type TrainStationDepartures struct {
	StationName string           `json:"stationName"`
	Departures  []TrainDeparture `json:"departures"`
}

type WeatherStationData struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type WeatherForecast struct {
	Trend    WeatherForecastTrend    `json:"trend"`
	Days     []WeatherForecastDays   `json:"days"`
	Forecast WeatherForecastForecast `json:"forecast"`
}

type WeatherForecastTrend struct {
	Start                        int   `json:"start"`
	Temperature                  []int `json:"temperature"`
	TemperatureStd               []int `json:"temperatureStd"`
	PrecipitationProbabilityHigh []int `json:"precipitationProbabilityHigh"`
	PrecipitationProbabilityLow  []int `json:"precipitationProbabilityLow"`
}

type WeatherForecastDays struct {
	StationID      string `json:"stationId"`
	DayDate        string `json:"dayDate"`
	TemperatureMin int    `json:"temperatureMin"`
	TemperatureMax int    `json:"temperatureMax"`
	Precipitation  int    `json:"precipitation"`
	WindSpeed      int    `json:"windSpeed"`
	WindGust       int    `json:"windGust"`
	WindDirection  int    `json:"windDirection"`
	Sunshine       int    `json:"sunshine"`
	Sunrise        int    `json:"sunrise"`
	Sunset         int    `json:"sunset"`
	Moonrise       int    `json:"moonrise"`
	Moonset        int    `json:"moonset"`
	MoonPhase      int    `json:"moonPhase"`
	Icon           int    `json:"icon"`
	Icon1          int    `json:"icon1"`
	Icon2          int    `json:"icon2"`
}

type WeatherForecastForecast struct {
	Start              int   `json:"start"`
	Temperature        []int `json:"temperature"`
	TemperatureStd     []int `json:"temperatureStd"`
	WindDirection      []int `json:"windDirection"`
	WindGust           []int `json:"windGust"`
	WindSpeed          []int `json:"windSpeed"`
	Icon               []int `json:"icon"`
	PrecipitationTotal []int `json:"precipitationTotal"`
}
