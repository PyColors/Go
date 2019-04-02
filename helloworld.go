package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

// type + name
func Greet(salutation Salutation) {
	// 2 multiples possibilities
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	fmt.Println(message)
	fmt.Println(alternate)
}

// Having 2 returns string values (message string, alternate string)
func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey" + name
	return
}

func main() {

	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
