package app

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

type Options struct {
	Locations []string
	Sort      string
	Reverse   bool
}

type Station struct {
	Id        string
	Name      string
	UVIndex   float64
	Time      time.Time
	Available bool
}

func (s Station) String() string {
	return fmt.Sprintf("%-17v %4.1f  %v", s.Name, s.UVIndex, s.Time.Format("Mon 2 Jan 3:04 pm"))
}

func Run(options Options) error {
	stations, err := getStations()
	if err != nil {
		return err
	}

	if len(options.Locations) > 0 {
		stations = filter(stations, options.Locations)
	}

	if options.Reverse {
		slices.Reverse(stations)
	}

	for _, s := range stations {
		fmt.Println(s)
	}

	return nil
}

func filter(stations []Station, names []string) []Station {
	for _, n := range names {
		strings.ToLower(n)
	}

	var filtered []Station
	for _, station := range stations {
		if slices.Contains(names, strings.ToLower(station.Id)) || slices.Contains(names, strings.ToLower(station.Name)) {
			filtered = append(filtered, station)
		}
	}
	return filtered
}
