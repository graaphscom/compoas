package compoas

import (
	"encoding/json"
	"os"
)

// Dump allows for dumping specification into a file.
func (oas OAS) Dump(prettyPrint bool, fileName string) error {
	var (
		data []byte
		err  error
	)

	if prettyPrint {
		data, err = json.MarshalIndent(oas, "", "  ")
	} else {
		data, err = json.Marshal(oas)
	}

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

// Merge allows for merging multiple specifications.
// It merges paths, components.schemas and components.securitySchemes.
// Root specification must have initialized these fields:
//  rootOpenapi := compoas.OAS{
//		Openapi: "3.0.0",
//		Info: compoas.Info{
//			Title:   "merged spec",
//			Version: "1.0.0",
//		},
//		Components: &compoas.Components{
//			Schemas:         map[string]compoas.Schema{},
//			SecuritySchemes: map[string]compoas.SecurityScheme{},
//		},
//		Paths: map[string]compoas.PathItem{},
//	}
//  rootOpenapi.Merge(anotherOpenapi)
func (oas *OAS) Merge(source OAS) *OAS {
	for k, v := range source.Paths {
		oas.Paths[k] = v
	}
	for k, v := range source.Components.Schemas {
		oas.Components.Schemas[k] = v
	}
	for k, v := range source.Components.SecuritySchemes {
		oas.Components.SecuritySchemes[k] = v
	}
	return oas
}
