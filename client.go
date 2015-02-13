package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type DICT3 struct{
	Key string
	Relationship string
	//Value interface{}
}


func main() {
	
	fmt.Println("start client")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	encoder := json.NewEncoder(conn)
	m := DICT3{"Whatup", "Dog"}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)

	encoder.Encode(&b)

	conn.Close()
	fmt.Println("done")
}
