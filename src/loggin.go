package main

import (
	"log"
	"os"
)

func InitLoggin(outPath string, prefix string) *os.File {
	file, err := os.OpenFile(outPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("No se pudo crear el archivo -> ", outPath, " -> ", err)
	}

	log.SetOutput(file)
	log.SetPrefix(prefix)
	log.Print(prefix + "Inicia Loggin en Go!")
	return file
}

func CloseLogFile(file *os.File) {
	defer file.Close()
}
