package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nearbyren/blog-service/global"
	"github.com/nearbyren/blog-service/internal/model"
	"github.com/nearbyren/blog-service/internal/routers"
	"github.com/nearbyren/blog-service/pkg/logger"
	"github.com/nearbyren/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	//配置
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	//日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	//初始化数据库
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

//读全剧配置信息
func setupSetting() error {
	setting, err := setting.NewSetting()
	log.Println("33333")
	if err != nil {
		log.Println("-33333")
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

//数据库
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

//日志
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

//@title 博客系统
//@version 1.0
//@description Go 语言编程之旅：一起用 Go 做项目
//@termsOfService https://github.com/nearbyren/
func main() {
	////加载yaml配置文件
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           router,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeOut,
		WriteTimeout:      global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
	//router := routers.NewRouter()
	//s := &http.Server{
	//	Addr:              ":8000",
	//	Handler:           router,
	//	ReadHeaderTimeout: 10 * time.Second,
	//	WriteTimeout:      10 * time.Second,
	//	MaxHeaderBytes:    1 << 20,
	//}
	//s.ListenAndServe()
	//测试日志
	//global.Logger.Infof("%s: go-programming-tur-book/%s", "eddycjy", "blog-service")
}
