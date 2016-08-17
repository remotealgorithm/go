package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
)

func main() {
	fmt.Print(c.Clear)
	fmt.Println(c.Red + "What's your name?")
	name := input.Ask(c.Blue + "> ")
	fmt.Println(name + " is a really bad name lol")
}
