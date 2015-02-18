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


type DICT3 struct{
	Key string
	Relationship string
	Value interface{}
}


func main() {
	
	fmt.Println("start client")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	encoder := json.NewEncoder(conn)
	//d3 := DICT3{"keyA", "relA", map[string]interface{}{"a": 3, "b" : 1}}  
	d3 := []interface{}{"keyA", "relA", map[string]interface{}{"a": 3, "b" : 1}}
	m := Request{"insert", d3, "1"}

	//b, err := json.Marshal(m)
	//fmt.Println(string(m))
	
	
	if err != nil {
		fmt.Println("error: ", err)
	}
	//os.Stdout.Write(b)

	encoder.Encode(m)

	conn.Close()
	fmt.Println("done")
}
