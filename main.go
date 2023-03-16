/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/3/15 07:46
 */

// @Title ck blog v1

package main

import (
	"github.com/CKzcb/ckBlog/global"
	"github.com/CKzcb/ckBlog/internal/model"
	"github.com/CKzcb/ckBlog/internal/routers"
	"github.com/CKzcb/ckBlog/pkg/logger"
	"github.com/CKzcb/ckBlog/pkg/setting"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	// config
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	// db
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	// logger
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	global.Logger.Info("the app init ...")
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

//
// setupSetting
//  @Description: init setting from configs
//  @return error
//
func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	// read server config
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	// read app config
	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	// read database config
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	// handle timeout duration
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName +
		global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
