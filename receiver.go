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

type detail struct {
	method string
	url    string
	ua     string
	src    string
}

func (io *IO) requestDetail() string {
	return fmt.Sprintf("%+v", detail{
		method: io.R.Method,
		url:    io.R.URL.Path,
		ua:     io.R.UserAgent(),
		src:    io.R.RemoteAddr,
	})
}

func (io IO) logIncoming() {
	log.Printf("INCOMING[%v]: %v\n", io.id, io.requestDetail())
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

func (io IO) Params(param string) (*string, error) {
	params := mux.Vars(io.R)

	val, ok := params[param]
	if !ok {
		return nil, fmt.Errorf("param does not exist")
	}

	return &val, nil
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
