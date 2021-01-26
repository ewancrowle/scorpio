package core

type Config struct {
	Port int
}

var DefaultConfig = Config{
	Port: 11042,
}
