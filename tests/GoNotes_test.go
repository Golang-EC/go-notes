package GoNotes_test

import (
	"fmt"
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

	loggin.initLoggin("../loggin/logs.log", "logger: ")
	log.Println("Do stuff BEFORE the tests!")
	fmt.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	fmt.Println("Do stuff AFTER the tests!")
	Log.Println("Do stuff AFTER the tests!")
	os.Exit(exitVal)
}
