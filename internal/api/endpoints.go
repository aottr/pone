package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/aottr/pone/internal/config"
	"github.com/aottr/pone/internal/types"
)

// {
// 	"path": "/allDom",
// 	"schema": "{path}.{format}",
// 	"format": [
// 	  "json",
// 	  "yaml"
// 	],
// 	"description": ""
//   },

// this is actually a list of endpoint prefixes for different APIs
func FetchApiEndpoints(region string) ([]types.EndpointMeta, error) {
	resp, err := http.Get(region)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data struct {
		Apis []types.EndpointMeta `json:"apis"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Apis, nil
}

func FetchModelsForEndpoint(ctx *config.RuntimeContext) ([]string, error) {

	resp, err := http.Get(config.GetEndpointURL(config.EU, ctx.APIVersion, config.JSON, ctx.APIEndpoint))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data struct {
		Resources []types.Resource       `json:"apis"` // must be named endpoint, the other is the resource
		Models    map[string]types.Model `json:"models"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	ri := slices.IndexFunc(data.Resources, func(r types.Resource) bool {
		return r.Path == ctx.APIEndpoint
	})

	if ri == -1 {
		return nil, fmt.Errorf("resource not found")
	}

	var affectedModels []string
	for _, operation := range data.Resources[ri].Operations {
		affectedModels = append(affectedModels, operation.ResponseType)
	}

	return affectedModels, nil
}
