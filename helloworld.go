package main

import "./greeting"

func main() {

	// var s = greeting.Salutation{"Lisa", "Hello"}

	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Maria", "Why?"},
	}

	greeting.Greet(slice, greeting.CreatePrintFuction("??"), true, 6)
}
