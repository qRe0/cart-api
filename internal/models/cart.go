package models

type Cart struct {
	Entity
	Items []CartItem `json:"items"`
}
