
package chess

import(
	log "log"
	os  "os"
)

const (
	SCREEN_WIDTH  = 400
	SCREEN_HEIGHT = 600
	SCREEN_TITLE  = "AI chess player"
)

var (
	INFO  *log.Logger
	WARN  *log.Logger
	ERROR *log.Logger
)

func init(){
	INFO  = log.New(os.Stdout, "INFO: " , log.Ldate|log.Ltime|log.Lshortfile)
	WARN  = log.New(os.Stdout, "WARN: " , log.Ldate|log.Ltime|log.Lshortfile)
	ERROR = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
