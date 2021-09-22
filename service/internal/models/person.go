package models

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Work    string `json:"work"`
	Address string `json:"address"`
}
