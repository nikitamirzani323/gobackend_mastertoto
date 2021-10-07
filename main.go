package main

import (
	"log"
	s "strings"

	"github.com/nikitamirzani323/gobackend_mastertoto/configs"
	"github.com/nikitamirzani323/gobackend_mastertoto/router"
)

var key string = configs.Keymap[60]
var source string = configs.Sourcechar

func main() {
	configs.Init()
	app := router.Init()

	data := "Indosuperbet.com dan nuke.com"

	temp_encr := ""
	temp_decp := ""
	log.Println(len(configs.Keymap))
	log.Println(len(source))
	log.Println(len(key))
	// log.Println(string(key[1]))
	for i := 0; i < len(data); i++ {
		// log.Println(string(data[i]))
		// temp_encr += string(key[i])
		temp_indexsource := s.Index(source, string(data[i]))
		temp_indexkey := s.Index(key, string(key[temp_indexsource]))
		// temp_stringkey := keyback(temp_indexkey)
		temp_encr += string(key[temp_indexkey])
		// log.Printf("%d - %d - %s", temp_indexsource, temp_indexkey, temp_stringkey)
	}
	for i := 0; i < len(temp_encr); i++ {
		// log.Println(string(data[i]))
		// temp_encr += string(key[i])
		temp_indexkey := s.Index(key, string(key[i]))
		temp_indexsource := s.Index(source, string(data[temp_indexkey]))
		// temp_stringkey := sourceback(temp_indexsource)
		temp_decp += string(source[temp_indexsource])
		// log.Printf("%d - %d - %s", temp_indexsource, temp_indexkey, temp_stringkey)
	}
	// for x := 0; x < len(source); x++ {
	// 	if source[x] == 80 {
	// 		log.Println(string(source[x]))
	// 	}
	// }

	// log.Println(string(80))
	log.Println(data)
	log.Println(temp_encr + "|60")
	log.Println(temp_decp + "|60")

	log.Fatal(app.Listen(":7071"))
}

func keyback(index int) string {
	result := ""
	for i := 0; i < len(key); i++ {
		if index == i {
			result = string(key[i])
		}
	}
	return result
}
func sourceback(index int) string {
	result := ""
	for i := 0; i < len(source); i++ {
		if index == i {
			result = string(source[i])
		}
	}
	return result
}
