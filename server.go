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


func handleConnection(conn net.Conn, triplets *db.Col) {
	dec := json.NewDecoder(conn)
	req := new(Request)
	dec.Decode(&req)
	fmt.Println()
	fmt.Printf("Received : %+v", req)
	fmt.Println()
	
	// Switch to see what method to call
	
	switch req.Method {
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

func lookup(req *Request, triplets *db.Col){
	fmt.Printf("Looking up %v", req)

}


func insert(req *Request, triplets *db.Col){
	fmt.Printf("InSerting %v", req)

	p := req.Params
	arr := p.([]interface{})
	
	key := arr[0].(string)
	rel := arr[1].(string)
	val := arr[2]

	fmt.Println(key, rel, val)

	/*
	doc := make(map[string]interface{})
	reldoc := make(map[string]interface{})
	reldoc[rel] = value
	doc[key] = reldoc
	*/
	// Inserting document into DB as ["key,relationship"] : value

	
	docID, err := triplets.Insert(map[string]interface{}{
		"key": key,
		"rel": rel,
		"val": val})
	if err != nil {
		panic(err)
	}
	fmt.Println(docID)

	// Create indexes here??
	// TODO: Do not create index if it already exists?
	/*if err := triplets.Index([]string{"key", "rel"}); err != nil {
		//panic(err)
		fmt.Printf(err.Error())
	}
        */
	
        
	
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

func testretrive(triplets *db.Col){
	var query interface{}
	//json.Unmarshal([]byte(`[{"eq": "keyA", "in": ["key"]}, {"eq": "relA", "in": ["rel"]}]`), &query)
	json.Unmarshal([]byte(`{"eq": "keyA", "in": ["key"]}`), &query)
	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalQuery(query, triplets, &queryResult); err != nil {
		panic(err)
	}

	// Query result are document IDs
	for id := range queryResult {
		// To get query result document, simply read it
		readBack, err := triplets.Read(id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Query returned document %v\n", readBack)
		fmt.Println(readBack["key"])
		fmt.Println(readBack["rel"])
		fmt.Println(readBack["val"])
	}
	

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
	/*if err := myDB.Scrub("Triplets"); err != nil {
		panic(err)
	}
        */

	// Start using a collection (the reference is valid until DB schema changes or Scrub is carried out)
	triplets := myDB.Use("Triplets")

	// Remove index
	/*
	if err := triplets.Unindex([]string{"key", "rel"}); err != nil {
		panic(err)
	}
	*/

	// Create indexes here??
	// TODO: Do not create index if it already exists?
	if err := triplets.Index([]string{"key"}); err != nil {
		//panic(err)
		fmt.Printf(err.Error())
	}
        

	
	for _, path := range triplets.AllIndexes() {
		fmt.Printf("I have an index on path %v\n", path)
	}

	testretrive(triplets)
	
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
