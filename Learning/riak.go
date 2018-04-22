package main

import (
	"fmt"
	"os"
	"reflect"
	riak "github.com/basho/riak-go-client"
)

func main() {
	var err error
	nodeOpts1 := &riak.NodeOptions{
		RemoteAddress: "204.236.157.206:8087",
	}

	nodeOpts2 := &riak.NodeOptions{
		RemoteAddress: "52.52.69.30:8087",
	}

	nodeOpts3 := &riak.NodeOptions{
		RemoteAddress: "54.153.118.197:8087",
	}

	var node1 *riak.Node
	var node2 *riak.Node
	var node3 *riak.Node

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

	cluster, err := riak.NewCluster(opts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer func() {
		if err := cluster.Stop(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	if err := cluster.Start(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	obj := &riak.Object{
	    ContentType:     "application/json",
	    Charset:         "utf-8",
	    ContentEncoding: "utf-8",
	    Value:           []byte("{'user':'data'}"),
	}

	cmd, err := riak.NewStoreValueCommandBuilder().
		WithBucket("testBucketName").
		WithContent(obj).
		Build()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := cluster.Execute(cmd); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	svc := cmd.(*riak.StoreValueCommand)
	rsp := svc.Response
	fmt.Println(reflect.TypeOf(rsp))
	fmt.Println(rsp.GeneratedKey)
}
