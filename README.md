## GO CLI

- `go build`: compiles a bung of go source code files
- `go run`: compiles and executes one or two files
- `go fmt`: formats all the code in each file in the current directory
- `go install`: compiles and installs a package
- `go get`: downloads the raw source code of someone else's package
- `go test`: runs any tests associated with the current project

## Package

`package main` defines a package that can be compiled and the executed. Must have a func called 'main'

`package calculator` and `package uploader` defines a package that can be used as a dependency (helper code)

## import "fmt"

Give my package all the functionally inside "fmt"

## Basic Go Types

```go
bool: true - false
string: "Hi"
int: 0 - 999
float64: 10.0001
```

## Declare variables

```go
var card string = "Ace of Spades"
card := "Ace of Spades"
```

## Slices

```go
cards := []string{"Ace of Diamonds", "Other string"}
```

### Iterate

```go
	for i, card := range cards {
		fmt.Println(i, card)
	}
```

## Types

```go
type deck []string

cards := deck{"Ace of Diamonds", newCard()}
```

### Byte Slice 

```go
"Hi there!" -> String

[72 105 32 116 104 101 114 101 33] -> Byte slice

data []byte

asciitable.com
```

### Type Convertion

```go
[]byte("Hi there!")
```

Example: 

```go
greeting :=  "Hi!"

fm.Println([]byte(greeting))
```

## Receiver function

```go
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

Any variable of type `deck` now gets access to the `print` method.
For convention, we call the iterable with the first char of the name of the type. 


