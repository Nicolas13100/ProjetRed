package game

import (
	"encoding/json"
	"io/ioutil"
)

func SaveGame(data SaveData, filename string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadGame(filename string) (SaveData, error) {
	var data SaveData

	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
