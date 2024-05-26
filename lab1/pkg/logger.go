package pkg

import (
	"log"
	"os"
)

var Logfile, log_err = os.OpenFile("files/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

func SetLogger() {
	log.SetOutput(Logfile)
}
