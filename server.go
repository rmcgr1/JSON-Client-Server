package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"encoding/json"
	"os"
)

//https://golang.org/src/net/rpc/jsonrpc/client.go
//http://blog.golang.org/json-and-go
//https://golang.org/src/encoding/json/example_test.go
//http://json.org/

type P struct {
	M, N int64
}

type DICT3 struct{
	Key string
	Relationship string
	//Value interface{}
}

/*
func insert(d3 DICT3){

}
*/


func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	p := &P{}
	dec.Decode(p)
	fmt.Printf("Received : %+v", p)
}

func main() {
	fmt.Println("start")
	
	// Test JSON code
	m := DICT3{"TheKey", "The___Value"}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)

	// For unmarshall
	var message []DICT3

	err = json.Unmarshal(b, &message)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Printf("%+v", message)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
	}
}
