package transponster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/gorilla/mux"
)

func (io IO) logIncoming() {
	log.Printf("INCOMING[%v]: %v\n", io.id, io.R)
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
		log.Println("Struct: ", reflect.TypeOf(i))
		return err
	}

	return nil
}

func (io IO) Params() map[string]string {
	return mux.Vars(io.R)
}

func (io IO) ContextInt(key string) (*int, error) {
	value := io.R.Context().Value(key)

	number, ok := value.(int)
	if !ok {
		return nil, fmt.Errorf("value in context was not an int")
	}

	return &number, nil
}

func (io IO) ContextStrSlice(key string) (*[]string, error) {
	value := io.R.Context().Value(key)

	slice, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("value in context was not a string slice")
	}

	return &slice, nil
}
