
package chess

import(
	log "log"
	os  "os"
)

var (
	SCREEN_WIDTH  = 480
	SCREEN_HEIGHT = 740
	SCREEN_TITLE  = "Automatic memory learning chess(AMLC)"
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
