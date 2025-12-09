# Go Learning Journey 🚀

Welcome to my Go programming learning repository! This README documents my progress as I learn the fundamentals of Go, up through loops and control flow.

## What I've Learned So Far

### Basic Syntax
- Package declaration and imports
- Variable declarations using `var` and `:=`
- Basic data types: `int`, `float64`, `string`, `bool`
- Functions and return values
- `fmt.Println()` for output

### Control Structures
- `if` statements and `if-else` chains
- Comparison operators: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Logical operators: `&&` (AND), `||` (OR), `!` (NOT)

### Loops
- `for` loop - Go's only looping construct
- Traditional three-component loop: `for i := 0; i < 10; i++ {}`
- While-style loop: `for condition {}`
- Infinite loop: `for {}`
- `break` and `continue` statements
- Iterating over ranges with `range`

## Sample Code

### Simple For Loop
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### While-style Loop
```go
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}
```

### Range Loop
```go
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

## Resources I'm Using
- Official Go documentation: https://go.dev/doc/
- Go by Example: https://gobyexample.com/
- A Tour of Go: https://go.dev/tour/

## Next Steps
- Arrays and slices
- Maps
- Structs
- Pointers
- Methods and interfaces

## Notes
- Go uses `gofmt` for code formatting
- Run programs with `go run filename.go`
- Build executables with `go build`

---

**Last Updated:** December 2025
