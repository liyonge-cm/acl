package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"acl/config"
	"acl/router"
	_ "acl/service/apis"
	"acl/service/logger"
	"acl/service/mysql"
)

func main() {
	configFile := flag.String("c", "./config/config.yml", "yml config")
	flag.Parse()
	if *configFile == "" {
		fmt.Println("invalid config file")
		os.Exit(0)
	}
	// 加载配置文件
	config.LoadConfig(*configFile)
	logger.Logger.Info("Service Start")
	db := mysql.NewMySQL(context.Background(), &config.Cfg.MySQL)
	db.WithLogger(logger.Logger)
	db.ConnDB()
	router.Init()
}
