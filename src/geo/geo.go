package geo

import (
	"log"

	"github.com/abh/geoip"
)

type Geo struct {
	gi *geoip.GeoIP
}
type Loc struct {
	Iso2Code    string  `json:"iso2_code"`
	Iso3Code    string  `json:"iso3_code"`
	CountryName string  `json:"country_name"`
	Region      string  `json:"region"`
	City        string  `json:"city"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
}

func New() *Geo {
	g := &Geo{}
	gi, err := geoip.Open("../github.com/abh/geoip/db/GeoLiteCity.dat")
	g.gi = gi
	if err != nil {
		log.Fatalln(err.Error())
	}
	return g
}

func (g *Geo) GetLoc(ipAddress string) Loc {
	record := g.gi.GetRecord(ipAddress)
	if record != nil {
		loc := Loc{
			Iso2Code:    record.CountryCode,
			Iso3Code:    record.CountryCode3,
			CountryName: record.CountryName,
			Region:      record.Region,
			City:        record.City,
			Latitude:    record.Latitude,
			Longitude:   record.Longitude,
		}
		return loc
	}
	return Loc{}
}
