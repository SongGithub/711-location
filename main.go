package main

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"net/http"
	"time"
)

var API_URL = "https://projectzerothree.info/api.php?format=json"
var FUEL_TYPE = "U91"
var REGION = "All"

type Config struct {
	FuelType string `envconfig:"FUEL_TYPE" default:"U91"`
}

type ServoPrices struct {
	Lat      float32 `json:"lat,omitempty"`
	Lng      float32 `json:"lng,omitempty"`
	Postcode string  `json:"postcode,omitempty"`
	Price    float32 `json:"price,omitempty"`
	State    string  `json:"state,omitempty"`
	Suburb   string  `json:"suburb,omitempty"`
	Type     string  `json:"type,omitempty"`
}

type RegionalPrices struct {
	Region string `json:"region,omitempty"`
	Prices []ServoPrices
}

type AllRegionResults struct {
	Updated int              `json:"updated,omitempty"`
	Regions []RegionalPrices `json:"regions,omitempty"`
}

// curl prototype
// curl -L https://projectzerothree.info/api.php\?format\=json | jq '.regions[] | select(.region=="All") | .prices[] | select(.type=="U91") | .lng,.lat'
func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("Getting data from Project03 API... \n")
	results := new(AllRegionResults)
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(API_URL)
	if err != nil {
		//
	}
	defer r.Body.Close()
	error := json.NewDecoder(r.Body).Decode(&results)
	if error != nil {
		fmt.Print("error happened\n")
		fmt.Print(json.NewDecoder(r.Body))
		fmt.Print(error)
		return
	}
	regions := results.Regions
	fmt.Printf("Searching received data for fuel type: %s, in region \"%s\" ... \n", cfg.FuelType, REGION)
	for _, r := range regions {
		if r.Region == REGION {
			prices := r.Prices
			for _, t := range prices {
				if t.Type == cfg.FuelType {
					// TODO: add error-handling here when coordinates can't be located
					fmt.Print("Latitude（经度）:", t.Lat, "\n", "Longitude（纬度）:", t.Lng, "\n")
				}
			}
		}
	}
}
