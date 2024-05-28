package models

type Cart struct {
	ID    int        `json:"id"`
	Items []CartItem `json:"items"`
}
