package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

// Declaring a Function Type - do Printer
type Printer func(string)

// type + name
func Greet(salutation Salutation, do Printer, isFormal bool) {
	// 2 multiples possibilities
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func GetPrefix(name string) (prefix string) {
	switch name {
	case "Bob":
		prefix = "Mr "
	case "Joe":
		prefix = "Sir "
	case "Maria":
		prefix = "Mrs "
	default:
		prefix = "Dude"
	}

	return
}

// Having 2 returns string values (message string, alternate string)
func CreateMessage(name, greeting string) (message string, alternate string) {
	// len() as length
	//fmt.Println(len(greeting))
	message = greeting + " " + name
	alternate = "Hey" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

// Closure is calling the context
// Function to create another function
func CreatePrintFuction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
