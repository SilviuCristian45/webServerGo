package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func processCars(jsonFile string) ([]byte, error) {
	// read file
	data, err := ioutil.ReadFile(fmt.Sprintf("./%s.json", "cars"))
	if err != nil {
		log.Fatalf("Error when reading json from db %s", err)
		return nil, err
	}
	return data, nil
}

func getCars(w http.ResponseWriter, req *http.Request) {
	var cars Cars
	var data, error = processCars("cars")
	if error != nil {
		w.WriteHeader(500)
	}
	var err = json.Unmarshal(data, &cars)
	if err != nil {
		log.Fatalf("Error when parsing json to Cars model %s", err)
		w.WriteHeader(500)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func handleRoutes() {
	http.HandleFunc("/", getCars)
}

func main() {

	handleRoutes()
	http.ListenAndServe(":8090", nil)

}
