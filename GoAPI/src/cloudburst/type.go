package main


type Pincode struct {
	Id  uint64 `json:"pincode"`
}

type Restaurant struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}

type RestaurantList struct {
	Restaurants []Restaurant `json:"restaurantlist"`
}

type Item struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Description string `json:"description"`
}

type ItemList struct {
	Menu []Item `json:"menu"`
}

type Cart struct{
	UserId uint64 `json:"userid"`
	Id string `json:"id"`
	RestaurantId uint64 `json:"restaurantId"`
	Items []CartItem `json:"items"`
}

type CartItem struct {
	Id uint64 `json:"id"` //item id
	Quantity uint8 `json:"quantity"`
}

type Order struct {
	Id             	string
	OrderStatus 	string
}