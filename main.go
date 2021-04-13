package main

import (
	"fmt"
	"log"
)

func main() {
	initLoggin("../loggin/logs.log", "logger: ")
	log.Print("Hello logger!!")
	fmt.Println("Hello")
}
