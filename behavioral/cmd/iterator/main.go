package main

import (
	"design-patterns/behavioral/internal/iterator"
	"fmt"
)

func main() {
	// Create users
	user1 := &iterator.User{
		Name: "Vasya",
		Age:  19,
	}
	user2 := &iterator.User{
		Name: "Tanya",
		Age:  19,
	}
	user3 := &iterator.User{
		Name: "Sanya",
		Age:  21,
	}

	// Create collection
	userCollection := iterator.CreateUserCollection([]*iterator.User{user1, user2, user3})

	// Create iterator
	myIterator := userCollection.CreateIterator()

	// Print users from the iterator
	for myIterator.HasNext() {
		user := myIterator.GetNext()
		fmt.Printf("User name: %s, age: %d\n", user.Name, user.Age)
	}
}
