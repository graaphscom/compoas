package compoas

import (
	"encoding/json"
	"os"
)

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
