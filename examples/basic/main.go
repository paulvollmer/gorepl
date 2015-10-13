package main

import "github.com/paulvollmer/gorepl"

func main() {
	// Initialize a new REPL instance.
	app := repl.NewREPL("SAMPLE")
	app.Version = "0.0.1"
	app.Author = "Paul Vollmer"

	// Set the commands and the handlers.
	app.Command("hello", "some words about the 'hello' command", func(ctx repl.Context) {
		ctx.Writeln(ctx.Keyword, "world")
	app.CommandUnknown(func(ctx repl.Context) {
		ctx.Writef("the command '%s' is unknown...\n", ctx.Keyword)
	})

	// Run the REPL as an infinite loop.
	app.Run()
}
