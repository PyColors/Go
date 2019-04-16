package main

import "./greeting"

func main() {

	// var s = greeting.Salutation{"Lisa", "Hello"}

	// `make()` takes three parameters (array, length, capacity)
	// Normally `capacity` will be the same as the length
	// The third parameter is not necessary
	// 3 is the capacity
	//s = make([]int, 3)
	//s[0] = 1
	//s[0] = 10
	//s[0] = 500
	//s[0] = 23

	// Short Slice syntax
	//s := []int{1, 10, 500, 25}

	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Maria", "Why?"},
	}

	// Augment the capacity of the slice
	// slice = append(slice, slice...) special one to expand the slice

	// Deleting from a Slice :1 2:
	slice = append(slice[:1], slice[2:]...)

	// Just pick up Joe...
	// 1 is the start of the index - if empty, will be the index 0
	// 2 is the end of the index - if empty, going to be the same as that slice
	// len(slice) short way to get everything after start of the index
	// slice = slice[1:len(slice)]

	greeting.Greet(slice, greeting.CreatePrintFuction("??"), true, 6)
}
