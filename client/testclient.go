package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"rpc-demo/entity"
)

func Testclient() {
	registerMessage := &entity.RegisterMessage{
		Address:     "127.0.0.1",
		Port:        "text",
		Method:      "8081",
		Description: "bb",
	}
	messagebyte, _ := json.Marshal(registerMessage)
	//a := "qqqq"
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	client, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	client.Write(messagebyte)
	fmt.Println(messagebyte)
	b := make([]byte, 4096)
	reader := bufio.NewReader(client)
	i := 0
	for {
		n, err := reader.Read(b[i:])
		if err != nil || err != io.EOF {
			fmt.Println(err)
			break
		} else if err == io.EOF {
			break
		}
		i += n
	}
	fmt.Println(string(b))
}
