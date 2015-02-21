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

func insert(key string, rel string, value interface{}){
	d3 := []interface{}{key, rel, value}
	m := Request{"insert", d3, "1"}

	encoder := getEncoder()
	encoder.Encode(m)

}

func delete(key string, rel string){
	d2 := []interface{}{key, rel}
	m := Request{"delete", d2, "1"}

	encoder := getEncoder()
	encoder.Encode(m)

}

func listKeys(){
	m := Request{"listKeys", []interface{}{}, "1"}

	encoder := getEncoder()
	encoder.Encode(m)
}


func listIDs(){
	m := Request{"listIDs", []interface{}{}, "1"}

	encoder := getEncoder()
	encoder.Encode(m)
}

func shutdown(){
	m := Request{"shutdown", []interface{}{}, "1"}

	encoder := getEncoder()
	encoder.Encode(m)
}


func getEncoder() (encoder *json.Encoder){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	e := json.NewEncoder(conn)

	return e
}

func main() {
	
	fmt.Println("start client")


	//TODO id values?
	
	//Insert 
	insert("keyA", "relA", map[string]interface{}{"a":3, "b": 1111})
	insert("keyB", "relA", map[string]interface{}{"a":3, "b": 1111})
	insert("keyC", "relA", map[string]interface{}{"a":3, "b": 1111})
	insert("keyC", "relA", map[string]interface{}{"a":3, "b": 1111})
	insert("keyC", "relB", map[string]interface{}{"a":3, "b": 1111})
	
	//Delete
	//delete("keyA", "relA")
	//delete("keyB", "relA")
	//delete("keyC", "relA")
	//delete("keyC", "relB")	

	
	
	//List Keys
	listKeys()

	listIDs()

	//shutdown()


	//conn.Close()
	fmt.Println("done")
}
