package repl

import (
	"fmt"
)

type Context struct {
	Keyword string
	History string
}

func (c *Context) Write(a ...interface{}) {
	fmt.Print(a...)
}

func (c *Context) Writef(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (c *Context) Writeln(a ...interface{}) {
	fmt.Println(a...)
}

type ContextFn func(Context)
