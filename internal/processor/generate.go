package processor

import (
	"fmt"
	"os"

	"github.com/aottr/pone/internal/api"
	"github.com/aottr/pone/internal/config"
	"github.com/aottr/pone/pkg/generator"
)

func ToTypeScript(ctx *config.RuntimeContext) error {

	models, err := api.FetchEndpointModels(config.GetEndpointURL(config.EU, ctx.APIVersion, config.JSON, ctx.APIEndpoint))
	if err != nil {
		return err
	}
	affectedModels, err := api.FetchModelsForEndpoint(ctx)
	if err != nil {
		return err
	}
	var typeScriptContent string
	for _, model := range affectedModels {
		if _, ok := models[model]; !ok {
			fmt.Printf("model %s not found\n", model)
			continue
		}
		typeScriptContent += generator.GenerateTypeScript(models, model)
	}

	if err := os.WriteFile(ctx.OutputPath, []byte(typeScriptContent), 0644); err != nil {
		return fmt.Errorf("failed to write decrypted file to %s: %w", ctx.OutputPath, err)
	}

	return err
}
