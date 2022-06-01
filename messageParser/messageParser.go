package messageParser

import (
	"fmt"
	"log"
	"strings"
)

type MessageSource struct {
	nickName string
	host     string
}

type MessageCommand struct {
	command    string
	channel    string
	botCommand string
	parameters string
}

type ParsedMessage struct {
	command MessageCommand
	source  MessageSource
}

func ParseMessage(line string) *ParsedMessage {
	fmt.Println("Message received was: ", line)
	var parsedMessage *ParsedMessage
	if strings.HasPrefix(line, "PING") {
		parsedMessage = parsePingMessage(line)
		log.Println("received PING message, responding with: ", "PONG"+parsedMessage.command.parameters)
		//sendData(conn, "P
		//ONG"+parsedMessage.parameters, "error replying to PING message with PONG message")
	} else if strings.HasPrefix(line, ":") {
		trimmedLine := line[1:]
		splitStrings := strings.SplitN(trimmedLine, " ", 2)
		messageSource := parseSourceString(splitStrings[0])
		messageCommand := parseCommandString(splitStrings[1])
		parsedMessage = &ParsedMessage{command: *messageCommand, source: *messageSource}
	} else {
		log.Println("received a message")
	}
	return parsedMessage
}

func parsePingMessage(line string) *ParsedMessage {
	returnMessage := strings.TrimSpace(strings.TrimPrefix(line, "PING"))
	return &ParsedMessage{command: MessageCommand{command: "PING", parameters: returnMessage}}
}

func parseCommandString(commandString string) *MessageCommand {
	messageCommand := MessageCommand{}
	commandAndParams := strings.Split(commandString, ":")
	if len(commandAndParams) == 1 {
		//there might be no params or the command might have no args such as the GLOBALUSERSTATE message
		//Figure out which case we're in
		commandAndArgs := strings.Split(strings.TrimSpace(commandAndParams[0]), " ")
		if len(commandAndArgs) == 1 {
			messageCommand.command = commandAndArgs[0]
		} else {
			messageCommand.command = commandAndArgs[0]
			messageCommand.channel = commandAndArgs[1]
		}
	} else {
		commandAndChannel := strings.SplitN(strings.TrimSpace(commandAndParams[0]), " ", 2)
		messageCommand.command = commandAndChannel[0]
		messageCommand.channel = commandAndChannel[1]
		messageCommand.parameters = commandAndParams[1]
		//check to see if the parameter is bot command, aka a string prefixed with !
		if strings.HasPrefix(commandAndParams[1], "!") {
			messageCommand.botCommand = commandAndParams[1][1:]
		}
	}
	return &messageCommand
}

func parseSourceString(sourceString string) *MessageSource {
	messageSource := MessageSource{}
	sourceStringParts := strings.Split(sourceString, "!")
	if len(sourceStringParts) == 2 {
		messageSource.nickName = sourceStringParts[0]
		messageSource.host = sourceStringParts[1]
	} else {
		messageSource.host = sourceStringParts[0]
	}
	return &messageSource
}
