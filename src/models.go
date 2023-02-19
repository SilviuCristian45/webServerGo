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
