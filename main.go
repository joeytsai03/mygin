package main
//import "net/http"
//参考 ： http://www.okyes.me/2016/05/03/go-gin.html
import (
	"./com.zy/config"
	"net/http"
	"time"
)

func main(){
	router := config.Init_router()
	//监听 方法一 ： router.Run(":8888")
	//监听 方法二 ：
	server := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
