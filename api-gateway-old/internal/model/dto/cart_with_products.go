package dto

// CartWithProducts represents a cart with product details
type CartWithProducts struct {
	Cart      interface{}     `json:"cart"`
	Products  []interface{} `json:"products"`
}