package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func readFile(fileName string) *os.File {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return jsonFile
}

func readJsonFromFile(fileName string, v any) {
	jsonFile := readFile(fileName)
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValue, v)

}
