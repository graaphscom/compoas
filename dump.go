package compoas

import (
	"encoding/json"
	"os"
)

func Dump(input OAS, doIndent bool, fileName string) error {
	var (
		data []byte
		err  error
	)

	if doIndent {
		data, err = json.MarshalIndent(input, "", "  ")
	} else {
		data, err = json.Marshal(input)
	}

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
