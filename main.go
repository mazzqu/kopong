package main

import "fmt"

func main() {
	fmt.Println(greeting("lord jengkloz"))
}

func greeting(name string) string {
	return "Hello " + name
}
