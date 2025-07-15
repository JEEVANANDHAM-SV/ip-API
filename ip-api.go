package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// IPInfo maps the API response structure
type IPInfo struct {
	Query      string  `json:"query"`
	Country    string  `json:"country"`
	RegionName string  `json:"regionName"`
	City       string  `json:"city"`
	ISP        string  `json:"isp"`
	Org        string  `json:"org"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lon"`
	Timezone   string  `json:"timezone"`
	Status     string  `json:"status"`
	Message    string  `json:"message"` // for error message if failed
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <ip_address>")
		return
	}

	ip := os.Args[1]
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer resp.Body.Close()

	var info IPInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	if info.Status != "success" {
		fmt.Printf("API Error: %s\n", info.Message)
		return
	}

	// Output the result
	fmt.Println("----- IP Information -----")
	fmt.Printf("IP Address:  %s\n", info.Query)
	fmt.Printf("Country:     %s\n", info.Country)
	fmt.Printf("Region:      %s\n", info.RegionName)
	fmt.Printf("City:        %s\n", info.City)
	fmt.Printf("ISP:         %s\n", info.ISP)
	fmt.Printf("Organization:%s\n", info.Org)
	fmt.Printf("Location:    %f, %f\n", info.Latitude, info.Longitude)
	fmt.Printf("Timezone:    %s\n", info.Timezone)
}
