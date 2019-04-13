package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

// Declaring a Function Type - do Printer
type Printer func(string)

// type + name
func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {

	// Go to salutation collection and then, populate i and s
	// i = index, s = the current value
	// _ to skip i
	for _, s := range salutation {
		// 2 multiples possibilities
		message, alternate := CreateMessage(s.Name, s.Greeting)

		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Bob":   "Mr ",
		"Joe":   "Dr ",
		"Lisa":  "Dr ",
		"Maria": "Mrs ",
	}

	return prefixMap[name]

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
