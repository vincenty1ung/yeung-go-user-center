package main

import (
	"flag"
	"fmt"

	"github.com/uncleyeung/yeung-user-center/api/user/internal/config"
	"github.com/uncleyeung/yeung-user-center/api/user/internal/handler"
	"github.com/uncleyeung/yeung-user-center/api/user/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "api/user/etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
