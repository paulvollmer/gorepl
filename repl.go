package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type REPL struct {
	Name           string
	Prompt         string
	Version        string
	Author         string
	AuthorEmail    string
	WelcomeMessage string
	Commands       []Command
	commandUnknown Command
	History        []string
}

func NewREPL(name string) *REPL {
	r := REPL{}
	// Set some default values
	r.Name = name
	r.Prompt = name + ">"
	r.Version = "0.0.0"
	r.WelcomeMessage = "Welcome to the " + name + " REPL (" + r.Version + ")\nType 'help' to get a list of all commands.\nType 'exit' to get out here...\n\n"
	// Set some default commands
	r.Command("help", "display this text", func(ctx Context) {
		ctx.Writeln(r.GetHelpText())
	})
	r.Command("version", "version of the application", func(ctx Context) {
		ctx.Writeln(r.Version)
	})
	r.Command("exit", "exit the application", func(ctx Context) {
		ctx.Writeln("Bye bye...")
		os.Exit(0)
	})
	r.Command("author", "the author", func(ctx Context) {
		ctx.Writeln(r.Author)
	})
	r.Command("author-email", "the author email address", func(ctx Context) {
		ctx.Writeln(r.AuthorEmail)
	})
	return &r
}

func (r *REPL) GetHelpText() string {
	commandList := "help:\n"
	for _, v := range r.Commands {
		commandList += "- " + v.Keyword + " --- " + v.Description + "\n"
	}
	return commandList
}

func (r *REPL) Command(keyword, description string, action ContextFn) {
	c := Command{keyword, description, action}
	r.Commands = append(r.Commands, c)
}

func (r *REPL) CommandUnknown(action ContextFn) {
	c := Command{"unknown", "unknown command", action}
	r.commandUnknown = c
}

func (r *REPL) Run() {
	fmt.Print(r.WelcomeMessage)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(r.Prompt + " ")
		textInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("REPL-ERROR!", err)
		}
		text := strings.TrimSpace(textInput)
		r.History = append(r.History, text)

		// Check if the command exist...
		commandFound := false
		for _, v := range r.Commands {
			if v.Keyword == text {
				ctx := Context{}
				ctx.Keyword = v.Keyword
				ctx.History = r.History
				commandFound = true
				v.Action(ctx)
			}
		}

		if !commandFound {
			ctx := Context{}
			ctx.Keyword = text
			r.commandUnknown.Action(ctx)
		}
	}
}
