package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net/textproto"
	"strings"
)

func handleMessages(conn *tls.Conn) {
	connectionReader := textproto.NewReader(bufio.NewReader(conn))
	for {
		line, err := connectionReader.ReadLine()
		if err == nil {
			parseMessage(line, conn)
		} else {
			log.Println("error: ", err)
			break
		}
	}
}

func parseMessage(line string, conn *tls.Conn) {
	fmt.Println(line)
	if strings.HasPrefix(line, "PING") {
		log.Println("received PING message, responding with PONG")
		returnMessage := strings.TrimPrefix(line, "PING")
		sendData(conn, "PONG"+returnMessage, "error replying to PING message with PONG message")
	}
}
