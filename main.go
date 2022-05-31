package main

import (
	"flag"
	"fmt"
	"log"
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

	//refreshToken()
	startupBot(userNameFlag, channelNameFlag, oauthToken)
	fmt.Println("exiting program")
}

func validateCLArg(arg *string, validCommandLine *bool, errorMessage string) {
	if len(*arg) == 0 {
		log.Println(errorMessage)
		*validCommandLine = false
	}
}
