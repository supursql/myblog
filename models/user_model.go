package models

import (
	"fmt"
	"myblog/utils"
)

type User struct {
	Id int
	Username string
	Password string
	Status int //正常状态0 删除1
	Createtime int64
}

//插入
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username, password, status, createtime) values(?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

//按照条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username = '%s'", username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码查询
func QueryUserWithParams(username, password string) int {
	sql := fmt.Sprintf("where username = '%s' and password = '%s'", username, password)
	return QueryUserWightCon(sql)
}

