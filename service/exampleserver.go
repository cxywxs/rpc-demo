package service

import (
	"bufio"
	"encoding/json"
	"huangqirichang3.0/client"
	"huangqirichang3.0/dao"
	"huangqirichang3.0/entity"
	"huangqirichang3.0/register_service"
	"io"
	"log"
	"net"
)

//这是一个服务端实例，值得注意的是需要将ExampleServer（）函数中的MathMathod替换成其他想要注册的服务
//目前还需要一个简单快捷的为函数注册到注册服务器的东西，来确保其完整性

//注册服务开启，需要将提供的方法输入与输出需要[]byte
func StartServer(method func([]byte) []byte, methodname string, description string) {
	RegisterServer(methodname, description)
	ExampleServer(method)
}

func ExampleServer(method func([]byte) []byte) {
	lner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println("listener creat error", err)
	}
	for {
		conn, err := lner.Accept()
		if err != nil {
			log.Println("accept error", err)
		}
		go TestMethod(conn, method)
	}
}

//注册服务
func RegisterServer(method, description string) {
	a := dao.AnalysisApplication()
	var register entity.RegisterMessage
	register.SetMessage(a.ServerInfo.Server, a.ServerInfo.Port, method, description)
	messagebyte, _ := json.Marshal(register)
	messagebyte = client.ClientSend(messagebyte, a.Register.Server+":"+a.Register.Port)
}

func TestMethod(conn net.Conn, method func([]byte) []byte) {
	defer conn.Close()
	readconn := bufio.NewReader(conn)
	//var messagebyte []byte
	//_, _ = readconn.Read(messagebyte)

	messagebyte := make([]byte, 4096)
	reader := bufio.NewReader(readconn)
	i := 0
	for {
		n, err := reader.Read(messagebyte[i:])
		if err != nil || err != io.EOF {
			log.Println(err)
			break
		} else if err == io.EOF {
			break
		}
		i += n
	}
	log.Println(string(messagebyte))

	////这是接收的参数结构体
	//var mathAdd MathAdd
	//messagebyte=register_service.CutByte(messagebyte)
	//_ = json.Unmarshal(messagebyte, &mathAdd)
	//fmt.Println(string(messagebyte))
	//messageReturn := Float64ToByte(MathMathod(mathAdd))
	messagebyte = register_service.CutByte(messagebyte)
	messageReturn := method(messagebyte)
	_, _ = conn.Write(messageReturn)
	log.Println(string(messageReturn))

}

//方法示例
func MathMathod(b []byte) []byte {
	var mathAdd entity.MathAdd
	var out entity.Output
	err := json.Unmarshal(b, &mathAdd)
	if err != nil {
		log.Println(err)
	}
	log.Println(mathAdd)
	out.C = float64(mathAdd.An) + mathAdd.Bn
	by, err := json.Marshal(&out)
	if err != nil {
		log.Println(err)
	}
	log.Println("方法内by：", string(by), "out.c", out.C)
	return by
}
