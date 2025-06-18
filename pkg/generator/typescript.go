package generator

import (
	"strings"

	"github.com/aottr/pone/internal/types"
)

var processedTypes = make(map[string]string)
var output strings.Builder

func toTSBaseType(p *types.Property) string {
	switch p.Type {
	case "string", "uuid", "ip", "ipBlock":
		return "string"
	case "integer":
		return "number"
	case "boolean":
		return "boolean"
	case "array":
		return "Array<any>"
	case "object":
		return "{ [key: string]: any }"
	case "date", "datetime":
		return "Date"
	default:
		if typeName, ok := processedTypes[p.Type]; ok {
			return typeName
		}
		return "any"
	}
}

func toTSType(p *types.Property) string {
	// Check if type ends with [] for array notation
	if strings.HasSuffix(p.Type, "[]") {
		baseType := strings.TrimSuffix(p.Type, "[]")
		return "Array<" + baseType + ">"
	}
	return toTSBaseType(p)
}

func GenerateTypeScript(models map[string]types.Model, modelName string) string {
	output.Reset()
	generateType(models, modelName)
	return output.String()
}

func generateEnum(models map[string]types.Model, modelName string) {
	if processedTypes[modelName] != "" {
		return
	}
	model, exists := models[modelName]
	if !exists {
		return
	}
	processedTypes[modelName] = model.Id
	output.WriteString("/* " + modelName + " */\n")
	output.WriteString("export enum " + model.Id + " {\n")
	for _, value := range model.Enum {
		output.WriteString("  " + value.(string) + ",\n")
	}
	output.WriteString("};\n\n")
}

func generateType(models map[string]types.Model, modelName string) {

	if processedTypes[modelName] != "" {
		return
	}
	model, exists := models[modelName]
	if !exists {
		return
	}

	if model.Properties == nil {
		generateEnum(models, modelName)
		return
	}
	for _, property := range model.Properties {
		if _, isCustomType := models[property.Type]; isCustomType {
			generateType(models, property.Type)
		}
	}

	processedTypes[modelName] = model.Id
	output.WriteString("/* " + modelName + " */\n")
	output.WriteString("export type " + model.Id + " = {\n")

	for key, value := range model.Properties {
		output.WriteString("  " + key + ": " + toTSType(&value) + ";\n")
	}
	output.WriteString("};\n\n")
}
