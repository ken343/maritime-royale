package render

import "log"

func Init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
