package main

import (
	"net/http"
	"log"
)

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "GET" {
		pincode := r.URL.Query().Get("pincode")

		if pincode != "" {
			resp, err := queryObjects("restaurants", pincode)
			if err != nil {
				log.Println("[RIAK DEBUG] " + err.Error())
			}

			w.Write(resp.Values[0].Value)
		}
	}
}

func getMenu(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "GET" {
		resp, err := queryObjects("restaurants", "menu")
		if err != nil {
			log.Println("[RIAK DEBUG] " + err.Error())
		}

		w.Write(resp.Values[0].Value)
	}
}