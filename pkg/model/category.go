package model

// category represents category information
type Category struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`
}
