package main

import "fmt"

func main() {
	fmt.Println("Hello, sdWorld!")
}

/*
Five basic questions:
1. How to run the code in project?
-> go run main.go -> Runs the code in main.go
-> go build -> Compiles the code in main.go
-> go run -> Runs the compiled code
-> go install -> Compiles the code and installs the binary in $GOPATH/bin
-> go fmt -> Formats the code
-> go get -> Downloads the package
-> go test - > Runs the test

2. What does package main mean?
Package == Project == Workspace
Package is a collection of source files in the same directory that are compiled together.
There are two types of packages: executable and reusable.
Reusable packages are used as dependencies for other packages.
Word main is used to create an executable package.

3. What does import "fmt" mean?
Grant access to the fmt package.
fmt is a standard library package that provides formatting for input and output.

4. How are main.go file organized in a Go project?

5. What is func mean?
*/
