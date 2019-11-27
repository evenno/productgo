package models

import (
	"fmt"
	"productgo/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	CreateTime int64
}

func InsertUser(user *User) (int64, error) {
	sql := "insert into user(username, password, createtime) values(?, ? , ?)"
	return utils.ExecSQL(sql, user.Username, user.Password, user.CreateTime)
}

//1 a 12345
//2 b 12345
//a 12345
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("select id from user where username='%s'", username)
	row := utils.QueryRow(sql)

	id := 0
	_ = row.Scan(&id)

	return id
}

func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("select id from user where username='%s' and password='%s'", username, password)
	row := utils.QueryRow(sql)

	id := 0
	_ = row.Scan(&id)

	return id
}
