package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

var carQueue Queue = Queue{data: []Car{}}

func processCars(jsonFile string) ([]byte, error) {
	// read file
	data, err := os.ReadFile(fmt.Sprintf("../%s.json", "cars"))
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

	isSorted := req.URL.Query().Get("isSorted")

	if isSorted == "true" {
		sort.Slice(cars, func(i, j int) bool {
			return cars[i].YearReleased < cars[j].YearReleased
		})
	}

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
	os.WriteFile("../cars.json", carsBytes, fs.ModeAppend)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func getCarImage(w http.ResponseWriter, req *http.Request) {
	imageName := req.URL.Query().Get("image")
	fileBytes, err := os.ReadFile(fmt.Sprintf("../images/%s", imageName))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}

func getCarByYear(w http.ResponseWriter, req *http.Request) {
	year := req.URL.Query().Get("year")
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
	var yearCars = map[int]Cars{}
	for i := 0; i < len(cars); i++ {
		yearReleased := int(cars[i].YearReleased)
		yearCars[yearReleased] = append(yearCars[yearReleased], cars[i])
	}

	w.Header().Set("Content-Type", "application/json")
	result, err := strconv.ParseUint(year, 10, 32)
	json.NewEncoder(w).Encode(yearCars[int(result)])
}

func handleRoutes() {
	http.HandleFunc("/", getCars)
	http.HandleFunc("/addCar", addCar)
	http.HandleFunc("/files", getCarImage)
	http.HandleFunc("/cars", getCarByYear)
}

func main() {

	handleRoutes()
	http.ListenAndServe(":8090", nil)

}
