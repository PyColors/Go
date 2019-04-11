package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

// Declaring a Function Type - do Printer
type Printer func(string)

// type + name
func Greet(salutation Salutation, do Printer, isFormal bool, times int) {
	// 2 multiples possibilities
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)

	i := 0
	for i < times {
		if prefix := GetPrefix(salutation.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
		i++
	}
}

func GetPrefix(name string) (prefix string) {
	switch {
	case name == "Bob":
		prefix = "Mr "
	case name == "Joe", name == "Lisa", len(name) > 10:
		prefix = "Dr "
	case name == "Maria":
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
