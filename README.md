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
create a file ```main.go``` or ```hello.go```
write this code in the ```main.go``` file
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

* ```go.mod``` file is a important file so never ignore this file while pushing on github.

* ```main``` function is the entry point of go programs like cpp

#### questions:
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

## External Package
make a directory ```externalPack/``` and a ```main.go``` inside it
```bash
mkdir externalPack/
cd externalPack/
```
init module
```bash
go mod init externalpack.com/hello
```
import external package using import "package_name"
example:
main.go
```go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
```
run the program using ```go run main.go```

output
```text
Hello, World!
I can eat glass and it doesn't hurt me.
Don't communicate by sharing memory, share memory by communicating.
Hello, world.
If a program is too slow, it must have a loop.
```