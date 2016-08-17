package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
)

func main() {
	fmt.Print(c.Clear)
	for {
		fmt.Print(c.Rc() + "NYAN ")

	}
}
