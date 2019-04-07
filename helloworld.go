package main

import "./greeting"

func main() {

	var s = greeting.Salutation{"M	aria", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFuction("??"), true)
}
