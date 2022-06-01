package main

import (
	"bufio"
	"crypto/tls"
	"github.com/jgonis/norrisBot/messageParser"
	"github.com/jgonis/norrisBot/norrisFact"
	"log"
	"net/textproto"
)

func handleMessages(conn *tls.Conn, writeChannel chan SendMessage) {
	connectionReader := textproto.NewReader(bufio.NewReader(conn))
	for {
		line, err := connectionReader.ReadLine()
		if err == nil {
			log.Println("Received message from server: ", line)
			message := messageParser.ParseMessage(line)
			handleParsedMessage(message, writeChannel)
		} else {
			log.Println("error: ", err)
			break
		}
	}
}

func handleParsedMessage(message *messageParser.ParsedMessage, writeChannel chan SendMessage) {
	switch message.Command.Command {
	case "PING":
		SendMessageUnlessFull(writeChannel, SendMessage{MainMessage: "PONG" + message.Command.Parameters,
			ErrorMessage: "error replying to PING message with PONG message"})
		log.Println("Received PING message, responding with ", "PONG"+message.Command.Parameters)
	case "PRIVMSG":
		if message.Command.BotCommand == "norrisFact" {
			norrisFactMessage := "PRIVMSG " + message.Command.Channel + " :" + norrisFact.GetNorrisFact()
			SendMessageUnlessFull(writeChannel, SendMessage{MainMessage: norrisFactMessage,
				ErrorMessage: "error sending random Chuck Norris fact"})
			log.Println("received !norrisFact bot message, responding with", norrisFactMessage)
		}
	default:
		log.Println("Received a command that the bot doesn't handle,", message.Command.Command, ", doing nothing")
	}

}
