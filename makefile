hello: 
	echo "hello"

test:
	go test src/GoNotes_test.go

build: 
	go build -o bin/main src/*.go

run:
	go run src/main.go

coverage: 
	go test src/GoNotes_test.go -race -covermode=atomic --coverprofile=coverage.out

check:
	go vet src/*.go