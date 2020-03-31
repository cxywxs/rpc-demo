package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"log"
	"rpc-demo/entity"
)

//连接数据库
func UserDao() *gorm.DB {
	application := AnalysisApplication()
	db, err := gorm.Open("mysql", application.Databases.Root+":"+application.Databases.Password+"@tcp("+application.Databases.Server+":"+application.Databases.Port+")/"+application.Databases.Database+"?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	log.Println("mysql connect succeed!")
	db.SingularTable(true)
	return db
}

func InsertRegisterInfo(message entity.RegisterMessage, db *gorm.DB) string {
	var message1 entity.RegisterMessage
	db.Where("address = ? and port = ? and method = ?", message.Address, message.Port, message.Method).First(&message1)
	if message1.Method != "" {
		return "same"
	}
	db.Create(&message)
	db.Where("address = ? and port = ? and method = ?", message.Address, message.Port, message.Method).First(&message1)
	if message.Method == message1.Method {
		return "true"
	}
	return "false"
}

func FindMethodInfo(method string, db *gorm.DB) entity.RegisterMessage {
	var message entity.RegisterMessage
	fmt.Println(method, "FindMethodInfo method")
	db.Where("method = ?", method).First(&message)
	//test
	fmt.Println(message, "FindMethodInfo")
	return message
}
