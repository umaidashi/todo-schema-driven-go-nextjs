package main

import (
	"server/common/api"
	"server/common/config"
	"server/common/db"
	"server/infra/applogger"
)

func main() {
	config.Init()
	applogger.Init()
	db := db.Init()
	api.Init(db)
	applogger.Info("Stock Service started.")
}
