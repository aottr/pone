package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aottr/pone/internal/types"
)

func FetchEndpointModels(ApiEndpoint string) (map[string]types.Model, error) {

	resp, err := http.Get(ApiEndpoint)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data struct {
		Models map[string]types.Model `json:"models"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Models, nil
}
