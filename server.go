package main

import (
	"fmt"
	"net"
	"encoding/json"

	"github.com/HouzuoGuo/tiedot/db"
	//"github.com/HouzuoGuo/tiedot/dberr"

)

//https://golang.org/src/net/rpc/jsonrpc/client.go
//http://blog.golang.org/json-and-go
//https://golang.org/src/encoding/json/example_test.go
//http://json.org/

type Request struct{
	Method string `json:"method"` 
	Params interface{} `json:"params`
	Id interface{} `json:"id"`
}

/*
type DICT3 struct{
	Key string
	Relationship string
	Value interface{}
}
*/



func handleConnection(conn net.Conn, triplets *db.Col) {
	dec := json.NewDecoder(conn)
	req := new(Request)
	dec.Decode(&req)
	fmt.Println()
	fmt.Printf("Received : %+v", req)
	fmt.Println()
	
	// Switch to see what method to call
	
	switch m.Method {
	case "lookup" :
		lookup(req, triplets)
	case "insert" :
		insert(req, triplets)
	case "insertOrUpdate":
		insertOrUpdate(req)
	case "delete" :
		deletekey(req)
	case "listKeys" :
		listKeys(req)
	case "listIDs" :
		listIDs(req)
	case "shutdown" :
		shutdown(req)
	}
	
}

func lookup(m *Request, triplets *db.Col){
	fmt.Printf("Looking up %v", m)

}


func insert(m *Request, triplets *db.Col){
	fmt.Printf("Inserting %v", m)
	
	/*docID, err := feeds.Insert(map[string]interface{}{
		"name": "Go 1.2 is released",
		"url":  "golang.org"})
	if err != nil {
		panic(err)
	}*/
}

func insertOrUpdate(m *Request){
	fmt.Printf("%v", m)
}

func deletekey(m *Request){
	fmt.Printf("%v", m)
}

func listKeys(m *Request){
	fmt.Printf("%v", m)
}

func listIDs(m *Request){
	fmt.Printf("%v", m)
}

func shutdown(m *Request){
	fmt.Printf("%v", m)
}




func main() {
	fmt.Println("start")

	//DB code
	
	fmt.Println("Intalizing DB")
	myDBDir := "tiedotDB"

	// (Create if not exist) open a database
	myDB, err := db.OpenDB(myDBDir)
	if err != nil {
		panic(err)
	}

	if err := myDB.Create("Triplets"); err != nil {
		//panic(err)
		fmt.Printf(err.Error())
	}
	
	// Scrub (repair and compact) "Feeds"
	/* if err := myDB.Scrub("Feeds"); err != nil {
		panic(err)
	}
        */

	// Start using a collection (the reference is valid until DB schema changes or Scrub is carried out)
	triplets := myDB.Use("Triplets")


	
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
		go handleConnection(conn, triplets) // a goroutine handles conn so that the loop can accept other connections
	}
}
