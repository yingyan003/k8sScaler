package utils

import (
	"os"
	"strconv"
)

func LoadEnvVar(key, value string) string {
	var v string
	if v = os.Getenv(key); v == "" {
		v = value
	}
	return v
}

func LoadEnvVarInt(key string, value int) int {
	i, _ := strconv.Atoi(LoadEnvVar(key, strconv.Itoa(value)))
	return i
}
