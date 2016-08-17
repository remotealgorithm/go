package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
)

func main() {
	fmt.Print(c.Clear)
	fmt.Println(c.Red + "QUIZ!")
	waffles := i.Ask(c.Blue + "Do you like waffles? ")
	if waffles == "yes" {
		if pancakes := i.Ask(c.Rc() + "Do you like pancakes? "); pancakes == "yes" {
			if frenchtoast := i.Ask("Do you like french toast? "); frenchtoast == "yes" {
				fmt.Println("Can't wait to get a mouthful")
			}
		}
	} else {
		fmt.Println(c.Violet + "Well, we like waffles!")
	}
}
