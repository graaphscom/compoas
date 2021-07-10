package main

import (
	"github.com/graaphscom/compoas"
	"github.com/graaphscom/compoas/internal/docs/ecommerceapp/auth"
	"github.com/graaphscom/compoas/internal/docs/ecommerceapp/billing"
	"github.com/graaphscom/compoas/internal/docs/ecommerceapp/inventory"
	"log"
	"path"
)

func main() {
	const dir = "cmd/start_http_server/openapi"

	err := auth.Openapi.Dump(true, path.Join(dir, "auth.json"))
	err = billing.Openapi.Dump(true, path.Join(dir, "billing.json"))
	err = inventory.Openapi.Dump(true, path.Join(dir, "inventory.json"))

	rootOpenapi := compoas.OAS{
		Openapi: "3.0.0",
		Info: compoas.Info{
			Title:   "e-commerce app",
			Version: "1.0.0",
		},
		Components: &compoas.Components{
			Schemas:         map[string]compoas.Schema{},
			SecuritySchemes: map[string]compoas.SecurityScheme{},
		},
		Paths: map[string]compoas.PathItem{},
	}
	err = rootOpenapi.Merge(auth.Openapi).
		Merge(billing.Openapi).
		Merge(inventory.Openapi).
		Dump(true, path.Join(dir, "merged.json"))

	if err != nil {
		log.Fatalln(err)
	}
}
