# Go

## Hello World Program
### Program Creation steps
### Create and Run the program
make a directory and a file called main.go inside the directory
```bash
mkdir hello
cd hello/
code .
```
create a file ```main.go```
### initialize go module
```bash
go mod init hello
```
Go has entry point which is ```main``` same like c and c++.
### write this code in ```main.go```
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello from Bugs Founder")
}
```
### run the file
```bash
go run main.go
```


### Variables and Types 
declaring a variable

```go
// var <var_name> <DT_TYPE> = <value>
var username string = "BugsFounder"
fmt.Println(username)
fmt.Printf("Variable is of type: %T \n", username)

// boolean
var isLoggedIn bool = false
fmt.Println(isLoggedIn)
fmt.Printf("Variable is of type: %T \n", isLoggedIn)

// integers
// var smallVal uint8 = 256 // cannot use 256 (untyped int constant) as uint8 value in variable declaration (overflows)
var smallVal uint8 = 255
fmt.Println(smallVal)
fmt.Printf("variable is of type: %T \n", smallVal)

// float
// var smallFloat float32 = 255.32434343434343434545 // 255.32434
var smallFloat float64 = 255.32434343434343434545 // 255.32434343434343
fmt.Println(smallFloat)
fmt.Printf("Variable is of type: %T \n", smallFloat)
```
### Other important things
#### default values and some aliases
```go
var anotherVariable int // type int value 0
// var anotherVariable string // type string value ''
fmt.Println(anotherVariable)
fmt.Printf("Variable is of type: %T \n", anotherVariable) 
```
#### Implicit type
```go
	var website = "bugsfounder.com"
	fmt.Println(website)
``` 
#### Walrus operator 
```go
// no var style
numberOfUser := 300000
fmt.Println(numberOfUser)
```

#### Walrus operator is not allowed outside of the method
```go
package main

import "fmt"

// jwtToken := 300000 // error
// var jwtToken int = 300000 // ok

// public variable because first letter is capital "LoginToken"
const LoginToken string = "sadfassdlksdwerweuiwer" // Public

func main() {

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}

```
Note: if you've created a variable or imported something you must have to use it otherwise you're going to get an error
### Taking input from the user

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok || error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Thanks for rating %T\n", input)
}
```

### Conversion
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to our pizza app"
	fmt.Println(welcome)
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	// this will also work but not preferable
	// inp, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// fmt.Println(inp)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

	// conversion 
	
	// strconv.ParseFloat: parsing "4\n": invalid syntax
	numRating, err := strconv.ParseFloat(input, 64) // error

	// correct way
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
```

### Time and Date 
time module

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday")) // 02 -> day, 01 -> month, 2006 -> year, Monday --> day, casesensitive it has to be Monday not monday

	// learn more : https://go.dev/play/p/d_92jfpw8Xq

	createDate := time.Date(2020, time.August, 15, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
	

}
```

while working with dates and times formatters in go you always have to maintain the strings as constant which is given in the go doc. for Ex:- ```02 -> day, 01 -> month, 2006 -> year, Monday --> weekday name```, casesensitive it has to be ```Monday``` not ```monday```
* [Learn More in go Playground](https://go.dev/play/p/d_92jfpw8Xq)
* [Time Doc](https://pkg.go.dev/time#Time)
