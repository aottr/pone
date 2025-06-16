package config

import "fmt"

type Region string

const (
	EU Region = "eu"
	CA Region = "ca"
	US Region = "us"
)

var RegionAPIBase = map[Region]string{
	EU: "https://eu.api.ovh.com",
	CA: "https://ca.api.ovh.com",
	US: "https://api.us.ovhcloud.com",
}

func GetAPIBaseURL(region Region) string {
	if baseURL, ok := RegionAPIBase[region]; ok {
		return baseURL
	}
	return RegionAPIBase[EU] // default to EU
}

func GetApiVersionURL(region Region, version APIVersion) string {
	return fmt.Sprintf("%s/%s", GetAPIBaseURL(region), version)
}

func GetEndpointsURL(region Region, version APIVersion, format APIFormat) string {
	return fmt.Sprintf("%s/%s.%s", GetAPIBaseURL(region), version, format)
}

func GetEndpointURL(region Region, version APIVersion, format APIFormat, endpoint string) string {
	return fmt.Sprintf("%s/%s%s.%s", GetAPIBaseURL(region), version, endpoint, format)
}
