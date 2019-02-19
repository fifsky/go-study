package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilibs/gosql"
)

type Users struct {
	Id        int       `form:"id" json:"id" db:"id"`
	Name      string    `form:"name" json:"name" db:"name"`
	Password  string    `form:"password" json:"password" db:"password"`
	NickName  string    `form:"nick_name" json:"nick_name" db:"nick_name"`
	Email     string    `form:"email" json:"email" db:"email"`
	Status    int       `form:"status" json:"status" db:"status"`
	Type      int       `form:"type" json:"type" db:"type"`
	CreatedAt time.Time `form:"-" json:"created_at" db:"created_at"`
	UpdatedAt time.Time `form:"-" json:"updated_at" db:"updated_at"`
}

func (u *Users) DbName() string {
	return "default"
}

func (u *Users) TableName() string {
	return "users"
}

func (u *Users) PK() string {
	return "id"
}

func (u *Users) AfterChange() {
	fmt.Println(fmt.Sprintf("user %s(%d) nikename change to %s", u.Name, u.Id, u.NickName))
}

func main() {
	configs := make(map[string]*gosql.Config)

	configs["default"] = &gosql.Config{
		Enable:  true,
		Driver:  "mysql",
		Dsn:     "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Asia%2FShanghai",
		ShowSql: true,
	}

	//connection database
	gosql.Connect(configs)

	//get user
	user := &Users{Id: 1}
	gosql.Model(user).Get()
	fmt.Println(user)

	gosql.Model(&Users{Id: 1, NickName: "fifsky2"}).Update()
}
