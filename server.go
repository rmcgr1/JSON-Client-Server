package main

import (
	"fmt"
	"net"
	"encoding/json"
	"os"
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

type Response struct{
	Result interface{} `json:"result"`
	Id interface{} `json:"id"`
	Error interface{} `json:"error"`      // Error must be null if there was no error
}

func handleConnection(conn net.Conn, triplets *db.Col, myDB *db.DB) {
	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)
	req := new(Request)
	decoder.Decode(&req)
	fmt.Println()
	fmt.Printf("Received : %+v", req)
	fmt.Println()
	
	// Switch to see what method to call
	
	switch req.Method {
	case "lookup" :
		lookup(req, encoder, triplets)
	case "insert" :
		insert(req, encoder, triplets, false)
	case "insertOrUpdate":
		insert(req, encoder, triplets, true)
	case "delete" :
		delete(req, encoder, triplets)
	case "listKeys" :
		listKeys(encoder, triplets)
	case "listIDs" :
		listIDs(encoder, triplets)
	case "shutdown" :
		shutdown(myDB)
	}
	
}

func lookup(req *Request, encoder *json.Encoder, triplets *db.Col){
	fmt.Printf("Looking up %v", req)

	p := req.Params
	arr := p.([]interface{})
	
	key := arr[0].(string)
	rel := arr[1].(string)

	// See if there this key/val is already in DB
	queryResult := query_key_rel(key, rel, triplets)
	if len(queryResult) != 0 {
		for id := range queryResult {
			
			readBack, err := triplets.Read(id)
			if err != nil {
				panic(err)
			}
			
			val := readBack["val"].(map[string]interface {})
			fmt.Println(val)
			//fmt.Println(key + " " + rel + " has value " + val)
			//TODO get ID value
			nillslice := []int{}
			m := Response{val, "ChangeMeID", nillslice}
			encoder.Encode(m)
		}
		
	} else {
		// Key/rel not in DB
		fmt.Println("key/rel not in DB return null in result")
	}
	
}

func listKeys(encoder *json.Encoder, triplets *db.Col){
	//TODO make sure UNIQUE keys
	fmt.Println("Listing all unique keys")

	var query interface{}
	json.Unmarshal([]byte(`{"n": [{"has": ["key"]}, {"has": ["rel"]}]}`), &query)
	//json.Unmarshal([]byte(`{"eq": "keyA", "in": ["key"]}`), &query)
	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalQuery(query, triplets, &queryResult); err != nil {
		panic(err)
	}


	key_set := make(map[string]bool)
	// Query result are document IDs
	fmt.Println(queryResult)
	for id := range queryResult {

		readBack, err := triplets.Read(id)
		if err != nil {
			panic(err)
		}
		
		key_set[readBack["key"].(string)] = true
	}

	for i := range key_set{
		fmt.Println(i)
	}
	
	
}

func listIDs(encoder *json.Encoder, triplets *db.Col){
	//TODO make sure UNIQUE keys
	fmt.Println("Listing all unique IDs")

	var query interface{}
	json.Unmarshal([]byte(`{"n": [{"has": ["key"]}, {"has": ["rel"]}]}`), &query)
	//json.Unmarshal([]byte(`{"eq": "keyA", "in": ["key"]}`), &query)
	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalQuery(query, triplets, &queryResult); err != nil {
		panic(err)
	}

	fmt.Println(queryResult)

	id_set := make(map[[2]string]bool)
	// Query result are document IDs
	for id := range queryResult {

		readBack, err := triplets.Read(id)
		if err != nil {
			panic(err)
		}
		fmt.Println(readBack)
		id_set[[2]string{readBack["key"].(string), readBack["rel"].(string)}] = true 
	}

	for i := range id_set{
		fmt.Println(i)
	}
	
	
}


func insert(req *Request, encoder *json.Encoder, triplets *db.Col, update bool){
	fmt.Printf("Inserting %v", req)

	p := req.Params
	arr := p.([]interface{})
	
	key := arr[0].(string)
	rel := arr[1].(string)
	val := arr[2]

	fmt.Println(key, rel, val)

	// See if there this key/val is already in DB
	queryResult := query_key_rel(key, rel, triplets)
	if len(queryResult) != 0 {
		fmt.Println("Insert: key " + key + " rel " + rel + " already exists")

		if update{
			// insertOrUpdate() now replaces the key/rel with an updated value
			// delete old value, insert new
			for id := range queryResult {
				fmt.Println("Deleting ", id)
				if err := triplets.Delete(id); err != nil {
					panic(err)
				}
			}

			//insert new value
			docID, err := triplets.Insert(map[string]interface{}{
				"key": key,
				"rel": rel,
				"val": val})
			if err != nil {
				panic(err)
			}
			fmt.Println("Inserting ", docID)


		} else{
			// insert() fails if key/rel already exists
			fmt.Println("Insert did not happen, need to return false")
			return
		}
	} else {
	
		docID, err := triplets.Insert(map[string]interface{}{
			"key": key,
			"rel": rel,
			"val": val})
		if err != nil {
			panic(err)
		}
		fmt.Println("Inserting ", docID)
	}
}


func delete(req *Request, encoder *json.Encoder, triplets *db.Col){
	
	p := req.Params
	arr := p.([]interface{})
	
	key := arr[0].(string)
	rel := arr[1].(string)


	queryResult := query_key_rel(key, rel, triplets)

	// Query result are document IDs
	for id := range queryResult {
		fmt.Println("Deleting ", id)
		if err := triplets.Delete(id); err != nil {
			panic(err)
		}
	}
}

func query_key_rel(key string, rel string, triplets *db.Col) (queryResult map[int]struct{}){

	var query interface{}

	//{"n" means "intersection" of the two queries, logical AND

	json.Unmarshal([]byte(`{"n": [{"eq": "` + key + `", "in": ["key"]}, {"eq": "` + rel + `", "in": ["rel"]}]}`), &query)

	q_result := make(map[int]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalQuery(query, triplets, &q_result); err != nil {
		panic(err)
	}

	return q_result
}

func shutdown(myDB *db.DB){
	fmt.Println("Shutting Down DB")
	myDB.Close()
	os.Exit(0)
}


func testretrive(triplets *db.Col){
	var query interface{}
	json.Unmarshal([]byte(`[{"eq": "keyA", "in": ["key"]}]`), &query)
	//json.Unmarshal([]byte(`{"eq": "keyA", "in": ["key"]}`), &query)
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

	defer myDB.Close()
	
	if err := myDB.Create("Triplets"); err != nil {
		//panic(err)
		fmt.Println(err.Error())
	}

	
	triplets := myDB.Use("Triplets")


	// Create indexes here??
	// TODO: Do not create index if it already exists?
	if err := triplets.Index([]string{"key"}); err != nil {
		//panic(err)
		fmt.Println(err.Error())
	}
        
	if err := triplets.Index([]string{"rel"}); err != nil {
		//panic(err)
		fmt.Println(err.Error())
	}

	
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
		go handleConnection(conn, triplets, myDB) // a goroutine handles conn so that the loop can accept other connections
	}
}
