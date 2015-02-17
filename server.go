package main

import (
	"fmt"
	"net"
	"encoding/json"
	//"os"
)

//https://golang.org/src/net/rpc/jsonrpc/client.go
//http://blog.golang.org/json-and-go
//https://golang.org/src/encoding/json/example_test.go
//http://json.org/

type DICT3 struct{
	Key string `json:"method"`
	Relationship string `json:"params"`
	Id string `json:"id"`
}

/*
func insert(d3 DICT3){

}
*/


func handleConnection(conn net.Conn) {
	dec := json.NewDecoder(conn)
	m := new(DICT3)
	dec.Decode(&m)
	fmt.Printf("Received : %+v", m)

	// Switch to see what method to call
	
	switch m.Key {
	case "lookup" :
		lookup(m)
	case "insert" :
		insert(m)
	case "insertOrUpdate":
		insertOrUpdate(m)
	case "delete" :
		deletekey(m)
	case "listKeys" :
		listKeys(m)
	case "listIDs" :
		listIDs(m)
	case "shutdown" :
		shutdown(m)
	}
	
}

func lookup(m *DICT3){
	fmt.Printf("%v", m)
}


func insert(m *DICT3){
	fmt.Printf("%v", m)
}

func insertOrUpdate(m *DICT3){
	fmt.Printf("%v", m)
}

func deletekey(m *DICT3){
	fmt.Printf("%v", m)
}

func listKeys(m *DICT3){
	fmt.Printf("%v", m)
}

func listIDs(m *DICT3){
	fmt.Printf("%v", m)
}

func shutdown(m *DICT3){
	fmt.Printf("%v", m)
}




func main() {
	fmt.Println("start")
	
	// Test JSON code
	/*
        m := DICT3{"TheKey", "TheValue"}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)

	// For unmarshall
	message := new(DICT3)

	err = json.Unmarshal(b, &message)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Printf("%+v", message)
	//fmt.Println(m.Key)
        */


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
