# Go

### Hello World Program
#### Program Creation steps
#### Create and Run the program
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

### Build executables for different operating systems
open the terminal and move to the go project directory and run below commands:
```bash
go env
```
output:
```bash
GO111MODULE=''
GOARCH='amd64'
GOBIN=''
GOCACHE='/home/bugsfounder/.cache/go-build'
GOENV='/home/bugsfounder/.config/go/env'
GOEXE=''
GOEXPERIMENT=''
GOFLAGS=''
GOHOSTARCH='amd64'
GOHOSTOS='linux'
GOINSECURE='' 
GOMODCACHE='/home/bugsfounder/go/pkg/mod'
GONOPROXY=''
GONOSUMDB='' 
GOOS='linux' # this line will help
GOPATH='/home/bugsfounder/go'
GOPRIVATE=''
GOPROXY='https://proxy.golang.org,direct'
GOROOT='/usr/local/go'
GOSUMDB='sum.golang.org'
GOTMPDIR=''
GOTOOLCHAIN='auto'
GOTOOLDIR='/usr/local/go/pkg/tool/linux_amd64'
GOVCS=''
GOVERSION='go1.22.4'
GCCGO='gccgo'
GOAMD64='v1'
AR='ar'
CC='gcc'
CXX='g++'
CGO_ENABLED='1'
GOMOD='/dev/null'
GOWORK=''
CGO_CFLAGS='-O2 -g'
CGO_CPPFLAGS=''
CGO_CXXFLAGS='-O2 -g'
CGO_FFLAGS='-O2 -g'
CGO_LDFLAGS='-O2 -g'
PKG_CONFIG='pkg-config'
GOGCCFLAGS='-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build3988505862=/tmp/go-build -gno-record-gcc-switches'
```

as you see ```GOOS='linux'`` line in the output we are going to use this to build executables, my os is linux i'm going to create executable for windows, linux or mac also.

run this command to build executable for windows from linux
``` bash
$ GOOS="windows" go build
```
you can update ```GOOS="windows" go build``` command as you're requirement
``` bash
$ GOOS="linux" go build
```
``` bash
$ GOOS="darwin" go build
```
or you can simply use ```go build``` to create executable for your os
```bash
go build
```

### Pointers
```go
package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int
	fmt.Println("value of pointer is:", ptr)

	// pointer -> *, reference -> &
	myNumber := 23
	// var ptr *int = &myNumber
	var ptr1 = &myNumber
	// var ptr *int = myNumber // error
	fmt.Println("Value of actual pointer is:", ptr1) // 0xc0000a4010
	fmt.Println("Value of actual pointer is:", *ptr1) // 23

	// updating variable value through pointer
	*ptr1 = *ptr1 + 3
	fmt.Println("New Value is:", *ptr1)

}
```

### Arrays in Go 
```go
package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Grapes"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList)) // 4

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Vegy list is:", vegList)
	fmt.Println("Vegy list is:", len(vegList))

}
```

### Slices
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	// CREATION
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// APPENDING
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// GETTING SLICES 
	// fruitList = append(fruitList[1:3])
	// fruitList = fruitList[1:3]
	fruitList = fruitList[0:3]
	fruitList = fruitList[0:]
	fruitList = fruitList[0:5]
	fmt.Println(fruitList)

	// CREATING SLICE
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 944
	highScores[2] = 454
	highScores[3] = 876

	// NOT WORK
	//	highScores[4] = 777 // error

	// WORK
	highScores = append(highScores, 555, 666, 321) // ok
	fmt.Println(highScores)
	fmt.Println(len(highScores))

	// SORTING
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	fmt.Println(highScores)

}
```

#### How to remove a value from slices based on index
```go
var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
fmt.Println(courses)
var index int = 2
courses = append(courses[:index], courses[index+1:]...)
fmt.Println(courses)
```
```output
[reactjs javascript swift python ruby]
[reactjs javascript python ruby]
```

