package main

import "./greeting"

func main() {

	var s = greeting.Salutation{"Lisa", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFuction("??"), true)
}
