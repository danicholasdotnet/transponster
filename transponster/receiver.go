package transponster

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
)

func (io IO) LogRequest() {
	log.Println("Incoming: { ", Details(io.R), " }")
}

func (io IO) RequestToStruct(i interface{}) error {
	b, err := ioutil.ReadAll(io.R.Body)
	if err != nil {
		log.Println("Error reading the request body.")
		log.Printf("Request: %#v", &io.R)
		return err
	}

	if err := json.Unmarshal(b, i); err != nil {
		log.Println("Error unmarshaling the JSON that was sent.")
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, b, "", "\t")
		log.Println("JSON: ", &prettyJSON)
		return err
	}
	decoder := json.NewDecoder(io.R.Body)
	if err := decoder.Decode(i); err != nil {
		log.Println("Error decoding request body to struct.")
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, b, "", "\t")
		log.Println("Source: ", &prettyJSON)
		log.Println("Destination: ", reflect.TypeOf(i))
		return err
	}

	return nil
}
