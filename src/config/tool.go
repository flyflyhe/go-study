package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func parseDb(path string, st *Mysql) error {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, st); err != nil {
		return err
	}

	return nil
}
