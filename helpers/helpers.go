package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"math"
	"math/rand"
	"strconv"
	s "strings"
	"time"

	"github.com/nikitamirzani323/gobackend_mastertoto/configs"
)

func HashPasswordMD5(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
func Generator_number() string {
	min := 0
	max := 999
	temp_number := 0
	temp_fixed := ""
	rand.Seed(time.Now().UnixNano())
	temp_number = rand.Intn(max-min) + min

	if temp_number >= 0 && temp_number < 10 {
		temp_fixed = "00" + strconv.Itoa(temp_number)
	}
	if temp_number >= 10 && temp_number < 100 {
		temp_fixed = "0" + strconv.Itoa(temp_number)
	}
	if temp_number >= 100 && temp_number < 999 {
		temp_fixed = strconv.Itoa(temp_number)
	}
	return temp_fixed
}
func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}
func Encryption(datatext string) (string, int) {
	min := 0
	max := 149
	rand.Seed(time.Now().UnixNano())
	// keymap := rand.Intn(max-min) + min
	keymap := rand.Intn(max-min) + min
	var key string = configs.Keymap[keymap]
	var source string = configs.Sourcechar
	result := ""
	for i := 0; i < len(datatext); i++ {
		temp_indexsource := s.Index(source, string(datatext[i]))
		temp_indexkey := s.Index(key, string(key[temp_indexsource]))
		result += string(key[temp_indexkey])
	}
	return result, keymap
}
func Decryption(dataencrypt string) string {
	temp := s.Split(dataencrypt, "|")
	keymap, _ := strconv.Atoi(temp[1])
	var key string = configs.Keymap[keymap]
	var source string = configs.Sourcechar
	result := ""
	for i := 0; i < len(temp[0]); i++ {
		temp_indexkey := s.Index(key, string(dataencrypt[i]))
		temp_indexsource := s.Index(source, string(source[temp_indexkey]))
		result += string(source[temp_indexsource])
	}
	return result
}
