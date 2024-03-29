package app

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var locations = map[string]string{
	"Adelaide":         "Australia/Adelaide",
	"Alice Springs":    "Australia/Darwin",
	"Brisbane":         "Australia/Brisbane",
	"Canberra":         "Australia/Sydney",
	"Casey":            "Antarctica/Casey",
	"Darwin":           "Australia/Darwin",
	"Davis":            "Antarctica/Davis",
	"Emerald":          "Australia/Brisbane",
	"Gold Coast":       "Australia/Brisbane",
	"Kingston":         "Australia/Hobart",
	"Macquarie Island": "Antarctica/Macquarie",
	"Mawson":           "Antarctica/Mawson",
	"Melbourne":        "Australia/Melbourne",
	"Newcastle":        "Australia/Sydney",
	"Perth":            "Australia/Perth",
	"Sydney":           "Australia/Sydney",
	"Townsville":       "Australia/Brisbane",
}

type XMLStations struct {
	XMLName   xml.Name      `xml:"stations"`
	Locations []XMLLocation `xml:"location"`
}

type XMLLocation struct {
	XMLName     xml.Name `xml:"location"`
	Id          string   `xml:"id,attr"`
	Name        string   `xml:"name"`
	Index       string   `xml:"index"`
	Time        string   `xml:"time"`
	Date        string   `xml:"date"`
	FullDate    string   `xml:"fulldate"`
	UTCDateTime string   `xml:"utcdatetime"`
	Status      string   `xml:"status"`
}

func (xmlStations *XMLStations) Stations() []Station {
	var stations []Station
	for _, xmlLoc := range xmlStations.Locations {
		stations = append(stations, xmlLoc.Station())
	}
	return stations
}

func (xmlLoc *XMLLocation) Station() Station {
	station := Station{
		Id:        strings.ToUpper(xmlLoc.Name),
		Name:      xmlLoc.Id,
		Available: xmlLoc.Status == "ok",
	}

	station.UVIndex, _ = strconv.ParseFloat(xmlLoc.Index, 64)

	station.Time, _ = time.Parse("2006/01/02 15:04", xmlLoc.UTCDateTime)
	if locName, err := getLocationName(station.Name); err == nil {
		loc, _ := time.LoadLocation(locName)
		station.Time = station.Time.In(loc)
	}

	if xmlLoc.Status != "ok" {
		station.Status = xmlLoc.Status
	}

	return station
}

func getLocationName(stationName string) (string, error) {
	if name, ok := locations[stationName]; !ok {
		return name, fmt.Errorf("no time zone configured for station %v", stationName)
	} else {
		return name, nil
	}
}

func getStations() ([]Station, error) {
	resp, err := http.Get("https://uvdata.arpansa.gov.au/xml/uvvalues.xml")
	if err != nil {
		return nil, errors.New("ARPANSA data file unavailable")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("unable to read ARPANSA data file")
	}

	var xmlStations XMLStations
	err = xml.Unmarshal(body, &xmlStations)
	if err != nil {
		return nil, errors.New("unexpected format of ARPANSA data file")
	}

	return xmlStations.Stations(), nil
}
