package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	userNameFlag := flag.String("userName", "", "The user name the bot will use to authenticate itself with")
	channelNameFlag := flag.String("channelName", "", "The name of the channel the bot will join")
	oauthToken := flag.String("oauthToken", "", "The oauth token the bot will use to authenticate with")

	flag.Parse()

	validCommandLine := true
	validateCLArg(userNameFlag, &validCommandLine, "Expected to receive a user name to authenticate with the server")
	validateCLArg(channelNameFlag, &validCommandLine, "Expected to receive a user name to authenticate with the server")
	validateCLArg(oauthToken, &validCommandLine, "Expected to receive a user name to authenticate with the server")
	if validCommandLine == false {
		log.Fatal("Did not receive valid args to run the bot with. Exiting")
	}

	conn := createIrcConnection()

	//refreshToken()
	authenticateAndJoinChannel(conn, userNameFlag, channelNameFlag, oauthToken)
	handleMessages(conn)

	quitChan := waitForSigTerm()
	<-quitChan
	disconnect(conn)
	fmt.Println("exiting program")
}

func validateCLArg(arg *string, validCommandLine *bool, errorMessage string) {
	if len(*arg) == 0 {
		log.Println(errorMessage)
		*validCommandLine = false
	}
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

func handleMessages(conn *tls.Conn) {
	connectionReader := textproto.NewReader(bufio.NewReader(conn))
	for {
		line, err := connectionReader.ReadLine()
		if err == nil {
			parseMessage(line)
		} else {
			log.Println("error: ", err)
			break
		}
	}
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

func disconnect(conn net.Conn) {
	sendData(conn,
		"QUIT Bye",
		"error sending QUIT message")
	conn.Close()
}

func sendData(conn net.Conn, message string, errMessage string) {
	_, err := fmt.Fprintf(conn, "%s\r\n", message)
	if err != nil {
		log.Println(errMessage, err)
	}
}
