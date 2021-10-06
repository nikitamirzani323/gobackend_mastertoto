package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"math"
	"math/rand"
	"strconv"
	"time"
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
