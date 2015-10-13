package repl

type Command struct {
	Keyword     string
	Description string
	Action      ContextFn
}
