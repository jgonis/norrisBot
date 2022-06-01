package messageParser

import "testing"

func TestParseCAPACKMessage(t *testing.T) {
	//:tmi.twitch.tv CAP * ACK :twitch.tv/commands
}

func TestParse001Message(t *testing.T) {
	//:tmi.twitch.tv 001 jgonis :Welcome, GLHF!
	expected := ParsedMessage{Command: MessageCommand{Command: "001", Channel: "jgonis", Parameters: "Welcome, GLHF!"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 001 jgonis :Welcome, GLHF!")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}

}

func TestParse002Message(t *testing.T) {
	//:tmi.twitch.tv 002 jgonis :Your host is tmi.twitch.tv
	expected := ParsedMessage{Command: MessageCommand{Command: "002", Channel: "jgonis", Parameters: "Your host is tmi.twitch.tv"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 002 jgonis :Your host is tmi.twitch.tv")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse003Message(t *testing.T) {
	//:tmi.twitch.tv 003 jgonis :This server is rather new
	expected := ParsedMessage{Command: MessageCommand{Command: "003", Channel: "jgonis", Parameters: "This server is rather new"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 003 jgonis :This server is rather new")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse004Message(t *testing.T) {
	//:tmi.twitch.tv 004 jgonis :-
	expected := ParsedMessage{Command: MessageCommand{Command: "004", Channel: "jgonis", Parameters: "-"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 004 jgonis :-")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse375Message(t *testing.T) {
	//:tmi.twitch.tv 375 jgonis :-
	expected := ParsedMessage{Command: MessageCommand{Command: "375", Channel: "jgonis", Parameters: "-"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 375 jgonis :-")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse372Message(t *testing.T) {
	//:tmi.twitch.tv 372 jgonis :You are in a maze of twisty passages, all alike.
	expected := ParsedMessage{Command: MessageCommand{Command: "372", Channel: "jgonis", Parameters: "You are in a maze of twisty passages, all alike."},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 372 jgonis :You are in a maze of twisty passages, all alike.")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse376Message(t *testing.T) {
	//:tmi.twitch.tv 376 jgonis :>
	expected := ParsedMessage{Command: MessageCommand{Command: "376", Channel: "jgonis", Parameters: ">"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 376 jgonis :>")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParseGLOBALUSERSTATEMessage(t *testing.T) {
	//:tmi.twitch.tv GLOBALUSERSTATE
	expected := ParsedMessage{Command: MessageCommand{Command: "GLOBALUSERSTATE"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv GLOBALUSERSTATE")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParseJOINMessage(t *testing.T) {
	//:jgonis!jgonis@jgonis.tmi.twitch.tv JOIN #jgonis
	expected := ParsedMessage{Command: MessageCommand{Channel: "#jgonis", Command: "JOIN"},
		source: MessageSource{nickName: "jgonis", host: "jgonis@jgonis.tmi.twitch.tv"}}
	actual := ParseMessage(":jgonis!jgonis@jgonis.tmi.twitch.tv JOIN #jgonis")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse353Message(t *testing.T) {
	//:jgonis.tmi.twitch.tv 353 jgonis = #jgonis :jgonis
	expected := ParsedMessage{Command: MessageCommand{Channel: "jgonis = #jgonis", Command: "353", Parameters: "jgonis"},
		source: MessageSource{host: "jgonis.tmi.twitch.tv"}}
	actual := ParseMessage(":jgonis.tmi.twitch.tv 353 jgonis = #jgonis :jgonis")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse366Message(t *testing.T) {
	//:jgonis.tmi.twitch.tv 366 jgonis #jgonis :End of /NAMES list
	expected := ParsedMessage{Command: MessageCommand{Channel: "jgonis #jgonis", Command: "366", Parameters: "End of /NAMES list"},
		source: MessageSource{host: "jgonis.tmi.twitch.tv"}}
	actual := ParseMessage(":jgonis.tmi.twitch.tv 366 jgonis #jgonis :End of /NAMES list")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParseUSERSTATEMessage(t *testing.T) {
	//:tmi.twitch.tv USERSTATE #jgonis
	expected := ParsedMessage{Command: MessageCommand{Channel: "#jgonis", Command: "USERSTATE"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv USERSTATE #jgonis")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParseROOMSTATEMessage(t *testing.T) {
	//:tmi.twitch.tv ROOMSTATE #jgonis
	expected := ParsedMessage{Command: MessageCommand{Channel: "#jgonis", Command: "ROOMSTATE"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv ROOMSTATE #jgonis")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParsePRIVMSGMessage(t *testing.T) {
	//:jgonis!jgonis@jgonis.tmi.twitch.tv PRIVMSG #jgonis :a
	expected := ParsedMessage{Command: MessageCommand{Channel: "#jgonis", Command: "PRIVMSG", Parameters: "a"},
		source: MessageSource{nickName: "jgonis", host: "jgonis@jgonis.tmi.twitch.tv"}}
	actual := ParseMessage(":jgonis!jgonis@jgonis.tmi.twitch.tv PRIVMSG #jgonis :a")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParsePRIVMSGBotCommandMessage(t *testing.T) {
	//:jgonis!jgonis@jgonis.tmi.twitch.tv PRIVMSG #jgonis :!dice
	expected := ParsedMessage{Command: MessageCommand{Command: "PRIVMSG", Channel: "#jgonis", BotCommand: "dice", Parameters: "!dice"},
		source: MessageSource{nickName: "jgonis", host: "jgonis@jgonis.tmi.twitch.tv"}}
	actual := ParseMessage(":jgonis!jgonis@jgonis.tmi.twitch.tv PRIVMSG #jgonis :!dice")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParsePINGMessage(t *testing.T) {
	//PING :tmi.twitch.tv
	expected := ParsedMessage{Command: MessageCommand{Command: "PING", Parameters: ":tmi.twitch.tv"}}
	actual := ParseMessage("PING :tmi.twitch.tv")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}
