package main


type Pincode struct {
	Id  uint64 `json:"pincode"`
}

type Restaurant struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Phone string `json:"phone"`
	Menu []Item `json:"menu"`

}

type Item struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Description string `json:"description"`
}

type Cart struct{
	Id uint64 `json:"id"`
	RestaurantId uint64 `json:"restaurantId"`
	Items []CartItem `json:"items"`
}

type CartItem struct {
	Id uint64 `json:"id"`
	Quantity uint8 `json:"quantity"`
}