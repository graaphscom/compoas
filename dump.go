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
