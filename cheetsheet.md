
---

## 📘 Go CheetSheet

---

### 🔹 1. Packages and Modules

* Every Go file starts with a `package` declaration.
* `main` is the entry point for executable programs.
* Modules are initialized using:

```bash
go mod init module-name
```

```go
package main  // Entry point package
```

---

### 🔹 2. Importing Packages

Use `import` to include packages:

```go
import (
    "fmt"
    "strings"
    "booking-app/helper"  // Custom package
)
```

---

### 🔹 3. Package-Level Variables

Declared outside any function, accessible globally within the package.

```go
var conferenceName = "Go Conference"
const totalTickets = 50
var remainingTickets uint = 50
```

---

### 🔹 4. Exported vs Unexported Identifiers

* **Exported**: Starts with a capital letter (`Println`)
* **Unexported**: Starts with a lowercase letter

```go
// helper.go
func GreetUsers() {}      // Exported
func validateInput() {}   // Unexported
```

---

### 🔹 5. Data Types

* **Primitive**: `int`, `uint`, `float64`, `bool`, `string`
* **Derived**: `arrays`, `slices`, `structs`, `maps`

```go
var age int = 21
var name string = "GoLang"
```

---

### 🔹 6. Constants

```go
const appName = "BookingApp"
```

---

### 🔹 7. Arrays and Slices

#### Arrays

Fixed size.

```go
var arr [3]string
arr[0] = "Go"
```

#### Slices

Dynamic arrays.

```go
var names = []string{"Alice", "Bob"}
names = append(names, "Charlie")
```

---

### 🔹 8. Loops

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Looping through slices:

```go
for index, name := range names {
    fmt.Println(index, name)
}
```

---

### 🔹 9. Conditionals

```go
if x > 10 {
    fmt.Println("Greater")
} else if x == 10 {
    fmt.Println("Equal")
} else {
    fmt.Println("Less")
}
```

---

### 🔹 10. Functions

```go
func greetUser(name string) {
    fmt.Println("Hello", name)
}
```

Returning multiple values:

```go
func validate(name string) (bool, string) {
    return len(name) > 2, "Too short"
}
```

---

### 🔹 11. Structs

```go
type User struct {
    firstName string
    lastName  string
    email     string
    age       int
}
```

Creating and using:

```go
user := User{"John", "Doe", "john@example.com", 30}
fmt.Println(user.email)
```

---

### 🔹 12. Maps

```go
userMap := map[string]string{
    "name":  "Alice",
    "email": "alice@example.com",
}
```

---

### 🔹 13. Pointers

```go
a := 10
ptr := &a
fmt.Println(*ptr)  // Dereference
```

---

### 🔹 14. Defer

Runs after the function ends.

```go
defer fmt.Println("This runs last")
fmt.Println("Main logic")
```

---

### 🔹 15. Error Handling

```go
result, err := strconv.Atoi("123")
if err != nil {
    fmt.Println("Conversion error")
}
```

---

### 🔹 16. User Input (Console)

```go
var name string
fmt.Print("Enter name: ")
fmt.Scan(&name)
```

---

### 🔹 17. Strings Package

```go
strings.ToUpper("hello")
strings.Contains("email@example.com", "@")
```

---

### 🔹 18. Working with Time

```go
import "time"

now := time.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

---

### 🔹 19. Goroutines (Concurrency)

```go
go greet("User")  // Runs in a separate thread
```

---

### 🔹 20. Channels (Communication)

```go
ch := make(chan string)

go func() {
    ch <- "Hello"
}()

msg := <-ch
fmt.Println(msg)
```

---

### 🔹 21. File Structure Best Practices

```
booking-app/
├── go.mod
├── main.go
└── helper/
    └── helper.go
```

Use folders for logical grouping, and capitalize exported functions.

---

### 🔹 22. Go Build & Run

* Build: `go build`
* Run: `go run main.go`
* Test: `go test`

---

### 🔹 23. Go Formatting & Linting

* Format: `go fmt ./...`
* Vet (checks correctness): `go vet ./...`
* Lint (optional): Use `golangci-lint`

---

### 🔹 24. Unit Testing

```go
func Add(a, b int) int {
    return a + b
}
```

**Test File**: `add_test.go`

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
```

Run with:

```bash
go test
```

---

### 🔹 25. Dependency Management

* Initialize module: `go mod init`
* Add packages: `go get github.com/xyz/package`
* Tidy: `go mod tidy`

---

### 🔹 26. Interfaces (Polymorphism)

```go
type Speaker interface {
    Speak() string
}

type Human struct{}

func (h Human) Speak() string {
    return "Hello!"
}
```

---

### 🔹 27. Blank Identifier (`_`)

Ignore a returned value:

```go
_, err := someFunc()
```

---

### 🔹 28. Naming Conventions

* `CamelCase` for exported
* `camelCase` for internal
* File names should match package purpose

---

### 🔹 29. Go Tools

* `go run`: Compile and run
* `go build`: Compile
* `go test`: Run unit tests
* `go fmt`: Format code
* `go vet`: Analyze for bugs

