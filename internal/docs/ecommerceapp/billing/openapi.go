package billing

import (
	"github.com/graaphscom/compoas"
	"github.com/graaphscom/compoas/internal/docs/ecommerceapp/auth"
	"github.com/graaphscom/compoas/internal/docs/ecommerceapp/inventory"
)

var Openapi = compoas.OAS{
	Openapi: "3.0.0",
	Info: compoas.Info{
		Title:   "Billing API",
		Version: "1.0.0",
	},
	Components: &compoas.Components{
		Schemas: map[string]compoas.Schema{
			"Order": {Type: "object", Properties: map[string]compoas.Schema{
				"id":    {Type: "integer"},
				"buyer": {Ref: "#/components/schemas/User-Read"},
				"items": {Type: "array", Items: &compoas.Schema{Ref: "#/components/schemas/Product"}},
			}},
			"User-Read": auth.Openapi.Components.Schemas["User-Read"],
			"Product":   inventory.Openapi.Components.Schemas["Product"],
		},
		SecuritySchemes: auth.Openapi.Components.SecuritySchemes,
	},
	Paths: map[string]compoas.PathItem{
		"/billing/orders": {
			Get: &compoas.Operation{
				Tags: []string{"Billing"},
				Responses: map[string]compoas.Response{
					"200": {Content: map[string]compoas.MediaType{
						"application/json": {Schema: &compoas.Schema{
							Type:  "array",
							Items: &compoas.Schema{Ref: "#/components/schemas/Order"},
						}},
					}},
				},
				Security: []compoas.SecurityRequirement{{"bearerAuth": {}}},
			},
		},
	},
}
