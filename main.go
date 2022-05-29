package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello World!")
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	_, err := tls.Dial("tcp", "irc.chat.twitch.tv:6697", config)
	if err != nil {
		log.Fatal(err)
	}
}
