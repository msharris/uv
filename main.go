package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Stations struct {
	XMLName   xml.Name   `xml:"stations"`
	Locations []Location `xml:"location"`
}

type Location struct {
	XMLName xml.Name `xml:"location"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name"`
	Index   string   `xml:"index"`
	Time    string   `xml:"time"`
	Date    string   `xml:"date"`
	Status  string   `xml:"status"`
}

func (loc Location) String() string {
	return fmt.Sprintf("%-17v %-5v %v", loc.Id, loc.Index, loc.Time)
}

func main() {
	resp, err := http.Get("https://uvdata.arpansa.gov.au/xml/uvvalues.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var stations Stations
	err = xml.Unmarshal(body, &stations)
	if err != nil {
		log.Fatal(err)
	}

	for _, loc := range stations.Locations {
		fmt.Println(loc)
	}
}

//<stations>
//  ...
//  <location id="Newcastle">
//    <name>new</name>
//    <index>0.0</index>
//    <time>6:07 PM</time>
//    <date>6/03/2024</date>
//    <fulldate>Wednesday, 6 March 2024</fulldate>
//    <utcdatetime>2024/03/06 08:07</utcdatetime>
//    <status>ok</status>
//  </location>
//  ...
//</stations>
