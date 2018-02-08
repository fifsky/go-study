package main

import (
	"fmt"
	"net/http"

	"app/config"
	"app/models/paymodel"
	"app/repository/pay"

	"github.com/verystar/golib/db/upper"
	"github.com/verystar/golib/do"
	"github.com/verystar/golib/logger"
	"github.com/verystar/golib/monitor"
	"github.com/verystar/golib/redis"
)

func Run(port string) {

	http.HandleFunc("/getuser", func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		user_id := req.Form.Get("user_id")
		//POST params
		//store_sn := req.PostFormValue("store_sn")

		if user_id == "" {
			w.Write(do.H{
				"code": 201,
				"msg":  "user_id can not emtpy",
			}.MustJSON())
			return
		}

		user := &paymodel.Users{}
		has, err := pay.Users().Where("user_id = ?", user_id).Get(user)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if err != nil {
			w.Write(do.H{
				"code": 201,
				"msg":  err,
			}.MustJSON())
			return
		}
		if !has {
			w.Write(do.H{
				"code": 202,
				"msg":  "user not found",
			}.MustJSON())
			return
		}

		w.Write(do.H{
			"code": 200,
			"data": user,
		}.MustJSON())
	})
	fmt.Println("HTTP Server Start Listen:[localhost:" + port + "]")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	//connect redis
	redis.Connect(config.App.Redis)
	//connect db
	upper.Connect(config.App.DB)

	go monitor.NewRedisMonitor(config.App.StatDB).Run()

	//运行8080端口
	Run("8080")
}
