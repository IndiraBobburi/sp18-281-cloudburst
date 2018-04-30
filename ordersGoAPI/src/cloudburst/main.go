package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"time"
	"bytes"
	"github.com/satori/go.uuid"
	"strings"
)

var debug = true
var server1 = "http://localhost:9000"
var server2 = "http://localhost:9001"
var server3 = "http://localhost:9002"

type Client struct {
	Endpoint string
	*http.Client
}

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

func NewClient(server string) *Client {
	return &Client{
		Endpoint:  	server,
		Client: 	&http.Client{Transport: tr},
	}
}

func (c *Client) Ping() (string, error) {
	resp, err := c.Get(c.Endpoint + "/ping" )
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return "Ping Error!", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if debug { fmt.Println("[RIAK DEBUG] GET: " + c.Endpoint + "/ping => " + string(body)) }
	return string(body), nil
}

//var c *riak.Client

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, Welcome to goBurger")
}

func init(){
	var err error

	c1 := NewClient(server1)
	msg, err := c1.Ping( )
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Riak Ping Server1: ", msg)
	}

	c2 := NewClient(server2)
	msg, err = c2.Ping( )
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Riak Ping Server2: ", msg)
	}

	c3 := NewClient(server3)
	msg, err = c3.Ping( )
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Riak Ping Server3: ", msg)
	}
}

func main() {
	http.HandleFunc("/hi", handler)
	http.HandleFunc("/getRestaurants", getRestaurants)
	http.HandleFunc("/getMenu", getMenu)
	http.HandleFunc("/addToCart", addToCart)
	http.HandleFunc("/viewCart", viewCart)
	http.HandleFunc("/order", order)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	pincode := r.URL.Query().Get("pincode")

	if pincode != "" {
		c := NewClient(server1)
		resp, err := c.Get(c.Endpoint + "/buckets/restaurants/keys/"+pincode )
		if err != nil {
			fmt.Println("[RIAK DEBUG] " + err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		w.Write(body)
	}
}

func getMenu(w http.ResponseWriter, r *http.Request) {
	c := NewClient(server1)
	resp, err := c.Get(c.Endpoint + "/buckets/restaurants/keys/menu" )

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	w.Header().Set("content-type", "application/json")
	w.Write(body)
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

	fmt.Printf("%+v\n", cart)

	if cart.Id == "" {
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


	output, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	val := bytes.NewBuffer(output)

	c := NewClient(server1)
	resp, err := c.Post(c.Endpoint + "/buckets/cart/keys/"+cart.Id+"?returnbody=true",
		"application/json", val )

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
	}
	defer resp.Body.Close()

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}


func viewCart(w http.ResponseWriter, r *http.Request) {
	//unmarshall
	var cartid string
	cartid = r.Header.Get("cartid")

	fmt.Println("cart id is :", cartid)

	c := NewClient(server1)
	resp, err := c.Get(c.Endpoint + "/buckets/cart/keys/"+ cartid )

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	w.Write(body)
}

func order(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c := NewClient(server1)
		uuid, _ := uuid.NewV4()
		value := "Order Placed"

		reqbody := "{\"Id\": \"" +
			uuid.String() +
			"\",\"OrderStatus\": \"" +
			value +
			"\"}"

		resp, err := c.Post(c.Endpoint + "/buckets/orders/keys/"+ uuid.String() +"?returnbody=true",
			"application/json", strings.NewReader(reqbody) )
		if err != nil {
			fmt.Println("[RIAK DEBUG] " + err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
	if r.Method == "GET" {
		var orderid string
		orderid = r.URL.Query().Get("orderId")
		if orderid != "" {
			c := NewClient(server1)
			resp, err := c.Get(c.Endpoint + "/buckets/cart/keys/"+ orderid )
			if err != nil {
				fmt.Println("[RIAK DEBUG] " + err.Error())
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			w.Write(body)
		}
	}
}