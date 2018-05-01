package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/satori/go.uuid"
	"github.com/basho/riak-go-client"
	"encoding/json"
)

var debug = true
/*var server1 = "http://localhost:9000"
var server2 = "http://localhost:9001"
var server3 = "http://localhost:9002"*/

var s1 = "localhost:8002"
var s2 = "localhost:8003"
var s3 = "localhost:8004"

var cluster *riak.Cluster
//var c *riak.Client

/*type Client struct {
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
}*/



func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there! Welcome to goBurger")
}

func initCluster(){
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

func init(){
	initCluster()
	/*var err error

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
	}*/
}

func main() {
	http.HandleFunc("/hi", handler)
	http.HandleFunc("/getRestaurants", getRestaurants)
	http.HandleFunc("/getMenu", getMenu)
	http.HandleFunc("/addToCart", addToCart)
	http.HandleFunc("/viewCart", viewCart)
	http.HandleFunc("/order", order)
	//http.HandleFunc("/try", try)
	log.Fatal(http.ListenAndServe(":8080", nil))
	defer func() {
		if err := cluster.Stop(); err != nil {
			fmt.Println(err.Error())
		}
	}()
}

/*func try(w http.ResponseWriter, r *http.Request) {
	err := insertingObjects()
	if err == nil {
		w.WriteHeader(http.StatusOK)
	}else {
		log.Println(err)
	}
}*/

func insertingObjects(bucket string, key string, body []byte) error {
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

func queryingObjects(bucket string, key string ) (*riak.FetchValueResponse, error) {
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

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pincode := r.URL.Query().Get("pincode")

		if pincode != "" {
			/*c := NewClient(server1)
			resp, err := c.Get(c.Endpoint + "/buckets/restaurants/keys/"+pincode )
			if err != nil {
				fmt.Println("[RIAK DEBUG] " + err.Error())
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			w.Write(body)*/

			resp, err := queryingObjects("restaurants", pincode)
			if err != nil {
				fmt.Println("[RIAK DEBUG] " + err.Error())
			}

			w.Write(resp.Values[0].Value)
		}
	}
}

func getMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		/*c := NewClient(server1)
		resp, err := c.Get(c.Endpoint + "/buckets/restaurants/keys/menu" )

		if err != nil {
			fmt.Println("[RIAK DEBUG] " + err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		w.Header().Set("content-type", "application/json")
		w.Write(body)*/
		resp, err := queryingObjects("restaurants", "menu")
		if err != nil {
			fmt.Println("[RIAK DEBUG] " + err.Error())
		}

		w.Write(resp.Values[0].Value)
	}
}

func addToCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

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

		//val := bytes.NewBuffer(output)

		/*c := NewClient(server1)
		resp, err := c.Post(c.Endpoint + "/buckets/cart/keys/"+cart.Id+"?returnbody=true",
			"application/json", val )

		if err != nil {
			fmt.Println("[RIAK DEBUG] " + err.Error())
		}
		defer resp.Body.Close()*/
		if insertingObjects("cart",cart.Id, output) == nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}


func viewCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		//unmarshall
		var cartid string
		cartid = r.Header.Get("cartid")

		if debug { fmt.Println("cart id is :", cartid) }

		/*c := NewClient(server1)
		resp, err := c.Get(c.Endpoint + "/buckets/cart/keys/"+ cartid )

		if err != nil {
			fmt.Println("[RIAK DEBUG] " + err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		w.Write(body)*/

		resp, err := queryingObjects("cart", cartid)
		if err != nil {
			fmt.Println("[RIAK DEBUG] " + err.Error())
		}

		w.Write(resp.Values[0].Value)
	}
}

func order(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//c := NewClient(server1)
		uuid, _ := uuid.NewV4()
		value := "Order Placed"

		reqbody := "{\"Id\": \"" +
			uuid.String() +
			"\",\"OrderStatus\": \"" +
			value +
			"\"}"

		/*resp, err := c.Post(c.Endpoint + "/buckets/orders/keys/"+ uuid.String() +"?returnbody=true",
			"application/json", strings.NewReader(reqbody) )
		if err != nil {
			fmt.Println("[RIAK DEBUG] " + err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)*/

		if insertingObjects("orders", uuid.String(), []byte(reqbody)) == nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(reqbody))
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
	if r.Method == "PUT" {
		var orderid string
		orderid = r.URL.Query().Get("orderId")
		if orderid != "" {

			reqbody := "{\"Id\": \"" +
				orderid +
				"\",\"OrderStatus\": \"" +
				"Order processed" +
				"\"}"
			resp, err := updateObjects("orders", orderid, []byte(reqbody))
			if err != nil {
				fmt.Println("[RIAK DEBUG] " + err.Error())
			}

			w.Write(resp.Values[0].Value)
		}
	}
	if r.Method == "GET" {
		var orderid string
		orderid = r.URL.Query().Get("orderId")
		if orderid != "" {
			/*c := NewClient(server1)
			resp, err := c.Get(c.Endpoint + "/buckets/cart/keys/"+ orderid )
			if err != nil {
				fmt.Println("[RIAK DEBUG] " + err.Error())
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			w.Write(body)*/

			resp, err := queryingObjects("orders", orderid)
			if err != nil {
				fmt.Println("[RIAK DEBUG] " + err.Error())
			}

			w.Write(resp.Values[0].Value)
		}
	}
}