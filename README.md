# Learning Go

## Go CLI

- `go build`

  compile sorce code file(s)

- `go run`

  compile and execute

- `go fmt`

  format all the code in dir

- `go install`

  compile and installs a package

- `go get`

  download package

- `go test`

  run tests

## Go packages

Types of packages:

- executable (main)

  generates a file we can run must have a main function

- reusable

  place to store reusable logic

## Variables

```go
// explicit type
var card string = "ace of spades"
// inferred type
var card = "ace of spades"
// another style
card := "ace of spades"
```

## Arrays and Slices

**Arrays:** fixed lenght

**Slices:** dynamic array

## for range

```go
// slice
cards := []string{newCard()}
// append slice (a new array assigned to same variable, not extened)
cards = append(cards, "six of spades")

for i, card := range cards {
  fmt.Println(i, card)
}

// unused variables replaced with _
for _, card := range cards {
  fmt.Println(card)
}

for i, _ := range cards {
  fmt.Println(i)
}
```

## OO vs Go approach

Go does not have classes.

Go has custom types and methods can be defined on them. methods are functions with recievers.

### methods / reciever functions

```go
type deck []string

func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}
```

any variable with type `deck` gets acces to `print`, like in an instance would have access to instance methods.

two types of recievers:

- pointer reciever
- value reciever

pointer recivers can modify the value to which the reciever points.

## Value Types and Reference types

- **Value Types**

  need to use pointers to change inside function

- **Reference Types**

  dont need to wory about pointers.

| Value Types | Reference Types |
| :---------: | :-------------: |
|     int     |      slice      |
|    float    |       map       |
|   string    |     channel     |
|    bool     |     pointer     |
|   struct    |    function     |

## Go Tests

to make test files make file ending in `_test.go`

to test function add a function with `Test` then function name in test file. test functions take argument `t *testing.T` which is the test runner

```go
// deck.go
func newDeck() {
  // code
}

// deck_test.go
func TestNewDeck(t *testing.T){
  // test code
}
```

## Structs

kind of like a javascript object

### definition

```go
type person struct {
  firstName string
  lastName string
}
```

### declaring variables of type struct

```go
// rely on order of definition to pass values
lorem := person{"lorem", "ipsum"}

// explicitly state properties
lorem := person{firstName: "lorem", lastName: "ipsum"}

// declare variable without value, will use zero value (empty string)
var lorem person

lorem.firstName = "lorem"
lorem.lastName = "ipsum"
```

## Maps

keys all the same type and values all the same type.

```go
// one way to declare maps
colors := map[string]string{
  "red":   "#ff0000",
  "green": "#00ff00",
}

// without initializing a value
var colors map[string]string

// make empty map using make
colors := make(map[string]string)

// add new key value pai to map
colors["white"] = "#ffffff"

// delete key value pair
delete(colors, "white")
```

## Maps vs Structs

| Map                                                            |                                                Struct |
| :------------------------------------------------------------- | ----------------------------------------------------: |
| used to represent a collection of related properties           | used to represent a "thing" with different properties |
| all keys must of same type and all values must be of same type |                      values can be of different types |
| dont need to know all the keys at compile time                 |               need to know all feilds at compile time |
| keys are indexed - can iterate over them                       |                            keys dont support indexing |
| reference type                                                 |                                            value type |

## Interfaces

Interfaces applied implicitly to types that satisfy the interface, no need for `implements` keyword. Interfaces define method signatures. Interfaces can wrap other interfaces.

```go
// interface defintion
type bot interface {
  getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}


// implementin this function means that nglishBot satisfies bot interface and is of type bot
func (englishBot) getGreeting() string {
  return "hello"
}

func (spanishBot) getGreeting() string {
  return "holla"
}

// can now have a shared funtion between types that satisfy interface type
func printGreeting(b bot) {
  fmt.Println(b.getGreeting())
}
```
