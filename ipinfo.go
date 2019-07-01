package iptrace

import "encoding/json"

// IPInfo - represent the full ip data from the get response
type IPInfo struct {
	Country             string   `json:"country"`
	Continent           string   `json:"continent"`
	CountryFalg         string   `json:"country_falg"`
	CountryCode         string   `json:"CountryCode"`
	City                string   `json:"City"`
	Region              string   `json:"Region"`
	Lat                 float64  `json:"lat"`
	Lng                 float64  `json:"lng"`
	Tz                  string   `json:"tz"`
	Isp                 string   `json:"isp"`
	IsAnonymousProxy    bool     `json:"is_anonymous_proxy"`
	IsSatelliteProvider bool     `json:"is_satellite_provider"`
	Currency            []string `json:"currency"`
	CountryDetails      struct {
		Name struct {
			Common   string   `json:"common"`
			Official string   `json:"official"`
			Native   struct{} `json:"Native"`
		} `json:"name"`
		EuMember            bool        `json:"EuMember"`
		LandLocked          bool        `json:"LandLocked"`
		Nationality         string      `json:"Nationality"`
		Tld                 []string    `json:"tld"`
		Languages           interface{} `json:"Languages"`
		Translations        interface{} `json:"Translations"`
		Currency            []string    `json:"currency"`
		Borders             []string    `json:"Borders"`
		Cca2                string      `json:"cca2"`
		Cca3                string      `json:"cca3"`
		CIOC                string      `json:"CIOC"`
		CCN3                string      `json:"CCN3"`
		CallingCode         []string    `json:"callingCode"`
		InternationalPrefix string      `json:"InternationalPrefix"`
		Region              string      `json:"region"`
		Subregion           string      `json:"subregion"`
		Continent           string      `json:"Continent"`
		Capital             string      `json:"capital"`
		Area                int         `json:"Area"`
		Longitude           json.Number `json:"longitude"`
		Latitude            json.Number `json:"latitude"`
		MinLongitude        float64     `json:"MinLongitude"`
		MinLatitude         float64     `json:"MinLatitude"`
		MaxLongitude        float64     `json:"MaxLongitude"`
		MaxLatitude         float64     `json:"MaxLatitude"`
	} `json:"country_details"`
}
