package server

import (
	"os"
	"strconv"

	log "github.com/cihub/seelog"
	"github.com/joho/godotenv"
)

// Config Configuration for server object
type Config struct {
	Port int
}

// NewServerConf gets a new server configuration
func NewServerConf() (c *Config) {
	env, err := godotenv.Read()
	if err != nil {
		log.Critical("Error loading .env file")
		os.Exit(1)
	}

	port, err := strconv.Atoi(env["Port"])
	if err != nil {
		log.Error("Error assigning port from env: ", err)
		os.Exit(1)
	}
	c = &Config{
		Port: port,
	}

	return c
}
