package repl

import (
	"bufio"
	"fmt"
	"os"
)

type REPL struct {
	Name           string
	Prompt         string
	Version        string
	Author         string
	AuthorEmail    string
	WelcomeMessage string
	Commands       []Command
}

func NewREPL(name string) *REPL {
	r := REPL{}
	r.Name = name
	r.Prompt = name + ">"
	r.Version = "0.0.0"
	r.WelcomeMessage = "Welcome to the " + name + " REPL (" + r.Version + ")\nType 'help' to get a list of all commands.\nType 'exit' to get out here...\n\n"
	return &r
}

func (r *REPL) GetHelpText() string {
	commandList := ""
	for _, v := range r.Commands {
		commandList += "- " + v.Keyword + " --- " + v.Description + "\n"
	}
	return `help:
- help         display this text
- author       the author
- author-email the author email address
- exit         exit the application
- version      version of the application
` + commandList
}

func (r *REPL) Command(keyword, description string, action ContextFn) {
	c := Command{keyword, description, action}
	r.Commands = append(r.Commands, c)
}

func (r *REPL) Run() {
	fmt.Print(r.WelcomeMessage)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(r.Prompt + " ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("REPL-ERROR!", err)
		}

		// Check basic commands
		switch text {
		case "help\n":
			fmt.Println(r.GetHelpText())
			break
		case "version\n":
			fmt.Println(r.Version)
			break
		case "exit\n":
			fmt.Println("Bye bye...")
			os.Exit(0)
			break
		default:
			// println("default...")
		}

		// Check custom commands
		for _, v := range r.Commands {
			if v.Keyword+"\n" == text {
				ctx := Context{}
				ctx.Keyword = v.Keyword
				ctx.History = "TODO: history"
				v.Action(ctx)
			}
		}
	}
}
