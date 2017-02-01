package main

import (
	"../server"
	"../client"
	"../protocol"
	"fmt"
)

func Handler(v interface{})  (r interface{}, err error) {
	r = v.(string) + " | kay"
	return r, nil
}

func main() {
	server := server.NewServer("127.0.0.1:8080", protocol.NewJsonProtocol(), server.ServerHandler(Handler))
	server.Start()
	c,_:= client.Dail("tcp", server.Address, protocol.NewJsonProtocol())
	c.Send("good")
	v , _ := c.Receive()
	fmt.Println(v.(string))
}