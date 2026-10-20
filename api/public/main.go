package main

import (
	"fmt"

	"github.com/si-gui/api/public/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	var c config.Config
	conf.MustLoad("etc/public.yaml", &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	fmt.Printf("公开API网关启动成功，监听端口: %s\n", c.Port)
	logx.Infof("公开API网关启动成功，监听端口: %s", c.Port)

	server.Start()
}
