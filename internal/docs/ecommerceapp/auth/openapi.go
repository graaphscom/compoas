package auth

import "github.com/graaphscom/compoas"

var Openapi = compoas.OAS{
	Openapi: "3.0.0",
	Info: compoas.Info{
		Title:   "Auth API",
		Version: "1.0.0",
	},
	Components: &compoas.Components{
		SecuritySchemes: map[string]compoas.SecurityScheme{
			"bearerAuth": {
				Type:         "http",
				Scheme:       "bearer",
				BearerFormat: "JWT",
			},
		},
		Schemas: map[string]compoas.Schema{
			"User-Write": {
				Type:       "object",
				Properties: map[string]compoas.Schema{"username": {Type: "string"}, "password": {Type: "string"}},
			},
			"User-Read": {
				Type:       "object",
				Properties: map[string]compoas.Schema{"id": {Type: "integer"}, "username": {Type: "string"}},
			},
		},
	},
	Paths: map[string]compoas.PathItem{
		"/auth/sign-up": {
			Post: &compoas.Operation{
				Tags: []string{"Auth"},
				RequestBody: &compoas.RequestBody{Content: map[string]compoas.MediaType{
					"application/json": {Schema: &compoas.Schema{Ref: "#/components/schemas/User-Write"}},
				}},
				Responses: map[string]compoas.Response{
					"200": {Content: map[string]compoas.MediaType{
						"application/json": {Schema: &compoas.Schema{Ref: "#/components/schemas/User-Read"}},
					}},
					"400": {Description: "Validation errors."},
				},
			},
		},
	},
}
