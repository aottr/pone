package types

type EndpointMeta struct {
	Description string   `json:"description"`
	Format      []string `json:"format"`
	Path        string   `json:"path"`
	Schema      string   `json:"schema"`
}

type Property struct {
	Type        string `json:"type"`
	FullType    string `json:"fullType"`
	CanBeNull   bool   `json:"canBeNull"`
	ReadOnly    bool   `json:"readOnly"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type Model struct {
	Id          string              `json:"id"`
	Namespace   string              `json:"namespace"`
	Description string              `json:"description"`
	Properties  map[string]Property `json:"properties,omitempty"`
	Enum        []any               `json:"enum,omitempty"` // use interface{} for string or int enums
	EnumType    *string             `json:"enumType,omitempty"`
}

type Parameter struct {
	Name        string `json:"name"`
	DataType    string `json:"dataType"`
	Type        string `json:"paramType"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
}

type Operation struct {
	Id              string      `json:"operationId"`
	HttpMethod      string      `json:"httpMethod"`
	Parameters      []Parameter `json:"parameters"`
	ResponseType    string      `json:"responseType"` // because responseType can be void
	Unauthenticated bool        `json:"noAuthentication"`
}

type Resource struct {
	Path        string      `json:"path"`
	Description string      `json:"description"`
	Operations  []Operation `json:"operations"`
}

type ApiEndpoint struct {
	ResourcePath string     `json:"resourcePath"`
	Apis         []Resource `json:"apis"`
	BasePath     string     `json:"basePath"`
	Models       *[]Model   `json:"models"`
}
