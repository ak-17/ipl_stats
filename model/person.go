package model

type Gender int

const (
	Male Gender = iota
	Female
)

type Person struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender Gender `json:"gender"`
}
