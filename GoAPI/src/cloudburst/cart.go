package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

func cart(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		addToCart(w, r)
	} else if r.Method == "PUT" {
		updateCart(w, r)
	} else if r.Method == "GET" {
		viewCart(w, r)
	} else if r.Method == "DELETE" {
		deleteCart(w, r)
	}
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
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if debug { fmt.Printf("%+v\n", cart) }

	if cart.Id == "" {
		http.Error(w, "User ID is not sent", 500)
		return
	}

	if cart.RestaurantId == 0{
		http.Error(w, "Restaurant ID is not sent", 500)
		return
	}

	output, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if insertObjects("cart", cart.Id, output) == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func updateCart(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var cart Cart
	err = json.Unmarshal(b, &cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if debug { fmt.Printf("%+v\n", cart) }

	if cart.Id == "" {
		http.Error(w, "User ID is not sent", 500)
		return
	}

	if cart.RestaurantId == 0 {
		http.Error(w, "Restaurant ID is not sent", 500)
		return
	}

	output, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	resp, nil := updateObjects("cart", cart.Id, output)
	if err != nil {
		log.Println("[RIAK DEBUG] " + err.Error())
	}

	w.Write(resp.Values[0].Value)
}

func viewCart(w http.ResponseWriter, r *http.Request) {
	//unmarshall
	var userid string
	userid = r.Header.Get("id")

	if debug { fmt.Println("cart id is :", userid) }

	resp, err := queryObjects("cart", userid)
	if err != nil {
		log.Println("[RIAK DEBUG] " + err.Error())
	}

	w.Write(resp.Values[0].Value)
}

func deleteCart(w http.ResponseWriter, r *http.Request) {
	//unmarshall
	var cartid string
	cartid = r.Header.Get("id")

	if debug { fmt.Println("cart id is :", cartid) }

	err := deleteObjects("cart", cartid)
	if err != nil {
		log.Println("[RIAK DEBUG] " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}