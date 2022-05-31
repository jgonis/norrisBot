package main

import (
	"fmt"
	"log"
	"net"
)

func sendData(conn net.Conn, message string, errMessage string) {
	_, err := fmt.Fprintf(conn, "%s\r\n", message)
	if err != nil {
		log.Println(errMessage, err)
	}
}
