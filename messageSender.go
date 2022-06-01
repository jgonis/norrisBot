package main

import (
	"fmt"
	"log"
	"net"
)

type SendMessage struct {
	MainMessage  string
	ErrorMessage string
}

func sendData(conn net.Conn, messageChannel chan SendMessage) {
	for message := range messageChannel {
		_, err := fmt.Fprintf(conn, "%s\r\n", message.MainMessage)
		if err != nil {
			log.Println(message.ErrorMessage, err)
		}
	}
}
