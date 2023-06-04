package main

import (
	"fmt"
	"log"

	"github.com/ljcbaby/select/config"
	"github.com/ljcbaby/select/router"
)

var conf *config.Config

func main() {
	// 加载配置文件
	var err error
	conf, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化路由
	r := router.SetupRouter()

	// 启动 HTTP 服务器
	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	err = r.Run(addr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
