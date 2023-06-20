package main

import (
	"fmt"
	"github.com/renfy96/yy-layout-v1/pkg/config"
	"github.com/renfy96/yy-layout-v1/pkg/http"
	"github.com/renfy96/yy-layout-v1/pkg/log"
	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	app, cleanup, err := newApp(conf, logger)
	if err != nil {
		panic(err)
	}
	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
	defer cleanup()
}
