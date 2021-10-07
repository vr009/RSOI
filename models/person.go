package models

type Person struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Work    string `json:"work"`
	Address string `json:"address"`
}
