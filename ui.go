package compoas

import (
	"bytes"
	"embed"
	"encoding/json"
	"html/template"
	"net/http"
)

//go:embed assets
var assets embed.FS

//go:embed index.html
var indexTpl string

// UIHandler serves Swagger UI (index.html and all required js, css assets)
//
// uiBundle argument is a configuration that will be rendered inside index.html as input to the SwaggerUIBundle.
//
// pathPrefix argument allows for setting path under which Swagger UI will be accessible.
// Eg If we want to have Swagger UI under http://example.com/swagger-ui, we would set pathPrefix to "/swagger-ui"
// (leading slash, no trailing slash).
// If no nesting is needed, set pathPrefix to "/".
//
// log argument is being used for logging error when http.ResponseWriter could not write a response
func UIHandler(uiBundle SwaggerUIBundleConfig, pathPrefix string, log func(v ...interface{})) (http.Handler, error) {
	indexHTML, err := execIndexHTML(uiBundle, pathPrefix)
	if err != nil {
		return nil, err
	}

	assetsHandler := http.FileServer(http.FS(assets))
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "" || request.URL.Path == "/" {
			_, err = writer.Write(indexHTML)
			if err != nil {
				log(err)
			}
		} else {
			assetsHandler.ServeHTTP(writer, request)
		}
	}

	return http.StripPrefix(pathPrefix, handler), nil
}

func execIndexHTML(uiBundle SwaggerUIBundleConfig, pathPrefix string) ([]byte, error) {
	tp, err := template.New("index").Parse(indexTpl)
	if err != nil {
		return nil, err
	}
	indexBuff := bytes.Buffer{}
	uiBundleJson, err := json.Marshal(uiBundle)
	if err != nil {
		return nil, err
	}
	err = tp.Execute(&indexBuff, map[string]interface{}{"UIBundleConfig": template.JS(uiBundleJson), "PathPrefix": pathPrefix})
	if err != nil {
		return nil, err
	}

	return indexBuff.Bytes(), nil
}

// SwaggerUIBundleConfig reflects config to the SwaggerUIBundle.
// https://github.com/swagger-api/swagger-ui/blob/bb21c6df52eb12cd4bdbf8c29feb500795595fa8/dist/index.html#L41
type SwaggerUIBundleConfig struct {
	Url  string               `json:"url,omitempty"`
	Urls []SwaggerUIBundleUrl `json:"urls,omitempty"`
}

type SwaggerUIBundleUrl struct {
	Url  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}
