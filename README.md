# Go

## Hello to Go
make a folder hello
```bash
mkdir hello
cd hello
```
init go project
```bash
go mod init hello
```
create a file main.go or hello.go
write this code in the main.go file
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
run the program
```bash
go run main.go
```
you can also use ```go run .``` command to run go file of present directory


important points:
```go.mod``` file is a important file so never ignore this file while pushing on github
```main``` function is the entry point of go programs like cpp

questions:
1. what if i accidentally deleted go.mod file ?
reinitialize the module
```go
go mod init hello
```
run the program
```go
go run main.go
```
2. if we can regenerate go.mod file using ```go mod init <name>``` then why we are not ignoring it.
yes, you can regenerate go.mod file but it is not advisable to ignore it for several reasons.
1. dependency versions
2. collaborative development
3. dependency management
4. ease of use

go.mod file  in a go project in version controls  provide consistend development environment
