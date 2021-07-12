package main

import (
	"embed"
	"github.com/graaphscom/compoas"
	"log"
	"net/http"
)

//go:embed openapi
var dumpedSpecs embed.FS

func main() {
	oasHandler, err := compoas.UIHandler(
		compoas.SwaggerUIBundleConfig{Urls: []compoas.SwaggerUIBundleUrl{
			{Url: "/openapi/merged.json", Name: "All"},
			{Url: "/openapi/auth.json", Name: "Auth"},
			{Url: "/openapi/billing.json", Name: "Billing"},
			{Url: "/openapi/inventory.json", Name: "Inventory"},
		}},
		"/swagger-ui",
		log.Fatalln,
	)
	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/swagger-ui/", oasHandler)
	mux.Handle("/openapi/", http.FileServer(http.FS(dumpedSpecs)))

	log.Fatalln(http.ListenAndServe(":8080", mux))
}
