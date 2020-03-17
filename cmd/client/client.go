package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	buf := make([]byte, 51*51)
	conn.Read(buf)

	fmt.Print(string(buf))
}
