package main

import (
	"./greeting"
	"fmt"
)

// RenameToFrog func
func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

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

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Maria", "Why?"},
	}

	// This is a Goroutines that field the Channel

	// Augment the capacity of the slice
	// slice = append(slice, slice...) special one to expand the slice

	// Deleting from a Slice :1 2:
	// slice = append(slice[:1], slice[2:]...)

	// Just pick up Joe...
	// 1 is the start of the index - if empty, will be the index 0
	// 2 is the end of the index - if empty, going to be the same as that slice
	// len(slice) short way to get everything after start of the index
	// slice = slice[1:len(slice)]

	// Rename from the original Object from a `*` (Pointer)
	// salutations[0].Rename("John")

	// salutations has a pointer receiver `*`
	// Deference this with `&` and,
	// Rename the first name of salutation to newName pass in RenameToFrog func.
	// RenameToFrog(&salutations[0])

	// Have the ability to use Salutation Type as an argument to `Fprintf`
	fmt.Fprintf(&salutations[0], "The count is %d", 10)

	// Channel type of salutation
	c := make(chan greeting.Salutation)

	// Second Channel
	c2 := make(chan greeting.Salutation)

	// Goroutines
	go salutations.ChannelGreeter(c)

	// Second Goroutines
	go salutations.ChannelGreeter(c2)

	// Switch to determine what Channel are ready and print out from particular Channel
	// This infinite loop keep doing this select statement
	// If the Select statement can't read to one case he will be the default if there is one
	for {
		// Select statement
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}

		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}

		default:
			fmt.Println("Waiting...")
		}
	}

	//for s := range c {
	//	fmt.Println(s.Name)
	//}

	// Channel
	// `Buffered` of 2
	// The `Buffered` control how much data can be store before we block
	//done := make(chan bool, 2)

	// Simple Goroutines `go`
	//  do to the greeting as a closure
	//go func() {
	//	salutations.Greet(greeting.CreatePrintFuction("<C>"), true, 6)
	//	// The value true going through on the chanel done by the arrow
	//	// direction of the arrow is to the data
	//	done <- true
	//	time.Sleep(100 * time.Millisecond)
	//	done <- true
	//	fmt.Println("Done!")
	//}()

	// New method
	//salutations.Greet(greeting.CreatePrintFuction("?"), true, 6)
	// greeting.Greet(salutations, greeting.CreatePrintFuction("??"), true, 6)

	// Get enough time to run the Goroutines `go`
	// time.Sleep(100 * time.Millisecond)

	// Data
	//<-done
	// Adding time to be able to print "Done" even the Sleep into the anonymous function
	//for {
	//	time.Sleep(100 * time.Millisecond)
	//}
}
