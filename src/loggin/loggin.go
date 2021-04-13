package main

import (
	"log"
	"os"
)

type Logger struct {
	prefix string
	flag   int
}

func initLoggin(outPath string, prefix string) {
	file, err := os.OpenFile(outPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("No se pudo crear el archivo -> ", outPath, " -> ", err)
	}

	defer file.Close()
	log.SetOutput(file)
	log.SetPrefix(prefix)
	log.Print(prefix + "Inicia Loggin en Go!")
}
