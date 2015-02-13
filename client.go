package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type P struct {
	M, N int64
}

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
	encoder := gob.NewEncoder(conn)
	p := &P{1, 2}
	encoder.Encode(p)
	conn.Close()
	fmt.Println("done")
}
