package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)
// 全局变量初始化。> init 》 main
//控制应用程序的初始化流程
func init()  {
	err:=setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v",err)
	}
}

func setupSetting()error{
	setting,err:=setting.NewSetting()
	if err != nil {
		return err
	}
	err=setting.ReadSection("Server",&global.ServerSetting)
	if err != nil {
		return err
	}
	err=setting.ReadSection("App",&global.AppSetting)
	if err != nil {
		return err
	}
	err=setting.ReadSection("DataBase",&global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}
func main(){
	//r:=gin.Default()
	////engine.Use(Logger(), Recovery())  输出请求日志，并标准化日志的格式  /异常捕获，防止因为panic导致服务崩溃，同时将异常日志的格式标准化
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200,gin.H{"message": "pong"})
	//})
	//r.Run()
    gin.SetMode(global.ServerSetting.RunMode)
	router:=routers.NewRouter()
	s:=&http.Server{
		Addr: ":"+global.ServerSetting.HttpPort,
		Handler:router,
        ReadTimeout:global.ServerSetting.ReadTimeOut,
        WriteTimeout:global.ServerSetting.ReadTimeOut,
        MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}