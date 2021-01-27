package config

import "os"

var (
	Address string
)

func Configure() {
	Address = os.Getenv("ADDR")
}
