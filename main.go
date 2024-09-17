package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	message := Hello("World")
	fmt.Println(message)
}
