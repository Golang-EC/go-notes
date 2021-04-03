package main

import (
	"fmt"
)

func main() {
	// loggin.logger :=
	loggin.initLoggin("../loggin/logs.log", "logger: ")
	log.Print("Hello logger!!")
	fmt.Println("Hello")
}
