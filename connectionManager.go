package main

import (
	"crypto/tls"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startupBot(userNameFlag *string, channelNameFlag *string, oauthToken *string) {
	conn := createIrcConnection()
	writeChannel := make(chan SendMessage, 20)
	go handleMessages(conn, writeChannel)
	go sendData(conn, writeChannel)
	authenticateAndJoinChannel(writeChannel, userNameFlag, channelNameFlag, oauthToken)
	quitChan := waitForSigTerm()
	<-quitChan
	disconnect(conn, writeChannel)
}

func createIrcConnection() *tls.Conn {
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second},
		"tcp",
		"irc.chat.twitch.tv:6697",
		config)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func authenticateAndJoinChannel(writeChannel chan SendMessage, userName *string, channelName *string, oauthToken *string) {
	writeChannel <- SendMessage{MainMessage: "CAP REQ :twitch.tv/commands", ErrorMessage: "error send CAP request"}
	writeChannel <- SendMessage{MainMessage: "PASS oauth:" + *oauthToken, ErrorMessage: "error sending oauth token"}
	writeChannel <- SendMessage{MainMessage: "NICK " + *userName, ErrorMessage: "error sending user name"}
	writeChannel <- SendMessage{MainMessage: "JOIN #" + *channelName, ErrorMessage: "error joining channel"}
}

func disconnect(conn net.Conn, writeChannel chan SendMessage) {
	writeChannel <- SendMessage{MainMessage: "QUIT Bye", ErrorMessage: "error sending QUIT message"}
	close(writeChannel)
	<-time.After(1 * time.Second)
	conn.Close()
}

func waitForSigTerm() chan bool {
	quitChan := make(chan bool)
	sigHaltChan := make(chan os.Signal, 2)
	signal.Notify(sigHaltChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigHaltChan
		log.Println("received halt message, quitting")
		quitChan <- true
	}()

	return quitChan
}
