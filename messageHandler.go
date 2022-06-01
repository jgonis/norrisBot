package main

import (
	"bufio"
	"crypto/tls"
	"github.com/jgonis/norrisBot/messageParser"
	"github.com/jgonis/norrisBot/norrisFact"
	"log"
	"net/textproto"
	"time"
)

func handleMessages(conn *tls.Conn, writeChannel chan<- *SendMessage) {
	connectionReader := textproto.NewReader(bufio.NewReader(conn))
	for {
		line, err := connectionReader.ReadLine()
		if err == nil {
			go parseAndDispatch(line, writeChannel)
		} else {
			log.Println("ERROR: ", err)
			break
		}
	}
}

func parseAndDispatch(line string, writeChannel chan<- *SendMessage) {
	log.Println("RECEIVED:", line)
	message := messageParser.ParseMessage(line)
	handleParsedMessage(message, writeChannel)
}

func handleParsedMessage(message *messageParser.ParsedMessage, writeChannel chan<- *SendMessage) {
	switch message.Command.Command {
	case "PING":
		go handlePongMessage(message.Command.Parameters, writeChannel)
	case "PRIVMSG":
		if message.Command.BotCommand == "norrisFact" {
			go handleNorrisMessage(message.Command.Channel, writeChannel)
		}
	default:
		log.Println("MESSAGE: Parsed a command that the bot doesn't handle,", message.Command.Command, ", doing nothing")
	}
}

func handlePongMessage(messageParameter string, writeChannel chan<- *SendMessage) {
	pongMessage := "PONG " + messageParameter
	SendMessageUnlessFull(writeChannel, &SendMessage{MainMessage: pongMessage,
		ErrorMessage: "error replying to PING message with PONG message"})
}

func handleNorrisMessage(messageChannel string, writeChannel chan<- *SendMessage) {
	norrisFactChannel := make(chan string)
	go norrisFact.GetNorrisFact(norrisFactChannel)
	select {
	case norrisFactPayload := <-norrisFactChannel:
		norrisFactMessage := "PRIVMSG " + messageChannel + " :" + norrisFactPayload
		SendMessageUnlessFull(writeChannel, &SendMessage{MainMessage: norrisFactMessage,
			ErrorMessage: "error sending random Chuck Norris fact"})
	case <-time.After(5 * time.Second):
		log.Println("ERROR: Timed out trying to receive Norris fact")
	}

}
