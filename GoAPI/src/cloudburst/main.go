package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/basho/riak-go-client"
)

var debug = true

//connect to tcp ports for cluster
var s1 = "54.183.106.118:8087"   //"localhost:8002"
var s2 = "13.57.3.195:8087"      //"localhost:8003"
var s3 = "54.153.107.186:8087"   //"localhost:8004"

var cluster *riak.Cluster

func handler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Hi there! Welcome to goBurger")
}

func init(){
	nodeOpts1 := &riak.NodeOptions{
		RemoteAddress: s1,
	}

	nodeOpts2 := &riak.NodeOptions{
		RemoteAddress: s2,
	}

	nodeOpts3 := &riak.NodeOptions{
		RemoteAddress: s3,
	}

	var node1 *riak.Node
	var node2 *riak.Node
	var node3 *riak.Node
	var err error

	if node1, err = riak.NewNode(nodeOpts1); err != nil {
		fmt.Println(err.Error())
	}

	if node2, err = riak.NewNode(nodeOpts2); err != nil {
		fmt.Println(err.Error())
	}

	if node3, err = riak.NewNode(nodeOpts3); err != nil {
		fmt.Println(err.Error())
	}

	nodes := []*riak.Node{node1, node2, node3}
	opts := &riak.ClusterOptions{
		Nodes: nodes,
	}

	log.Println( nodes )

	cluster, err = riak.NewCluster(opts)
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := cluster.Start(); err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	http.HandleFunc("/hi", handler)
	http.HandleFunc("/getRestaurants", getRestaurants)
	http.HandleFunc("/getMenu", getMenu)
	http.HandleFunc("/cart", cart)
	http.HandleFunc("/order", order)
	http.HandleFunc("/orders", getOrders)
	http.HandleFunc("/user", user)

	http.ListenAndServe(":8080", nil)

	defer func() {
		if err := cluster.Stop(); err != nil {
			log.Println(err.Error())
		}
	}()
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func insertObjects(bucket string, key string, body []byte) error {
	obj := &riak.Object{
		ContentType:     "application/json",
		Key:             key,
		Value:           body,
	}

	cmd, err := riak.NewStoreValueCommandBuilder().
		WithBucket(bucket).
		WithContent(obj).
		Build()
	if err != nil {
		return err
	}

	if err = cluster.Execute(cmd); err != nil {
		return err
	}

	return nil
}

func queryObjects(bucket string, key string ) (*riak.FetchValueResponse, error) {
	cmd, err := riak.NewFetchValueCommandBuilder().
		WithBucket(bucket).
		WithKey(key).
		Build()
	if err != nil {
		return nil, err
	}

	if err = cluster.Execute(cmd); err != nil {
		return nil, err
	}

	fvc := cmd.(*riak.FetchValueCommand)
	rsp := fvc.Response

	log.Println(string(rsp.Values[0].Value))
	return rsp, nil
}

func updateObjects(bucket string, key string, newval []byte) (*riak.FetchValueResponse, error) {
	cmd, err := riak.NewFetchValueCommandBuilder().
		WithBucket(bucket).
		WithKey(key).
		Build()
	if err != nil {
		return nil, err
	}

	if err = cluster.Execute(cmd); err != nil {
		return nil, err
	}

	fvc := cmd.(*riak.FetchValueCommand)
	rsp := fvc.Response

	if debug { log.Println(string(rsp.Values[0].Value)) }
	rsp.Values[0].Value = newval
	return rsp, nil
}

func deleteObjects(bucket string, key string) error{
	cmd, err := riak.NewDeleteValueCommandBuilder().
		WithBucket(bucket).
		WithKey(key).
		Build()
	if err != nil {
		return err
	}

	return cluster.Execute(cmd)
}
