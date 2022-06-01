package messageParser

import "testing"

func TestParseCAPACKMessage(t *testing.T) {
	//:tmi.twitch.tv CAP * ACK :twitch.tv/commands
}

func TestParse001Message(t *testing.T) {
	//:tmi.twitch.tv 001 jgonis :Welcome, GLHF!
	expected := ParsedMessage{command: MessageCommand{command: "001", channel: "jgonis", parameters: "Welcome, GLHF!"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 001 jgonis :Welcome, GLHF!")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}

}

func TestParse002Message(t *testing.T) {
	//:tmi.twitch.tv 002 jgonis :Your host is tmi.twitch.tv
	expected := ParsedMessage{command: MessageCommand{command: "002", channel: "jgonis", parameters: "Your host is tmi.twitch.tv"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 002 jgonis :Your host is tmi.twitch.tv")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse003Message(t *testing.T) {
	//:tmi.twitch.tv 003 jgonis :This server is rather new
	expected := ParsedMessage{command: MessageCommand{command: "003", channel: "jgonis", parameters: "This server is rather new"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 003 jgonis :This server is rather new")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse004Message(t *testing.T) {
	//:tmi.twitch.tv 004 jgonis :-
	expected := ParsedMessage{command: MessageCommand{command: "004", channel: "jgonis", parameters: "-"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 004 jgonis :-")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse375Message(t *testing.T) {
	//:tmi.twitch.tv 375 jgonis :-
	expected := ParsedMessage{command: MessageCommand{command: "375", channel: "jgonis", parameters: "-"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 375 jgonis :-")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse372Message(t *testing.T) {
	//:tmi.twitch.tv 372 jgonis :You are in a maze of twisty passages, all alike.
	expected := ParsedMessage{command: MessageCommand{command: "372", channel: "jgonis", parameters: "You are in a maze of twisty passages, all alike."},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 372 jgonis :You are in a maze of twisty passages, all alike.")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse376Message(t *testing.T) {
	//:tmi.twitch.tv 376 jgonis :>
	expected := ParsedMessage{command: MessageCommand{command: "376", channel: "jgonis", parameters: ">"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv 376 jgonis :>")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParseGLOBALUSERSTATEMessage(t *testing.T) {
	//:tmi.twitch.tv GLOBALUSERSTATE
	expected := ParsedMessage{command: MessageCommand{command: "GLOBALUSERSTATE"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv GLOBALUSERSTATE")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParseJOINMessage(t *testing.T) {
	//:jgonis!jgonis@jgonis.tmi.twitch.tv JOIN #jgonis
	expected := ParsedMessage{command: MessageCommand{channel: "#jgonis", command: "JOIN"},
		source: MessageSource{nickName: "jgonis", host: "jgonis@jgonis.tmi.twitch.tv"}}
	actual := ParseMessage(":jgonis!jgonis@jgonis.tmi.twitch.tv JOIN #jgonis")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse353Message(t *testing.T) {
	//:jgonis.tmi.twitch.tv 353 jgonis = #jgonis :jgonis
	expected := ParsedMessage{command: MessageCommand{channel: "jgonis = #jgonis", command: "353", parameters: "jgonis"},
		source: MessageSource{host: "jgonis.tmi.twitch.tv"}}
	actual := ParseMessage(":jgonis.tmi.twitch.tv 353 jgonis = #jgonis :jgonis")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParse366Message(t *testing.T) {
	//:jgonis.tmi.twitch.tv 366 jgonis #jgonis :End of /NAMES list
	expected := ParsedMessage{command: MessageCommand{channel: "jgonis #jgonis", command: "366", parameters: "End of /NAMES list"},
		source: MessageSource{host: "jgonis.tmi.twitch.tv"}}
	actual := ParseMessage(":jgonis.tmi.twitch.tv 366 jgonis #jgonis :End of /NAMES list")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParseUSERSTATEMessage(t *testing.T) {
	//:tmi.twitch.tv USERSTATE #jgonis
	expected := ParsedMessage{command: MessageCommand{channel: "#jgonis", command: "USERSTATE"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv USERSTATE #jgonis")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParseROOMSTATEMessage(t *testing.T) {
	//:tmi.twitch.tv ROOMSTATE #jgonis
	expected := ParsedMessage{command: MessageCommand{channel: "#jgonis", command: "ROOMSTATE"},
		source: MessageSource{host: "tmi.twitch.tv"}}
	actual := ParseMessage(":tmi.twitch.tv ROOMSTATE #jgonis")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParsePRIVMSGMessage(t *testing.T) {
	//:jgonis!jgonis@jgonis.tmi.twitch.tv PRIVMSG #jgonis :a
	expected := ParsedMessage{command: MessageCommand{channel: "#jgonis", command: "PRIVMSG", parameters: "a"},
		source: MessageSource{nickName: "jgonis", host: "jgonis@jgonis.tmi.twitch.tv"}}
	actual := ParseMessage(":jgonis!jgonis@jgonis.tmi.twitch.tv PRIVMSG #jgonis :a")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParsePRIVMSGBotCommandMessage(t *testing.T) {
	//:jgonis!jgonis@jgonis.tmi.twitch.tv PRIVMSG #jgonis :!dice
	expected := ParsedMessage{command: MessageCommand{command: "PRIVMSG", channel: "#jgonis", botCommand: "dice", parameters: "!dice"},
		source: MessageSource{nickName: "jgonis", host: "jgonis@jgonis.tmi.twitch.tv"}}
	actual := ParseMessage(":jgonis!jgonis@jgonis.tmi.twitch.tv PRIVMSG #jgonis :!dice")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}

func TestParsePINGMessage(t *testing.T) {
	//PING :tmi.twitch.tv
	expected := ParsedMessage{command: MessageCommand{command: "PING", parameters: ":tmi.twitch.tv"}}
	actual := ParseMessage("PING :tmi.twitch.tv")
	if expected != *actual {
		t.Errorf("parsed string did not match expected structure")
	}
}
