package compoas

// OAS https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#schema
type OAS struct {
	Openapi    string              `json:"openapi"`
	Info       Info                `json:"info"`
	Components *Components         `json:"components,omitempty"`
	Paths      map[string]PathItem `json:"paths"`
}

// Info https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#info-object
type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

// Components https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#componentsObject
type Components struct {
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`
	Schemas         map[string]Schema         `json:"schemas,omitempty"`
}

// SecurityScheme https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#securitySchemeObject
type SecurityScheme struct {
	Type         string `json:"type"`
	Scheme       string `json:"scheme"`
	BearerFormat string `json:"bearerFormat,omitempty"`
}

// PathItem https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#path-item-object
type PathItem struct {
	Post   *Operation `json:"post,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Get    *Operation `json:"get,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
}

// Operation https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#operationObject
type Operation struct {
	Tags        []string              `json:"tags,omitempty"`
	RequestBody *RequestBody          `json:"requestBody,omitempty"`
	Responses   map[string]Response   `json:"responses"`
	Parameters  []Parameter           `json:"parameters,omitempty"`
	Security    []SecurityRequirement `json:"security,omitempty"`
}

// Response https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#responseObject
type Response struct {
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"content,omitempty"`
}

// RequestBody https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#requestBodyObject
type RequestBody struct {
	Content map[string]MediaType `json:"content"`
}

// MediaType https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#mediaTypeObject
type MediaType struct {
	Schema *Schema `json:"schema,omitempty"`
}

// Parameter https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#parameterObject
type Parameter struct {
	Name        string      `json:"name"`
	In          string      `json:"in"`
	Description string      `json:"description,omitempty"`
	Required    bool        `json:"required,omitempty"`
	Style       string      `json:"style,omitempty"`
	Explode     bool        `json:"explode,omitempty"`
	Schema      *Schema     `json:"schema,omitempty"`
	Example     interface{} `json:"example,omitempty"`
}

// SecurityRequirement https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#securityRequirementObject
type SecurityRequirement map[string][]string

// Schema https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#schemaObject
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
