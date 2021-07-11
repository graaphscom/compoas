package compoas

import (
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"testing"
)

func TestUIHandler(t *testing.T) {
	handler, _ := UIHandler(SwaggerUIBundle{Url: "/openapi.json"}, "/swagger-ui", log.Fatal)
	handlerFunc := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handler.ServeHTTP(writer, request)
	})
	require.HTTPBodyContains(t, handlerFunc, "get", "/swagger-ui", nil, "/openapi.json")
	require.HTTPBodyContains(t, handlerFunc, "get", "/swagger-ui/assets/swagger-ui-bundle.js", nil, "For license information please see swagger-ui-bundle.js.LICENSE.txt")
}
