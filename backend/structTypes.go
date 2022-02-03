package main

type search struct {
	ID       string `json:"ID"`
	Name     string `json: "Name"`
	Position string `json: "Position"`
}

type city struct {
	City_name string `json: "city_name"`
	City_id   int    `json: "city_id"`
}

type vendor struct {
	First_name string `json: "first_name"`
	Last_name  string `json: "last_name"`
	Phone      int    `json: "phn"`
}
