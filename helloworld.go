package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

// type + name
func Greet(salutation Salutation, do func(string)) {
	// 2 multiples possibilities
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "coco")
	do(message)
	do(alternate)
}

// Having 2 returns string values (message string, alternate string)
func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	// len() as length
	fmt.Println(len(greeting))
	message = greeting[1] + " " + name
	alternate = "Hey" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func main() {

	var s = Salutation{"Bob", "Hello"}
	Greet(s, Print)
}
