package filemanager

import (
	"encoding/json"
	"errors"
	"os"
)

func WriteJSON(path string, data any) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create json file")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to json")
	}
	file.Close()
	return nil
}
