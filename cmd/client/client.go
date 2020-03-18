package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/JosephZoeller/maritime-royale/pkg/mrp"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	var message = make([]byte, 0)
	var newMRP mrp.MRP

	for {
		var buf = make([]byte, 1024)
		conn.Read(buf)
		for k, v := range buf {
			if v == 0 {
				buf = buf[0:k]
				break
			}
		}

		message = append(message, buf...)
		newMRP, err = mrp.ReadMRP(message)
		if err == nil {
			break
		}

	}

	if newMRP.Request == "MAP" {
		var mapLines = strings.Split(newMRP.Body, newMRP.Footers[0])
		for k, v := range mapLines {
			if k != len(mapLines)-1 {
				fmt.Println(v)
			}
		}
	}
}
