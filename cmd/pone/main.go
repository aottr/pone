package main

import (
	"context"
	"log"
	"os"

	"github.com/aottr/pone/internal/config"
	"github.com/aottr/pone/internal/processor"
	"github.com/urfave/cli/v3"
)

func main() {

	var apiVersion string
	var apiEndpoint string
	var outputPath string

	cmd := &cli.Command{
		Name:                  "nox",
		Usage:                 "Manage and decrypt app secrets",
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "Generate typescript types from API endpoints",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "api-version",
						Aliases:     []string{"v"},
						Usage:       "API version",
						DefaultText: "v1",
						Value:       "v1",
						Destination: &apiVersion,
					},
					&cli.StringFlag{
						Name:        "api-endpoint",
						Aliases:     []string{"e"},
						Usage:       "API endpoint",
						Value:       "/me",
						DefaultText: "/me",
						Destination: &apiEndpoint,
					},
					&cli.StringFlag{
						Name:        "output-path",
						Aliases:     []string{"o"},
						Usage:       "Output path",
						DefaultText: "./src/types.ts",
						Destination: &outputPath,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return processor.ToTypeScript(&config.RuntimeContext{
						APIVersion:  config.APIVersion(apiVersion),
						APIEndpoint: apiEndpoint,
						OutputPath:  outputPath,
					})
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Welcome to pone CLI!")
	// apis, err := api.FetchApiEndpoints(config.GetEndpointsURL(config.EU, config.V1, config.JSON))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, api := range apis {
	// 	fmt.Println(api.Path)
	// }

	// models, err := api.FetchEndpointModels(config.GetEndpointURL(config.EU, config.V1, config.JSON, "/me"))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, ok := models["nichandle.Nichandle"]
	// if !ok {
	// 	fmt.Println("model not found")
	// } else {
	// 	fmt.Println(generator.GenerateTypeScript(models, "nichandle.Nichandle"))
	// }
}
