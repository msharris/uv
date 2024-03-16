package app

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	"text/tabwriter"
	"time"
)

type Options struct {
	Locations []string
	Sort      Field
	Reverse   bool
	Quiet     bool
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
	case "location":
		*f = Name
	case "uv":
		*f = UVIndex
	case "date":
		*f = Time
	default:
		return errors.New(`must be one of "id", "name", "location", index", "uv", "time", "date", "status"`)
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
	Status    string
}

// For debugging
func (s Station) String() string {
	status := s.Status
	if s.Available {
		status = "-"
	}
	return fmt.Sprintf("%-5v%-18v%-5.1f%-7v%-v", s.Id, s.Name, s.UVIndex, s.Time.Format("15:04"), status)
}

func Run(options Options) error {
	stations, err := getStations()
	if err != nil {
		return err
	}

	if len(options.Locations) > 0 {
		stations = filter(stations, options.Locations)
	}

	if len(stations) == 0 {
		return nil
	}

	sort(stations, options.Sort)

	if options.Reverse {
		slices.Reverse(stations)
	}

	show(stations, options.Quiet)

	return nil
}

func filter(stations []Station, names []string) []Station {
	for i := range names {
		names[i] = strings.ToLower(names[i])
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
			if a1-a2 == 0 {
				return strings.Compare(s1.Status, s2.Status)
			}
			return a1 - a2
		})
	}
}

func show(stations []Station, quiet bool) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	if quiet {
		for _, s := range stations {
			status := s.Status
			if !s.Available {
				status = "(" + status + ")"
			}
			fmt.Fprintln(w, fmt.Sprintf("%v\t%.1f\t%v\t", s.Name, s.UVIndex, status))
		}
	} else {
		fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t", "Id", "Location", "Index", "Time", "Status"))
		for _, s := range stations {
			status := s.Status
			if s.Available {
				status = "-"
			}
			fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%.1f\t%v\t%v\t", s.Id, s.Name, s.UVIndex, s.Time.Format("15:04"), status))
		}
	}
	w.Flush()
}
