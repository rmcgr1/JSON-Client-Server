package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	//"os"
)

type DICT3 struct{
	Key string `json:"method"`
	Relationship string `json:"params"`
	Id string `json:"id"`
}


func main() {
	
	fmt.Println("start client")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	encoder := json.NewEncoder(conn)
	m := DICT3{"lookup", "something", "1"}

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
