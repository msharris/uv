package app

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Options struct {
	Locations []string
	Sort      Field
	Reverse   bool
}

type Field string

const (
	Id        Field = "id"
	Name      Field = "name"
	UVIndex   Field = "index"
	Time      Field = "time"
	Available Field = "status"
)

// Implement pflag.Value interface on Field for Cobra
func (f *Field) String() string {
	return string(*f)
}

func (f *Field) Set(s string) error {
	s = strings.ToLower(s)
	switch s {
	case "id", "name", "index", "time", "status":
		*f = Field(s)
	case "uv":
		*f = UVIndex
	case "date":
		*f = Time
	default:
		return errors.New(`must be one of "id", "name", "index", "uv", "time", "date", "status"`)
	}
	return nil
}

func (f *Field) Type() string {
	return "field"
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

	sort(stations, options.Sort)

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

func sort(stations []Station, f Field) {
	switch f {
	case Id:
		slices.SortFunc(stations, func(s1, s2 Station) int {
			return strings.Compare(s1.Id, s2.Id)
		})
	case Name:
		slices.SortFunc(stations, func(s1, s2 Station) int {
			return strings.Compare(s1.Name, s2.Name)
		})
	case UVIndex:
		slices.SortFunc(stations, func(s1, s2 Station) int {
			if s1.UVIndex < s2.UVIndex {
				return -1
			} else if s1.UVIndex > s2.UVIndex {
				return 1
			} else {
				return 0
			}
		})
	case Time:
		slices.SortFunc(stations, func(s1, s2 Station) int {
			return s1.Time.Compare(s2.Time)
		})
	case Available:
		slices.SortFunc(stations, func(s1, s2 Station) int {
			a1, a2 := 0, 0
			if s1.Available {
				a1 = 1
			}
			if s2.Available {
				a2 = 1
			}
			return a1 - a2
		})
	}
}
