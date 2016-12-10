package start

import (
	"io"
	"os"
	"strconv"
	SYS "syscall"

	"github.com/grindlemire/Socket/server"
	DEATH "github.com/vrecan/death"
)

// Run starts the webserver
func Run() {
	var goRoutines []io.Closer
	death := DEATH.NewDeath(SYS.SIGINT, SYS.SIGTERM)

	server := server.NewServer(server.NewServerConf())
	server.Start()
	goRoutines = append(goRoutines, server)

	death.WaitForDeath(goRoutines...)

}

// GetEnvInt gets an environment variable and returns it as an int
func GetEnvInt(key string) (val int, err error) {
	strVal := os.Getenv(key)
	val, err = strconv.Atoi(strVal)
	return val, err
}
