package app

import (
	"fmt"
	"time"
)

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

func Run() error {
	stations, err := getStations()
	if err != nil {
		return err
	}

	for _, s := range stations {
		fmt.Println(s)
	}

	return nil
}
