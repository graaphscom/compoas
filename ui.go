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

func UiHandler(uiBundle SwaggerUIBundle, pathPrefix string, log func(v ...interface{})) (http.Handler, error) {
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

func execIndexHTML(uiBundle SwaggerUIBundle, pathPrefix string) ([]byte, error) {
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

type SwaggerUIBundle struct {
	Url  string               `json:"url,omitempty"`
	Urls []SwaggerUIBundleUrl `json:"urls,omitempty"`
}

type SwaggerUIBundleUrl struct {
	Url  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}
