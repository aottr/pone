package config

type APIVersion string

const (
	V1 APIVersion = "v1"
	V2 APIVersion = "v2"
)

type APIFormat string

const (
	JSON APIFormat = "json"
	YAML APIFormat = "yaml"
)
