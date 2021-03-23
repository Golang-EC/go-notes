package tests

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestSomething(t *testing.T) {
	assert := assert.New(t)

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(a, b, "The two words should be the same.")
}
