package inventory

import "github.com/graaphscom/compoas"

var Openapi = compoas.OAS{
	Openapi: "3.0.0",
	Info: compoas.Info{
		Title:   "Inventory API",
		Version: "1.0.0",
	},
	Components: &compoas.Components{
		Schemas: map[string]compoas.Schema{
			"Product": {Type: "object", Properties: map[string]compoas.Schema{
				"id":   {Type: "integer"},
				"name": {Type: "string"},
			}},
		},
	},
	Paths: map[string]compoas.PathItem{
		"/inventory/product/{id}": {
			Get: &compoas.Operation{
				Tags: []string{"Inventory"},
				Responses: map[string]compoas.Response{
					"200": {Content: map[string]compoas.MediaType{
						"application/json": {Schema: &compoas.Schema{Ref: "#/components/schemas/Product"}},
					}},
					"404": {Description: "Product not found"},
				},
				Parameters: []compoas.Parameter{
					{Name: "id", In: "path", Required: true, Schema: &compoas.Schema{Type: "integer"}},
				},
			},
		},
	},
}
