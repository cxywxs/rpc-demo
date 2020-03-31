package register_service

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"rpc-demo/dao"
	"rpc-demo/entity"
)

//在注册服务端中，需要注意的是将配置文件和mysql确定，以保证正确的运行

//开启服务监听
func ServerListen() {
	application := dao.AnalysisApplication()
	lner, err := net.Listen("tcp", ":"+application.Register.Port)
	if err != nil {
		log.Println("listener creat error", err)
	}
	for {
		conn, err := lner.Accept()
		if err != nil {
			log.Println("accept error", err)
		}
		go handleConnection(conn)
	}
}

//接收后函数
func handleConnection(conn net.Conn) {
	defer conn.Close()
	readconn := bufio.NewReader(conn)
	messagebyte := make([]byte, 4096)
	i := 0
	for {
		n, err := readconn.Read(messagebyte[i:])
		if err != nil || err != io.EOF {
			log.Println(err)
			break
		} else if err == io.EOF {
			break
		}
		i += n
	}
	log.Println("[]byte read", string(messagebyte))

	//读出信息
	var message entity.RegisterMessage
	messagebyte = CutByte(messagebyte)
	_ = json.Unmarshal(messagebyte, &message)
	log.Println(message, "handleConnection message")

	//判断是注册还是请求方法
	if message.Address == "" {
		messageReturn := dao.FindMethodInfo(message.Method, dao.UserDao())
		messagebyte, _ = json.Marshal(&messageReturn)
		log.Println(messageReturn, "handleConnection==")
		_, _ = conn.Write(messagebyte)
	} else {
		messageReturn := dao.InsertRegisterInfo(message, dao.UserDao())
		log.Println(messageReturn, "handleConnection!=")
		_, _ = conn.Write([]byte(messageReturn))
	}
}

//主要的作用是将多余的byte数组切掉
func CutByte(b []byte) []byte {

	for i := range b {
		if b[i] == 0 {
			b = b[:i]
			break
		}
	}
	return b
}
