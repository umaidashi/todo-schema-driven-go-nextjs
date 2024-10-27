package main

import (
	"server/infra/applogger"
	"server/presentation/http/ogen"
)

func main() {
	applogger.Init()
	ogen.Init()
}
