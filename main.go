package main

import (
	"encoding/json"
	"fmt"
	"rpc-demo/client"
	"rpc-demo/entity"
)

func main() {

	//register_service.ServerListen()

	c := entity.MathAdd{
		An: 12,
		Bn: 1.35,
	}
	a := client.RpcGet("test", c)
	fmt.Println("asdasd", string(a))

	//service.ExampleServer()
	//DoClient()

}

func DoClient() {
	registerMessage := &entity.RegisterMessage{
		Address:     "127.0.0.1",
		Port:        "text",
		Method:      "8081",
		Description: "bb",
	}
	messagebyte, _ := json.Marshal(registerMessage)
	messagebyte = client.ClientSend(messagebyte, "127.0.0.1:8888")
	fmt.Println(string(messagebyte))
}
