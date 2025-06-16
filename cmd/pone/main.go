package main

import (
	"fmt"

	"github.com/aottr/pone/internal/api"
	"github.com/aottr/pone/internal/config"
	"github.com/aottr/pone/pkg/generator"
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

	models, err := api.FetchEndpointModels(config.GetEndpointURL(config.EU, config.V1, config.JSON, "/me"))
	if err != nil {
		fmt.Println(err)
	}
	_, ok := models["audit.Log"]
	if !ok {
		fmt.Println("model not found")
	} else {
		fmt.Println(generator.GenerateTypeScriptRecursive2(models, "audit.Log"))
	}
}
