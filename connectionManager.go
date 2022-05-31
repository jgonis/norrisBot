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
	go handleMessages(conn)
	authenticateAndJoinChannel(conn, userNameFlag, channelNameFlag, oauthToken)
	quitChan := waitForSigTerm()
	<-quitChan
	disconnect(conn)
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

func authenticateAndJoinChannel(conn *tls.Conn, userName *string, channelName *string, oauthToken *string) {
	sendData(conn,
		"CAP REQ :twitch.tv/commands",
		"error send CAP request")
	sendData(conn,
		"PASS oauth:"+*oauthToken,
		"error sending oauth token")
	sendData(conn,
		"NICK "+*userName,
		"error sending user name")
	sendData(conn,
		"JOIN #"+*channelName,
		"error joining channel")
}

func disconnect(conn net.Conn) {
	sendData(conn,
		"QUIT Bye",
		"error sending QUIT message")
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
