package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
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
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err : %v",err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("")
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

func setupLogger()error{
	fileName:=global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt
	//使用lumberjack.Logger 作为日志库的io.Writer
	global.Logger=logger.NewLogger(&lumberjack.Logger{
		Filename:   fileName,
		//最大占用空间600MB
		MaxSize:    600,
		//日志文件最大生存周期10天
		MaxAge:     10,
		//日志文件名的时间格式为本地时间
		LocalTime:  true,
	},"",log.LstdFlags).WithCaller(2)
	return nil
}

func setupDBEngine()error{
	var err error
	//注意 ： 此处不能写成  :=  因为这存在很大问题，会重新声明并创建左侧的新局部变量。并不能赋值给全局变量global.DBEngine
	global.DBEngine,err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
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
	//global.Logger.Info("1111111test")
	s:=&http.Server{
		Addr: ":"+global.ServerSetting.HttpPort,
		Handler:router,
        ReadTimeout:global.ServerSetting.ReadTimeOut,
        WriteTimeout:global.ServerSetting.ReadTimeOut,
        MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}