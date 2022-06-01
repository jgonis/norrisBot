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
	burstyLimiter := initializeBurstyLimiter()

	for message := range messageChannel {
		<-burstyLimiter
		go refillRateLimitQueue(burstyLimiter)
		_, err := fmt.Fprintf(conn, "%s\r\n", message.MainMessage)
		if err != nil {
			log.Println("ERROR:", message.ErrorMessage, err)
		}
	}
}

func initializeBurstyLimiter() chan time.Time {
	queueSize := 20
	burstyLimiter := make(chan time.Time, queueSize)
	for i := 0; i < queueSize; i++ {
		burstyLimiter <- time.Now()
	}
	return burstyLimiter
}

func refillRateLimitQueue(rateLimitChannel chan<- time.Time) {
	<-time.Tick(30 * time.Second)
	rateLimitChannel <- time.Now()
}
