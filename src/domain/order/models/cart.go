package models

type CartItem struct {
	ProductID    string `json:"product_id"`    // required
	Quantity     int    `json:"quantity"`      // required
	ProductName  string `json:"product_name"`  // optional
	ProductPrice int    `json:"product_price"` // optional
}
type Cart struct {
	RestaurantID string `json:"restaurant_id"`
	Instruction  string `json:"instruction"`
	Restaurant   struct {
		Name     string `json:"name"`
		ImageUrl string `json:"image_url"`
	} `json:"restaurant"`
	Delivery struct {
		AddressID string `json:"address_id"`
		Time      string `json:"time"`
		Notes     string `json:"notes"`
	} `json:"delivery"`
	Payment struct {
		Method        string `json:"method"`
		CardID        string `json:"card_id"`
		DeliveryPrice int    `json:"delivery_price"`
		TotalPrice    int    `json:"total_price"`
	} `json:"payment"`
	Items []*CartItem `json:"items"`
}
