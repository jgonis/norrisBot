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

func SendMessageUnlessFull(messageChannel chan<- SendMessage, messageToSend SendMessage) {
	select {
	case messageChannel <- messageToSend:
		log.Println("SENDING:", messageToSend.MainMessage)
	default:
		log.Println("ERROR: Message send queue was full, so we're hitting rate limits, dropping message:", messageToSend)
	}
}

func sendData(conn net.Conn, messageChannel <-chan SendMessage) {
	limiter := time.Tick(1500 * time.Millisecond)
	for message := range messageChannel {
		<-limiter
		_, err := fmt.Fprintf(conn, "%s\r\n", message.MainMessage)
		if err != nil {
			log.Println("ERROR:", message.ErrorMessage, err)
		}
	}
}
