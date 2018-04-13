package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	//"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/satori/go.uuid"
	//"gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
	"github.com/gocql/gocql"
	//"github.com/GetStream/stream-go"
)
var Session *gocql.Session
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
func init() {
  var err error
   cluster := gocql.NewCluster("localhost")

  cluster.Keyspace = "shop"
  Session, err = cluster.CreateSession()
  if err != nil {
    panic(err)
  }
  fmt.Println("Cassandra is up and running")
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/users", getAllUsersHandler(formatter)).Methods("GET")
	mx.HandleFunc("/user/{id}", getUserByIdHandler(formatter)).Methods("GET")
	mx.HandleFunc("/user/{id}", deleteUserByIdHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/user/{id}", updateUserByIdHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/users", addNewUserHandler(formatter)).Methods("POST")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}


func pingHandler(formatter *render.Render) http.HandlerFunc {

}
func updateUserByIdHandler(formatter *render.Render) http.HandlerFunc {

}
func deleteUserByIdHandler(formatter *render.Render) http.HandlerFunc {

}


func getAllUsersHandler(formatter *render.Render) http.HandlerFunc {

}

func getUserByIdHandler(formatter *render.Render) http.HandlerFunc {

}

func addNewUserHandler(formatter *render.Render) http.HandlerFunc {

}
