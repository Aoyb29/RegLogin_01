package test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"testing"
)

var db *gorm.DB
var err error

type Users struct {
	Id   int
	Name string
	Psw  string
	Sex  string
	Age  int
	Tel  string
}

func init() {
	db, err = gorm.Open("postgres",
		"host=localhost user=postgres dbname=mydb1 sslmode=disable password=ayb17202114")
	if err != nil {
		panic(err)
		log.Println("数据库连接失败...")
	}
	log.Println("数据初始化成功...")

}

func Test1(t *testing.T) {
	us := &[]Users{}
	db.Table("users").Find(us)
	fmt.Println(us)
}