### Maps
maps are used to store key value pairs in go lang.
code example:
```go
language := make(map[string]string)
// adding values
language["JS"] = "Javascript"
language["RB"] = "Ruby"
language["PY"] = "Python"

// printing entire map and a specific key / value
fmt.Println("List of all language: ", language)
fmt.Println("Js Shorts for: ", language["JS"])

// deleting/removing item from map
delete(language, "RB")
fmt.Println("List of all language: ", language)

// loops are interesting in golang
for key, value := range language {
	fmt.Printf("For key %v, value is %v\n", key, value)
}

// comma ok
for _, value := range language {
	fmt.Printf("For key v, value is %v\n", value)
}
```

### Intro to Struct
There is no inheritance in go lang; No super or parent ...
```go

func main() {
	fmt.Println("Structs in go lang")

	bugs := User{"Bugs Founder", "bugsfounder2021@gmail.com", true, 20}
	fmt.Println(bugs)
	fmt.Printf("Bugs details are: %+v\n", bugs)
	fmt.Printf("Name is %v and email is %v\n", bugs.Name, bugs.Email)
	
}

type User struct {
	Name   string 
	Email  string
	Status bool
	Age    int
}

```

```output
{Bugs Founder bugsfounder2021@gmail.com true 20}
Bugs details are: {Name:Bugs Founder Email:bugsfounder2021@gmail.com Status:true Age:20}
Name is Bugs Founder and email is bugsfounder2021@gmail.com
```

```%+v``` is very useful to print struct with full detail.

# if else 
To check conditions.
```go
loginCount := 23
var result string

// if-else if-else ladder 
if loginCount < 10 {
	result = "Regular user"
} else if loginCount > 10 {
	result = "Watch out"
} else {
	result = "Exactly 10 login count"
}

fmt.Println(result)

// if else
if 9%2 == 0 {
	fmt.Println("Number is even")
} else {
	fmt.Println("Number is odd")
}

// creating variable and checking it in if
if num := 3; num < 10 {
	fmt.Println("Num is less than 10")
} else {
	fmt.Println("Num is NOT Less than 10")
}

// reverse condition statement '!'
// if err!= nil{

// }
```
### Switch Case
```go
func main() {
	fmt.Println("Switch and case in golang")

	// rand.Seed(time.Now().UnixNano()) // rand.Seed is deprecated
	// we can use crypto 

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough // runs all cases below this case except last
	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough // runs all cases below this case
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}
}
```

### Loops
Basic Kind of loops or traditional for loop
```go
days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

fmt.Println(days)

// Basic / old
for d := 0; d < len(days); d++ {
	fmt.Println(days[d])
}
```
#### Get index in i and iterate
```go
for i := range days {
	// i is index
	fmt.Println(days[i])
}
```
#### Kind of for-each (for each loop)
```go
for index, day := range days {
	fmt.Printf("index is %v and value is %v\n", index, day)
}
```
#### kind of while loop
```go
rougueValue := 1
for rougueValue < 10 {
	fmt.Println("Value is : ", rougueValue)
	// ++rougueValue // error
	rougueValue++
}
```
#### break statement
```go
rougueValue := 1
for rougueValue < 10 {
	if rougueValue == 5 {
		rougueValue++
		break
	}
	fmt.Println("Value is : ", rougueValue)
	rougueValue++
}
```
#### continue statemetn
```go
rougueValue := 1
for rougueValue < 10 {
	if rougueValue == 5 {
		rougueValue++
		continue
	}
	fmt.Println("Value is : ", rougueValue)
	rougueValue++
}
```
#### goto statement
defining a label
```go
bf:
	fmt.Println("Jumping at bugsfounder.com")
```
Using goto statement to jump to the label
```go
rougueValue := 1
for rougueValue < 10 {
	if rougueValue == 2 {
		goto bf // jumping to the level
	}
	fmt.Println("Value is : ", rougueValue)
	rougueValue++
}
```

### Functions
Simple Functions
```go
func greeter() {
	fmt.Println("Namastey from golang")
}
```

#### Get any number of arguments or params (variadic functions)
```go
func proAdder(values ...int) int {
	total := 0
	for _, values := range values {
		total += values
	}

	return total
}
```

