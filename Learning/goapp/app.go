package main


import (
	"net/http"
	"../../kalikalyandash/counter-burger/src/github.com/codegangsta/negroni"
	"../../kalikalyandash/counter-burger/src/github.com/gorilla/mux"
	"../../kalikalyandash/counter-burger/src/github.com/unrolled/render"
	"../../kalikalyandash/counter-burger/src/gopkg.in/mgo.v2"
	"../../kalikalyandash/counter-burger/src/gopkg.in/mgo.v2/bson"
	"../../kalikalyandash/counter-burger/src/github.com/satori/go.uuid"
	"log"
	"fmt"
	"encoding/json"
)

// MongoDB Config
var mongodb_server = "127.0.0.1:27017"
var mongodb_database = "cmpe281"
var mongodb_collection = "orders"

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/orders", getOrdersHandler(formatter)).Methods("GET")
	mx.HandleFunc("/orders/{id}", getOrderHandler(formatter)).Methods("GET")
	mx.HandleFunc("/orders", postOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/order/{id}", updateOrderHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/order/{id}", deleteOrderHandler(formatter)).Methods("DELETE")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hiiii:")
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}


// API Machine Handler
func getOrdersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var result[] bson.M
		err = c.Find(bson.M{}).All(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Orders:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}

func postOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Order Post:" )
		uuid := uuid.NewV4()
		fmt.Println(uuid.String())
		fmt.Println(req.Body)

		var order Order
		_ = json.NewDecoder(req.Body).Decode(&order)
		order.OrderId = uuid.String()
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		err = c.Insert(order)
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, order)
	}
}



