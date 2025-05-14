package config

import (
	"encoding/json"
	"os"
)

func LoadConfig(configFile string, c interface{}) error {
	file, err := os.Open("/Users/shalini/CumulativePoetry/src/config/json/" + configFile)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&c); err != nil {
		return err
	}
	return nil
}
