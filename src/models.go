// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cars, err := UnmarshalCars(bytes)
//    bytes, err = cars.Marshal()

package main

import "encoding/json"

type Cars []Car

func UnmarshalCars(data []byte) (Cars, error) {
	var r Cars
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Cars) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Car struct {
	Brand        string   `json:"brand"`
	Model        string   `json:"model"`
	YearReleased int64    `json:"yearReleased"`
	Options      []Option `json:"options"`
}

type Option struct {
	Name string `json:"name"`
	Cost int64  `json:"cost"`
}
