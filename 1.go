package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type user struct {
	Username string
	Nickname string
	Sex      uint8
	birthday time.Time
}

// 定义结构体时没有大写或者2次编码，没法解析
func main() {
	u := user{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}
