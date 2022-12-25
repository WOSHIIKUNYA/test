package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Username string `json:"Username"`
	Possward string `json:"Possward"`
}

func register(m *gin.Context) {
	dp, err := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/mine")
	var user user
	if err != nil {
		fmt.Println(err)
	}
	m.ShouldBind(&user)
	_, err1 := dp.Exec("insert into bookuser (name,Possward) values (?,?)", user.Username, user.Possward)
	if err1 != nil {
		fmt.Println(err1)
	}
}
func Login(m *gin.Context) {
	dp, err := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/mine")
	if err != nil {
		fmt.Println(err)
	}
	var User user
	m.ShouldBind(&User)
	h, err1 := dp.Query("select name,possward from bookuser")
	if err1 != nil {
		fmt.Println(err1)
	}
	for h.Next() {
		var user1 user
		err = h.Scan(&user1.Username, &user1.Possward)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(user1.Possward)
		if user1.Username == User.Username && user1.Possward == User.Possward {
			m.SetCookie("Name", user1.Username, 3600, "./", "localhost", false, true)
			m.String(200, "OK")
			m.String(200, "ok")
			return
		}
	}
	m.String(200, "false")
}
func mule(m *gin.Context) {
	b, err := m.Cookie("name")
	if err != nil {
		fmt.Println(err)
	}
	if b == "" {
		m.String(200, "false")
		return
	}
	dp, err1 := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/mine")
	if err1 != nil {
		fmt.Println(err)
	}

}
func main() {

	m := gin.Default()
	m.POST("/register", register)
	m.POST("/login", Login)
	m.GET(".mule", mule)
	m.Run()
}
