package env

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func GetInt(key string) (i int) {
	val := os.Getenv(key)
	i, _ = strconv.Atoi(val)
	return
}

func GetString(key string) (s string) {
	s = os.Getenv(key)
	return
}

func GetBool(key string) (b bool) {
	val := os.Getenv(key)
	b, _ = strconv.ParseBool(val)
	return
}
