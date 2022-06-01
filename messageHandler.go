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
			log.Println("Received message from server: ", line)
			message := messageParser.ParseMessage(line)
			handleParsedMessage(message, conn)
		} else {
			log.Println("error: ", err)
			break
		}
	}
}

func handleParsedMessage(message *messageParser.ParsedMessage, conn *tls.Conn) {
	switch message.Command.Command {
	case "PING":
		sendData(conn, "PONG"+message.Command.Parameters, "error replying to PING message with PONG message")
		log.Println("Received PING message, responding with ", "PONG"+message.Command.Parameters)
	case "PRIVMSG":
		if message.Command.BotCommand == "norrisFact" {
			norrisFactMessage := "PRIVMSG " + message.Command.Channel + " :Random Chuck Norris Fact"
			sendData(conn, norrisFactMessage, "error sending random Chuck Norris fact")
			log.Println("received !norrisFact bot message, responding with", norrisFactMessage)
		}
	default:
		log.Println("Received a command that the bot doesn't handle,", message.Command.Command, ", doing nothing")
	}

}
