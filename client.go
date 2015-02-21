package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	//"os"
)

type Request struct{
	Method string `json:"method"` 
	Params interface{} `json:"params"`
	Id interface{} `json:"id"`
}

//Params is a 3 element array: [key(string), relationship(string), value(JSON object)]

func insert(key string, rel string, value interface{}, encoder *json.Encoder){
	d3 := []interface{}{key, rel, value}
	m := Request{"insert", d3, "1"}

	encoder.Encode(m)

}

func delete(key string, rel string, encoder *json.Encoder){
	d2 := []interface{}{key, rel}
	m := Request{"delete", d2, "1"}

	encoder.Encode(m)

}

func listKeys(encoder *json.Encoder){
	m := Request{"listKeys", []interface{}{}, "1"}
	encoder.Encode(m)
}


func main() {
	
	fmt.Println("start client")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	encoder := json.NewEncoder(conn)

	//TODO id values?
	
	//Insert 
	//insert("keyA", "relA", map[string]interface{}{"a":3, "b": 1111}, encoder)
	//insert("keyB", "relA", map[string]interface{}{"a":3, "b": 1111}, encoder)
	//insert("keyC", "relA", map[string]interface{}{"a":3, "b": 1111}, encoder)

	//List Keys
	listKeys(encoder)
	
	//Delete
	//delete("keyA", "relA", encoder)



	conn.Close()
	fmt.Println("done")
}
