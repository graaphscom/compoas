package compoas

type OAS struct {
	Openapi    string          `json:"openapi"`
	Info       Info            `json:"info"`
	Components Components      `json:"components"`
	Paths      map[string]Path `json:"paths"`
}

type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type Components struct {
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes"`
	Schemas         map[string]Schema         `json:"schemas"`
}

type SecurityScheme struct {
	Type         string `json:"type"`
	Scheme       string `json:"scheme"`
	BearerFormat string `json:"bearerFormat"`
}

type Path struct {
	Post   *Operation `json:"post,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Get    *Operation `json:"get,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
}

type Operation struct {
	Tags        []string              `json:"tags,omitempty"`
	Responses   map[string]Response   `json:"responses,omitempty"`
	RequestBody *RequestBody          `json:"requestBody,omitempty"`
	Parameters  []Parameter           `json:"parameters,omitempty"`
	Security    []SecurityRequirement `json:"security,omitempty"`
}

type Response struct {
	Description string             `json:"description,omitempty"`
	Content     map[string]Content `json:"content,omitempty"`
}

type RequestBody struct {
	Content map[string]Content `json:"content"`
}

type Content struct {
	Schema Schema `json:"schema"`
}

type Parameter struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	In          string   `json:"in"`
	Schema      Schema   `json:"schema"`
	Required    bool     `json:"required"`
	Style       string   `json:"style,omitempty"`
	Explode     bool     `json:"explode,omitempty"`
	Example     []string `json:"example,omitempty"`
}

type SecurityRequirement map[string][]string

type Schema struct {
	Type        string            `json:"type,omitempty"`
	Format      string            `json:"format,omitempty"`
	Properties  map[string]Schema `json:"properties,omitempty"`
	Enum        []string          `json:"enum,omitempty"`
	Items       *Schema           `json:"items,omitempty"`
	Nullable    bool              `json:"nullable,omitempty"`
	Ref         string            `json:"$ref,omitempty"`
	Required    []string          `json:"required,omitempty"`
	MinLength   int               `json:"minLength,omitempty"`
	MaxLength   int               `json:"maxLength,omitempty"`
	MinItems    int               `json:"minItems,omitempty"`
	MaxItems    int               `json:"maxItems,omitempty"`
	Description string            `json:"description,omitempty"`
	OneOf       []Schema          `json:"oneOf,omitempty"`
}
