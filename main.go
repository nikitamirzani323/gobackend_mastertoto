package main

import (
	"log"

	"github.com/nikitamirzani323/gobackend_mastertoto/configs"
	"github.com/nikitamirzani323/gobackend_mastertoto/router"
)

func main() {
	configs.Init()
	app := router.Init()
	log.Fatal(app.Listen(":7071"))
}
