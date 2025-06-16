package api

import (
	"encoding/json"
	"fmt"
	"net/http"

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
