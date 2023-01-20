package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func UpdateFile() {
	content, err := json.Marshal(Users)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile("userfile.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	LoadFile()
}
func LoadFile() {
	content, err := ioutil.ReadFile("userfile.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &Users)
	if err != nil {
		log.Fatal(err)
	}
}
