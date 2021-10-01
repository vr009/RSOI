package models

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Work    string `json:"work"`
	Address string `json:"address"`
}

type PersonRequest struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}
