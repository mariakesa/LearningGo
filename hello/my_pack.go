package main

import (
	"fmt"
	"testpack"
)

func MyExtension() {
	fmt.Println("Hello, World!")
	maria := "Today!"
	fmt.Printf("Rasmus is hacking the world! %s \n", maria)
	testpack.Hello("Rasmus")
}