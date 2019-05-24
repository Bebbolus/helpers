package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//ReadFromJSON function load a json file into a struct or return error
func ReadJson(t interface{}, filename string) error {

	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(jsonFile), t)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	return nil
}