package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type SendMessage struct {
	MainMessage  string
	ErrorMessage string
}

func sendData(conn net.Conn, messageChannel chan SendMessage) {
	limiter := time.Tick(1500 * time.Millisecond)
	for message := range messageChannel {
		<-limiter
		_, err := fmt.Fprintf(conn, "%s\r\n", message.MainMessage)
		if err != nil {
			log.Println(message.ErrorMessage, err)
		}
	}
}
