package routes

import (
	"dream/serve/config"
	"dream/serve/dao"
	"dream/serve/utils"
	"log"
	"net/http"
	"time"
)

// InitGlobalVariable 初始化全局变量
func InitGlobalVariable() {
	// 读取配置文件
	utils.InitViper()
	// 初始化 Logger
	utils.InitLogger()
	// 初始化数据库
	dao.DB = utils.InitMySQLDB()
	// 初始化控制权限
	utils.InitCasbin(dao.DB)
}

// BackendServer 后台服务
//func BackendServer() *http.Server {
//	backPort := config.Cfg.Server.BackPort
//	log.Printf("后台服务启动于 %s 端口", backPort)
//	return &http.Server{
//		Addr:         backPort,
//		Handler:      BackRouter(),
//		ReadTimeout:  5 * time.Second,
//		WriteTimeout: 10 * time.Second,
//	}
//}

// FrontendServer 前台服务
func FrontendServer() *http.Server {
	frontPort := config.Cfg.Server.FrontPort
	log.Printf("前台服务启动于 %s 端口", frontPort)
	return &http.Server{
		Addr:         frontPort,
		Handler:      FrontRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
