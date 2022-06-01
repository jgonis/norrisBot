package main

import (
	"bufio"
	"crypto/tls"
	"github.com/jgonis/norrisBot/messageParser"
	"log"
	"net/textproto"
)

func handleMessages(conn *tls.Conn) {
	connectionReader := textproto.NewReader(bufio.NewReader(conn))
	for {
		line, err := connectionReader.ReadLine()
		if err == nil {
			messageParser.ParseMessage(line)
		} else {
			log.Println("error: ", err)
			break
		}
	}
}
