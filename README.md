```go
import "fmt"

func main(){
	fmt.Print("Hello World")
}
```

```go
go run app.go
```

- Main function is the entry point for whole GO application
- Print is a method that comes from the fmt package
- Strings are only allowed in double quotes or backticks only
- Every Go file needs a package imported that will be same as the file name

- Idea for packages is to organize code
- package can be made up of multiple files

## Explicit type declaration

- Go can by itself get the type of values being stored in the variable, but we can also declare a type so that we can say what value can be stored there and with that we can remove any extra type conversion

```go
	var investmentAmount float64 = 1000
```

- explicit type can be added after the variable name

## Variable Declarations

- variables can be declared via “var” keyword
- and also with special assignment operator “:=” so with that we can remove the “var” keyword
- this assignment operator is useful when we dont have to change the type of variable that is assigned by GO
- because with “:=” we can not declare the type of variable

```go
expectedReturn float64 := 5.5
```

above is invalid

```go
expectedReturn := 5.5
```

above is valid

- we can declare multiple variables at the same time with same type

```go
	var investmentAmount, years float64 = 1000, 10
```

- if we dont declare the type of variable on that line then we can also declare multiple variables with different types

```go
	var investmentAmount, years  = 1000, "10"
```

## User Input

- Getting values from users through cli is pretty similar to how it is done in C language
- we can use the Scan() method of the fmt package to do that.

```go
fmt.Scan(&investmentAmount)
```

- here the investmentAmount is already declared and it has a value in it so to override that we are using the “&” symbol, which is a pointer to the actual variable and here will override the value from user input
- If user dont give any value to the variable, it will be initialized with the default value of the type, for float it will be 0.0

## Declaring Variables without initial value

- We can declare a variable without initializing any value to that variable, but for that, we will need to explicitly declare the type of the variable, which required as GO Lang is a statically typed language
- This will be declared using the var keyword and then the variable name and then the type of variable

```go
var years float64
var expectedReturn float64
```

## All Methods to declare variables in GO Lang

```go
const inflationRate = 2.5
var investmentAmount float64 = 1000
var years float64
expectedReturn := 5.5
```

## We can format the Output Strings with Printf() Method

```go
fmt.Printf("Future Value : %v\nFuture Value (Adjusted for Inflation): %v", futureValue, futureRealValue)
```
