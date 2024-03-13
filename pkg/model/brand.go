package model

// brand represents brand information
type Brand struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`
}
