package main

import (
	"fmt"

	"github.com/aottr/pone/internal/api"
	"github.com/aottr/pone/internal/config"
)

func main() {
	fmt.Println("Welcome to pone CLI!")
	apis, err := api.FetchApiEndpoints(config.GetEndpointsURL(config.EU, config.V1, config.JSON))
	if err != nil {
		fmt.Println(err)
	}

	for _, api := range apis {
		fmt.Println(api.Path)
	}

	err = api.FetchEndpointModels(config.GetEndpointURL(config.EU, config.V1, config.JSON, "/me"))
	if err != nil {
		fmt.Println(err)
	}

}
