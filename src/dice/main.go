//basic setup things
//note: hiccups are terrible
package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//define dice
	one := `
	 -------
	|       |
	|   •   |
	|       |
	 -------
	`
	two := `
	 -------
	| •     |
	|       |
	|     • | 
	 -------
	`
	three := `
	 -------
	| •     |
	|   •   |
	|     • |
	 -------
	`
	four := `
	 -------
	| •   • |
	|       |
	| •   • |
	 -------
	`
	five := ` 
	 -------
	| •   • |
	|   •   |
	| •   • |
	 -------
	`
	six := `
	  -------
	 | •   • |
	 | •   • |
	 | •   • |
	  -------
	`
	sides := []string{one, two, three, four, five, six}
	//define the thingamajig
	roll := rand.Intn(6)
	//roll
	fmt.Println(c.Rc() + sides[roll])
}
