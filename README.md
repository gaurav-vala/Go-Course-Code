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

- there are different option to add to the output string that can format it
- here we have `%v` and `\n` that helped to add the variable value and also with `\n` everything after it will be on next line
- There are many of these “verbs that we can add to format, check the official docs
- We can use `%.0f` to round the floating numbers
- the number before f refers to the number we wanna show after the .
- so if we say `%.2f` that will print 2 number after .
-

```go
fmt.Printf("Future Value : %v\nFuture Value (Adjusted for Inflation): %v", futureValue, futureRealValue)
```

## Storing Formatted strings into a variable

- using the method `Sprintf()` we can store any formatted string into a variable and then use that variable instead of whole string

```go
formattedFV := fmt.Sprintf("Future Value : %.0f\n", futureValue)
formattedFRV := fmt.Sprintf("Future Value (Adjusted for Inflation): %0.f\n", futureRealValue)
```

- after that we can use the Print() method to print these strings which will just print the string without any formatting

```go
fmt.Print(formattedFV, formattedFRV)
```

## Multi line Strings

- We can use backticks ```` instead of double quotes to create multiline formatted string, this way

```go
fmt.Printf(`Future Value : %v
Future Value (Adjusted for Inflation): %v`, futureValue, futureRealValue)
```

## Defining Functions

- all the user defined Functions are defined below the `main` function
- a function in Go can be defined using `func` keyword
- We can add parameters to the function, when adding parameters we have to define the type of the parameter

```go
func outputText(text1 string, text2 string){
	fmt.Print()
}
```

- if the parameters are of same type then we can define the parameter and have the type only once, then we have to add the parameter name and at the end add the type of the parameter

```go
func outputText(text1, text2 string) {
	fmt.Print()
}
```

- After that we can use it like any other function in any language

```go
outputText("Investment Amount: ")
fmt.Scan(&investmentAmount)

func outputText(text1 string) {
	fmt.Print(text1)
}
```

- in Go lang, we can return multiple values from a single function
- In Go, variables are block scoped, which means any variables defined inside the function or block are scoped to that function and we can use them outside the function
- We can define the variables outside the main function or any other function to have them as global scope and then we can use them inside any function in the file.
- We can only define const and var declaration as global declaration, we cannot do this to any other declaration
- That means `:=` syntax cannot be used in the global scope

```go
const inflationRate = 2.5

func main() {
	var investmentAmount float64 = 1000
	var years float64
	var expectedReturn = 5.5
}
```

- when returning any value, we have to define the return type of the function
- If we are returning 2 values we have to wrap it into rounded parenthesis, and then define the type of specific return to its type
- with this we can define specific types to any return value `func demo ()(float64, int){}`

```go
func calculateFutureValue(investmentAmount, expectedReturn, years float64) (float64, float64) {
	fv := (investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
```

- when we want to store the return values in the variable, we can have 2 variable by comma separated variable names

```go
futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturn, years)
```
