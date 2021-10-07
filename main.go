package main

import (
	"log"
	"strconv"

	"github.com/nikitamirzani323/gobackend_mastertoto/db"
	"github.com/nikitamirzani323/gobackend_mastertoto/helpers"
	"github.com/nikitamirzani323/gobackend_mastertoto/router"
)

func main() {
	db.Init()
	app := router.Init()

	data := "aaaabr==isb"

	temp_encr, keymap := helpers.Encryption(data)
	temp_encr_final := temp_encr + "|" + strconv.Itoa(keymap)
	temp_decp := helpers.Decryption(temp_encr_final)

	log.Println(data)
	log.Println(temp_encr_final)
	log.Println(temp_decp)

	log.Fatal(app.Listen(":7071"))
}
