package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
)

var carQueue Queue = Queue{data: []Car{}}

func processCars(jsonFile string) ([]byte, error) {
	// read file
	data, err := ioutil.ReadFile(fmt.Sprintf("../%s.json", "cars"))
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

	fmt.Println(carQueue)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func addCar(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var newCar Car
	errt := decoder.Decode(&newCar)
	if errt != nil {
		panic(errt)
	}
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
	cars = append(cars, newCar)
	carsBytes, _ := json.Marshal(cars)
	ioutil.WriteFile("../cars.json", carsBytes, fs.ModeAppend)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func handleRoutes() {
	http.HandleFunc("/", getCars)
	http.HandleFunc("/addCar", addCar)
}

func main() {

	handleRoutes()
	http.ListenAndServe(":8090", nil)

}
