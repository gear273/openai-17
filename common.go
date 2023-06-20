package openai

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// https://json-schema.org/understanding-json-schema/
// JSON Type
const (
	JSONSchemaTypeNull    JSONSchemaType = "null"
	JSONSchemaTypeArray   JSONSchemaType = "array"
	JSONSchemaTypeObject  JSONSchemaType = "object"
	JSONSchemaTypeNumber  JSONSchemaType = "number"
	JSONSchemaTypeString  JSONSchemaType = "string"
	JSONSchemaTypeBoolean JSONSchemaType = "boolean"
)

type JSONSchemaType string

type JSONSchemaDefine struct {
	Type        JSONSchemaType               `json:"type,omitempty"`
	Description string                       `json:"description,omitempty"`
	Enum        []string                     `json:"enum,omitempty"`
	Properties  map[string]*JSONSchemaDefine `json:"properties,omitempty"`
	Required    []string                     `json:"required,omitempty"`
	Items       *JSONSchemaDefine            `json:"items,omitempty"`
}