#### return multiple values
```go
func proAdderMul(values ...int) (int, string) {
	total := 0
	for _, values := range values {
		total += values
	}

	return total, "Hello Pro Adder"
}
```

#### Using functions
``` go
func main() {
	fmt.Println("Functions in Go")

	// greeter // reference
	greeter() // call

	// Note: cannot create function inside function

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult := proAdder(1, 2, 3, 5, 4, 5, 4, 3)
	fmt.Println("Pro result is: ", proResult)

	mulR, myMessage := proAdderMul(1, 4, 3, 23, 5)

	fmt.Println(mulR, myMessage)

}
```
Note: Type of parameters and return types are compulsory

### Methods 
Methods are similar to functions, the difference is methods are those functions which are written into a class but we don't have classes in go so the methods are those functions which are written into the structs.
```go
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int // private 'first letter small'
}
```
```go
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}
```
```go
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
```

```go
func main() {

	bugs := User{"Bugs Founder", "bugsfounder2021@gmail.com", true, 20}
	fmt.Println(bugs)

	bugs.GetStatus()
	bugs.NewMail() // test@go.dev
	fmt.Println(bugs.Email) // "bugsfounder2021@gmail.com"
}
```
Note: for update the original value we use pointers or references.

### Defer in go lang
All defer statements in a codebase are aligned at the end of the code 'main' function, then it prints the aligned values in LIFO Manner. ex: 0,1,2,3,4 (aligned using defer), prints: 4,3,2,1,0
```go
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
```
```go
func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")

	myDefer() 
}
```
deferred prints values in LIFO manner
alignment: world, One, Two, 0, 1, 2, 3, 4
```output
4
3
2
1
0
Two
One
World
```

### Files in GO 

#### ReadFile function to read contents of a file and print the contents of that file into the terminal
```go
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename) // deprecated
	// default data type of ReadFile func is Byte

	checkNilError(err)

	fmt.Println("Text Data inside the file is \n", string(databyte))
}
```
#### This logic is repeating to many times so it is good to wrap it into a function.
```go
if err != nil {
	panic(err)
}
```
Created a function ```checkNilError()```
#### Function to check err so we don't need to repeat the same logic again and again.
```go
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
```
#### Function to Write into a file
```go
func writeFile(filename string, content string) {
	// Creating a file
	file, err := os.Create(filename)

	// checking error
	checkNilError(err)

	// writing to the file
	length, err := io.WriteString(file, content)

	// checking error
	checkNilError(err)

	fmt.Println(length)

	// close file at the end of the program 'end of the main function'
	defer file.Close()
}
```

```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - BugsFounder.com" // content
	filename := "./mybffile.txt"                              // filename

	// writing to file using writeFile function 
	writeFile(filename, content)



	// reading the file using readFile function 
	readFile(filename) // reading file './mybffile.txt'

}
```

### Web Requests in GO
Creating a function which send a get request to url and returns string content
```go
func request(url string) (data string, err error) {
	res, err := http.Get(url) // Get Request

	// checking status code if it is ok return str data 
	if res.StatusCode == 200 {

		databyte, err := io.ReadAll(res.Body)

		checkNilError(err)

		return string(databyte), nil
	}

	// return error
	return "", err
}
```

#### Using request function
sending request to url using request function which we have created.
```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

// dummy apis: https://jsonplaceholder.typicode.com/
const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("Welcome to web request")

	res_content, err := request(url)

	checkNilError(err)

	fmt.Println(res_content)

}
```

### URLs in Go (url)
```go
package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.bugsfounder.com/learn?coursename=go&payment_id=asdfa4345-bugs"

func main() {
	fmt.Println("Welcome to handling urls")

	result, _ := url.Parse(myurl)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("RawPath: ", result.RawPath)
	fmt.Println("RawQuery: ", result.RawQuery)
	fmt.Println("User: ", result.User)

	qparams := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["payment_id"])

	for key, value := range qparams {
		fmt.Println(key, value)
	}

	// url must be always a reference '&url' not 'url'
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "bugsfounder.com",
		Path:    "/go",
		RawPath: "user=bugs",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}
```