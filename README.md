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
