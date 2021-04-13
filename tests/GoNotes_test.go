package test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {

}

func TestSomething(t *testing.T) {
	assert := assert.New(t)

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(a, b, "The two words should be the same.")
}

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")
	os.Exit(exitVal)
}

func TestLoggin(t *testing.T) {

	// f := InitLoggin("../loggin/logs.log", "logger: ")
	// file, err := os.OpenFile(f)
	// // defer CloseLogFile(f)
	// if f != file {
	// 	t.Errorf("No se creo el archivo de loggin", err)
	// }
}
