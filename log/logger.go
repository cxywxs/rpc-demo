package log

import "log"

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
