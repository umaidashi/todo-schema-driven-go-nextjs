package main

import (
	"server/common/api"
	"server/common/db"
	"server/infra/applogger"
)

func main() {
	applogger.Init()
	db := db.Init()
	api.Init(db)
}
