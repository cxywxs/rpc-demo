package client

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"rpc-demo/entity"
	"rpc-demo/register_service"
)

//客户端的用法：使用RpcGet（）函数来进行rpc调用，值得注意的是RpcGet（）将注册服务器地址写死在里面
//之后需要修改

//用于客户端发送信息到服务端
func ClientSend(messagebyte []byte, serverAddr string) []byte {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Println("Resolve TCPAddr error", err)
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	defer conn.Close()
	log.Println(string(messagebyte))
	_, _ = conn.Write(messagebyte)
	//没有接收
	messagereturn := make([]byte, 4096)
	reader := bufio.NewReader(conn)
	i := 0
	for {
		n, err := reader.Read(messagereturn[i:])
		if err != nil || err != io.EOF {
			log.Println(err)
			break
		} else if err == io.EOF {
			break
		}
		i += n
	}
	return messagereturn

}

func GetMethod(messagebyte []byte) {
	var message entity.RegisterMessage
	err := json.Unmarshal(messagebyte, &message)
	if err != nil {
		log.Println(err)
	}

}

//获取函数结果
func RpcGet(method string, input interface{}) []byte {
	var message entity.RegisterMessage
	message.Method = method
	message.Address = ""
	message.Port = ""
	message.Description = ""
	//test
	log.Println(message, "RPC before conn ClientSend")
	messagebyte, _ := json.Marshal(&message)
	//test
	log.Println(messagebyte, "RPC before conn ClientSend")
	messagebyte = ClientSend(messagebyte, "127.0.0.1:8888")
	messagebyte = register_service.CutByte(messagebyte)
	err := json.Unmarshal(messagebyte, &message)
	if err != nil {
		log.Println(err)
	}
	methodbyte, _ := json.Marshal(&input)
	//test
	log.Println(messagebyte, "RPC after conn ClientSend")
	messagebyte = ClientSend(methodbyte, message.Address+":"+message.Port)
	messagebyte = register_service.CutByte(messagebyte)
	return messagebyte
}
