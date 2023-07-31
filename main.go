package main

import (
	"AlittleRequire/config"
	"AlittleRequire/pkg/logger"
	"AlittleRequire/repository/dao"
	"AlittleRequire/routes"
)

func main() {
	loadServer()
	r := routes.NewRoute()
	_ = r.Run(config.Config.System.HttpPort)
}

func loadServer() {
	config.InitConfig()
	dao.MysqlInit()
	logger.InitLog()
}
