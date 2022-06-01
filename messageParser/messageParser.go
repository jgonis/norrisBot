package messageParser

import (
	"log"
	"strings"
)

type MessageSource struct {
	nickName string
	host     string
}

type MessageCommand struct {
	Command    string
	Channel    string
	BotCommand string
	Parameters string
}

type ParsedMessage struct {
	Command MessageCommand
	source  MessageSource
}

func ParseMessage(line string) *ParsedMessage {
	var parsedMessage *ParsedMessage
	if strings.HasPrefix(line, "PING") {
		parsedMessage = parsePingMessage(line)
		log.Println("received PING message, responding with: ", "PONG"+parsedMessage.Command.Parameters)
	} else if strings.HasPrefix(line, ":") {
		trimmedLine := line[1:]
		splitStrings := strings.SplitN(trimmedLine, " ", 2)
		messageSource := parseSourceString(splitStrings[0])
		messageCommand := parseCommandString(splitStrings[1])
		parsedMessage = &ParsedMessage{Command: *messageCommand, source: *messageSource}
	} else {
		log.Println("received a message")
	}
	return parsedMessage
}

func parsePingMessage(line string) *ParsedMessage {
	returnMessage := strings.TrimSpace(strings.TrimPrefix(line, "PING"))
	return &ParsedMessage{Command: MessageCommand{Command: "PING", Parameters: returnMessage}}
}

func parseCommandString(commandString string) *MessageCommand {
	messageCommand := MessageCommand{}
	commandAndParams := strings.Split(commandString, ":")
	if len(commandAndParams) == 1 {
		//there might be no params or the Command might have no args such as the GLOBALUSERSTATE message
		//Figure out which case we're in
		commandAndArgs := strings.Split(strings.TrimSpace(commandAndParams[0]), " ")
		if len(commandAndArgs) == 1 {
			messageCommand.Command = commandAndArgs[0]
		} else {
			messageCommand.Command = commandAndArgs[0]
			messageCommand.Channel = commandAndArgs[1]
		}
	} else {
		commandAndChannel := strings.SplitN(strings.TrimSpace(commandAndParams[0]), " ", 2)
		messageCommand.Command = commandAndChannel[0]
		messageCommand.Channel = commandAndChannel[1]
		messageCommand.Parameters = commandAndParams[1]
		//check to see if the parameter is bot Command, aka a string prefixed with !
		if strings.HasPrefix(commandAndParams[1], "!") {
			messageCommand.BotCommand = commandAndParams[1][1:]
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
