package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var err error
var db *gorm.DB

//type Users struct {
//	Id   int    `json:"id"; gorm:"column:id"`
//	Name string `json:"name";gorm:"column:name"`
//	Psw  string `json:"psw";gorm:"column:psw"`
//	Sex  string `json:"sex";gorm:"column:sex"`
//	Age  int    `json:"age";gorm:"column:age"`
//	Tel  string `json:"tel";gorm:"column:tel"`
//}

type Users struct {
	Id   int
	Name string
	Psw  string
	Sex  string
	Age  int
	Tel  string
}

func main() {
	db, err = gorm.Open("postgres",
		"host=localhost user=postgres dbname=mydb1 sslmode=disable password=ayb17202114")
	if err != nil {
		panic(err)
		log.Println("数据库连接失败...")
	}
	log.Println("数据初始化成功...")

	r := gin.Default()
	r.GET("/login/:name/:psw", login) //用户登录
	r.POST("/reg", reg)               //用户注册
	r.Run("localhost:1118")
	defer db.Close()

}
func reg(context *gin.Context) {
	var ageInt int
	name := context.PostForm("name")
	psw := context.PostForm("psw")
	sex := context.PostForm("sex")
	age := context.PostForm("age")
	ageInt, err = strconv.Atoi(age) //string转int
	tel := context.PostForm("tel")
	us := Users{}
	db.Where("name=?", name).First(&us)
	log.Println(us)
	if us.Name != "" {
		rand.Seed(time.Now().UnixNano())
		nameNew := name + strconv.Itoa(rand.Intn(2020))
		context.JSON(401, gin.H{
			"message": "用户名已存在,您可以尝试使用:" + nameNew,
			"success": false,
		})
		return
	} else {
		us := Users{Name: name, Psw: psw, Sex: sex, Age: ageInt, Tel: tel}
		db.Create(&us)
		context.JSON(200, gin.H{
			"message": "注册成功...",
			"success": true,
		})
	}

}
func login(context *gin.Context) {
	//name := context.DefaultQuery("name", "admin")
	//psw := context.DefaultQuery("psw", "888888")
	name := context.Param("name")
	psw := context.Param("psw")
	us := Users{}
	db.Where(&Users{Name: name, Psw: psw}).First(&us)
	log.Println(us)
	if us.Name == "" || us.Psw == "" {
		context.JSON(401, gin.H{
			"message": "登陆失败，请检查用户名或密码",
			"success": false,
		})
		return
	}
	context.JSON(200, gin.H{
		"message": "登陆成功...",
		"success": true,
	})

}
