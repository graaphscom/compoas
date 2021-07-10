package compoas

import (
	"bytes"
	"embed"
	"html/template"
	"net/http"
)

//go:embed assets
var assets embed.FS

//go:embed index.html
var indexTpl string

func UiHandler(openapiSpecUrl string, pathPrefix string, log func(v ...interface{})) (http.Handler, error) {
	indexHTML, err := execIndexHTML(openapiSpecUrl, pathPrefix)
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

func execIndexHTML(openapiSpecUrl string, pathPrefix string) ([]byte, error) {
	tp, err := template.New("index").Parse(indexTpl)
	if err != nil {
		return nil, err
	}
	indexBuff := bytes.Buffer{}
	err = tp.Execute(&indexBuff, map[string]string{"OpenapiSpecUrl": openapiSpecUrl, "PathPrefix": pathPrefix})
	if err != nil {
		return nil, err
	}

	return indexBuff.Bytes(), nil
}
