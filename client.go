package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"bufio"
	"time"
)

type Request struct{
	Method string `json:"method"` 
	Params interface{} `json:"params"`
	Id interface{} `json:"id"`
}

type Response struct{
	Result interface{} `json:"result"`
	Id interface{} `json:"id"`
	Error interface{} `json:"error"`
}

//Params is a 3 element array: [key(string), relationship(string), value(JSON object)]

func insert(key string, rel string, value interface{}){
	d3 := []interface{}{key, rel, value}
	m := Request{"insert", d3, "1"}

	encoder, decoder := getEncoder()
	encoder.Encode(m)

	res := new(Response)
	decoder.Decode(&res)
	fmt.Println("Recieved: %+v", res)

}

func insertOrUpdate(key string, rel string, value interface{}){
	d3 := []interface{}{key, rel, value}
	m := Request{"insertOrUpdate", d3, "1"}

	encoder, _ := getEncoder()
	encoder.Encode(m)

}


func delete(key string, rel string){
	d2 := []interface{}{key, rel}
	m := Request{"delete", d2, "1"}

	encoder,_ := getEncoder()
	encoder.Encode(m)

}

func lookup(key string, rel string){
	d2 := []interface{}{key, rel}
	m := Request{"lookup", d2, "1"}

	encoder, decoder := getEncoder()
	encoder.Encode(m)

	
	res := new(Response)
	decoder.Decode(&res)
	fmt.Println("Recieved: %+v", res)

}

func listKeys(){
	m := Request{"listKeys", []interface{}{}, "1"}

	encoder,decoder := getEncoder()
	encoder.Encode(m)
	
	res := new(Response)
	decoder.Decode(&res)
	fmt.Println("Recieved: %+v", res)
}


func listIDs(){
	m := Request{"listIDs", []interface{}{}, "1"}

	encoder,decoder := getEncoder()
	encoder.Encode(m)

	res := new(Response)
	decoder.Decode(&res)
	fmt.Println("Recieved: %+v", res)
	
}

func shutdown(){
	m := Request{"shutdown", []interface{}{}, "1"}

	encoder,_ := getEncoder()
	encoder.Encode(m)
}


func getEncoder() (encoder *json.Encoder, decoder *json.Decoder){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	e := json.NewEncoder(conn)
	d := json.NewDecoder(conn)
	
	return e, d
}

func readInput(){

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		text := scanner.Text()

		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			log.Fatal("Connection error", err)
		}

		conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			
		ip_writer := bufio.NewWriter(conn)
		ip_reader := bufio.NewReader(conn)
		
		fmt.Println("Sending: ", text)
			
		ip_writer.WriteString(text)
		ip_writer.Flush()
		
		line, _ := ip_reader.ReadString('\n')
		fmt.Println("Recieved: ", line)

		conn.Close()
	}

		
}

func main() {
	

	readInput()
	

	//TODO id values?
	
	//Insert 
	//insert("keyA", "relA", map[string]interface{}{"a":3, "b": 1111})
	//lookup("keyC", "relA")
	//insert("keyB", "relA", map[string]interface{}{"a":3, "b": 1111})
	//lookup("keyB", "relA")
	//insert("keyC", "relA", map[string]interface{}{"a":3, "b": 1111})
	//insert("keyC", "relA", map[string]interface{}{"a":3, "b": 1111})
	//insertOrUpdate("keyC", "relA", map[string]interface{}{"a":9999, "b": 999})
	//insertOrUpdate("keyC", "relA", map[string]interface{}{"a":888, "b": 888})

	
	
	//delete("keyA", "relA")
	//delete("keyB", "relA")
	//delete("keyC", "relA")
	//delete("keyC", "relB")	

	//listKeys()

	//listIDs()

	//shutdown()
}
