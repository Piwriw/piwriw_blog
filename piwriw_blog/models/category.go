package models

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	*Time
}