package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

// Type of Salutations by nameType of slice of salutation
type Salutations []Salutation

// Declaring a Function Type - do Printer
type Printer func(string)

// Method of salutation type
func (salutations Salutations) Greet(do Printer, isFormal bool, times int) {

	// Go to salutation collection and then, populate i and s
	// i = index, s = the current value
	// _ to skip i
	for _, s := range salutations {
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

	// Simple way to write a map
	// Easy to read and maintain
	// key is a string and the type too
	prefixMap := map[string]string{
		"Bob":   "Mr ",
		"Joe":   "Dr ",
		"Lisa":  "Dr ",
		"Maria": "Mrs ",
	}

	// Update key in the map, if he exists
	prefixMap["Joe"] = "Jr "

	// Info: this the old way to delete a key in the map
	// prefixMap["Maria"] = "", false

	// Delete a key in the map in a new way
	// Delete would work even if the value is in the map or not
	// Don't need to verify in before doing a delete or an update with `Go`
	delete(prefixMap, "Maria")

	// As we can have multiple return in `Go`
	// exists is a boolean, he is actually checking if the value exists
	// Don't need to write if `exists === true` or `...false` as he a boolean itself
	if value, exists := prefixMap[name]; exists {
		return value
	}

	// return default value
	return "Dude "

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
