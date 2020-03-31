package register_service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"rpc-demo/entity"
)

func TestServer() {
	list, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := list.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go TestMethod1(conn)
	}
}

func TestMethod1(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	i := 0
	a := make([]byte, 4096)
	for {
		n, err := reader.Read(a[i:])
		if err != nil || err != io.EOF {
			fmt.Println(err)
			break
		} else if err == io.EOF {
			break
		}
		i += n
	}
	var messagebyte entity.RegisterMessage
	a = CutByte(a)
	_ = json.Unmarshal(a, &messagebyte)
	fmt.Println(a)
	fmt.Println(messagebyte.Description, messagebyte.Port)
	n, err := conn.Write([]byte("pong"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("n=", n)

	fmt.Println(conn.RemoteAddr().String(), "is disconnected")
}
