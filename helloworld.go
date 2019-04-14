package main

import "./greeting"

func main() {

	// var s = greeting.Salutation{"Lisa", "Hello"}

	var s []int
	// `make()` takes three parameters (array, length, capacity)
	// Normally `capacity` will be the same as the length
	// The third parameter is not necessary
	// 3 is the capacity
	s = make([]int, 3)
	s[0] = 1
	s[0] = 10
	s[0] = 500
	s[0] = 23

	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Maria", "Why?"},
	}

	greeting.Greet(slice, greeting.CreatePrintFuction("??"), true, 6)
}
