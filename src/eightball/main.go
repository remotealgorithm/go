package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//prep
	fmt.Print(c.Clear)
	rand.Seed(time.Now().UnixNano())
	//vars
	answers := []string{"YES.", "NO.", "MAYBE.", "Out to lunch, try again later!", "please help me I'm being attacked by space monsters"}
	for {
		question := i.Ask(c.Violet + "Please ask a question.\n(lowercase only please)\n> ")
		if strings.HasPrefix(question, "hi") {
			fmt.Println(c.Blue + "Friends is an American television sitcom, created by David Crane and Marta Kauffman, which originally aired on NBC from September 22, 1994, to May 6, 2004")
		} else if strings.HasPrefix(question, "hello") {
			fmt.Println(c.Blue + "it's time")
			time.Sleep(2 * time.Second)
			fmt.Println("to take out " + c.Red + "the trash")
		} else {
			fmt.Println(c.Blue + answers[rand.Intn(len(answers))])
		}
	}
}
