package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/satori/go.uuid"
	"fmt"
)

func order(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createOrder(w,r)
	} else if r.Method == "PUT" {
		updateOrder(w,r)
	} else if r.Method == "GET" {
		getOrder(w,r)
	}
}

func createOrder(w http.ResponseWriter, r *http.Request){
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var order Order
	err = json.Unmarshal(b, &order)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if order.Id == "" {
		uuid, _ := uuid.NewV4()
		order.Id = uuid.String()
	}

	if order.UserId == "" {
		http.Error(w, "User ID is not sent", 500)
		return
	}

	if order.RestaurantId == 0 {
		http.Error(w, "Restaurant ID is not sent", 500)
		return
	}

	order.OrderStatus = "Order Placed"

	if debug { fmt.Printf("%+v\n", order) }

	output, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if insertObjects("orders", order.Id, output) == nil {
		//delete cart
		err := deleteObjects("cart", order.UserId)
		if err != nil {
			log.Println("[RIAK DEBUG] " + err.Error())
		}

		w.WriteHeader(http.StatusOK)
		w.Write(output)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func updateOrder(w http.ResponseWriter, r *http.Request){
	var orderid string
	orderid = r.URL.Query().Get("orderid")

	if debug { log.Println(orderid) }

	if orderid != "" {
		resp, err := queryObjects("orders", orderid)
		if err != nil {
			log.Println("[RIAK DEBUG] " + err.Error())
		}

		var order Order
		err = json.Unmarshal(resp.Values[0].Value, &order)
		order.OrderStatus = "Order Processed"

		output, err := json.Marshal(order)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		newrsp, err := updateObjects("orders", orderid, []byte(output))
		if err != nil {
			log.Println("[RIAK DEBUG] " + err.Error())
		}

		w.Write(newrsp.Values[0].Value)
	}
}

func getOrder(w http.ResponseWriter, r *http.Request){
	var orderid string
	orderid = r.URL.Query().Get("orderid")

	if orderid != "" {
		resp, err := queryObjects("orders", orderid)
		if err != nil {
			log.Println("[RIAK DEBUG] " + err.Error())
		}

		w.Write(resp.Values[0].Value)
	}
}

func getOrders(w http.ResponseWriter, r *http.Request){

}