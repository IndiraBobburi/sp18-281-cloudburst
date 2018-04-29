package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var name string
	name = "Indy"
	fmt.Fprintf(w, "Hi there, I love %s!",name)
}

func main() {
	http.HandleFunc("/hi", handler)
	http.HandleFunc("/getRestaurants", getRestaurants)
	http.HandleFunc("/addToCart", addToCart)
	http.HandleFunc("/viewCart", viewCart)
	http.HandleFunc("/order", order)
	http.HandleFunc("/orderStatus", orderStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var pincode Pincode
	err = json.Unmarshal(b, &pincode)
	fmt.Printf("%+v\n", pincode)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//TODO: Get the restaurant list from the db

	restaurants := Restaurant{
		Id:     1,
		Name:   "Ulavacharu",
		Address:	"xyz",
		Phone:	"320-234-2384",
		Menu: []Item {
			Item{
				Id:	1,
				Price: 2.0,
				Name: "Idly",
				Description: "South Indian",
			},
			Item{
				Id:	1,
				Price: 2.0,
				Name: "Idly",
				Description: "South Indian",
			},
		},
	}

	output, err := json.Marshal(restaurants)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

}


func addToCart(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var cart Cart
	err = json.Unmarshal(b, &cart)

	if cart.Id == 0{
		http.Error(w, "ID is not sent", 500)
		return
	}

	if cart.RestaurantId == 0{
		http.Error(w, "Restaurant ID is not sent", 500)
		return
	}

	if cart.Items == nil{
		http.Error(w, "Items null", 500)
		return
	}

	//TODO: Loop through items and print error by checking id and qty


	fmt.Printf("%+v\n", cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//TODO: Do the db insertion

	output, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)


}

func orderStatus(w http.ResponseWriter, r *http.Request){
	orderId := r.URL.Query().Get("orderId")
	if orderId != "" {
		//TODO: query the db and return the order status
		fmt.Println("Order id is ",orderId)
	}
	//orderstatus := "Order id"+orderId +"is being processed"
	orderstatus:= fmt.Sprintf( "Order id %v is being processed", orderId)
	output, err := json.Marshal(orderstatus)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}


func viewCart(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var cart Cart
	err = json.Unmarshal(b, &cart)

	if cart.Id == 0{
		http.Error(w, "ID is not sent", 500)
		return
	}

	if cart.RestaurantId == 0{
		http.Error(w, "Restaurant ID is not sent", 500)
		return
	}

	if cart.Items == nil{
		http.Error(w, "Items null", 500)
		return
	}

	//TODO: Loop through items and print error by checking id and qty


	fmt.Printf("%+v\n", cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//TODO: Do the db insertion

	output, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)


}

func order(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var cart Cart
	err = json.Unmarshal(b, &cart)

	if cart.Id == 0{
		http.Error(w, "ID is not sent", 500)
		return
	}

	if cart.RestaurantId == 0{
		http.Error(w, "Restaurant ID is not sent", 500)
		return
	}

	if cart.Items == nil{
		http.Error(w, "Items null", 500)
		return
	}

	//TODO: Loop through items and print error by checking id and qty


	fmt.Printf("%+v\n", cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//TODO: Do the db insertion

	output, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)


}