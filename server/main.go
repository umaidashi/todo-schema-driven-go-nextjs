package main

import (
	"server/common/db"
	"server/infra/applogger"
	"server/presentation/http/ogen"
)

func main() {
	applogger.Init()
	db.Init()
	ogen.Init()
}
