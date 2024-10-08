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

### Explicit type declaration

- Go can by itself get the type of values being stored in the variable, but we can also declare a type so that we can say what value can be stored there and with that we can remove any extra type conversion

```go
	var investmentAmount float64 = 1000
```

- explicit type can be added after the variable name

### Variable Declarations

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

### User Input

- Getting values from users through cli is pretty similar to how it is done in C language
- we can use the Scan() method of the fmt package to do that.

```go
fmt.Scan(&investmentAmount)
```

- here the investmentAmount is already declared and it has a value in it so to override that we are using the “&” symbol, which is a pointer to the actual variable and here will override the value from user input
- If user dont give any value to the variable, it will be initialized with the default value of the type, for float it will be 0.0

### Declaring Variables without initial value

- We can declare a variable without initializing any value to that variable, but for that, we will need to explicitly declare the type of the variable, which required as GO Lang is a statically typed language
- This will be declared using the var keyword and then the variable name and then the type of variable

```go
var years float64
var expectedReturn float64
```

### All Methods to declare variables in GO Lang

```go
const inflationRate = 2.5
var investmentAmount float64 = 1000
var years float64
expectedReturn := 5.5
```

### We can format the Output Strings with Printf() Method

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

### Storing Formatted strings into a variable

- using the method `Sprintf()` we can store any formatted string into a variable and then use that variable instead of whole string

```go
formattedFV := fmt.Sprintf("Future Value : %.0f\n", futureValue)
formattedFRV := fmt.Sprintf("Future Value (Adjusted for Inflation): %0.f\n", futureRealValue)
```

- after that we can use the Print() method to print these strings which will just print the string without any formatting

```go
fmt.Print(formattedFV, formattedFRV)
```

### Multi line Strings

- We can use backticks ```` instead of double quotes to create multiline formatted string, this way

```go
fmt.Printf(`Future Value : %v Future Value (Adjusted for Inflation): %v`, futureValue, futureRealValue)
```

### Defining Functions

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

### Different way to declare return Variables

- when you declare the type of variables on function, there you can also add the variables name, and that way you dont have to declare those return variables
- also then you dont have to add those variable names after the return statement, Go will understand which variables to return

```go
func calculateFutureValue(investmentAmount, expectedReturn, years float64) (fv float64, rfv float64) {
	fv = (investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}
```

- starting new project with command

```go
go mod init example.com/bank
```

- Booleans for checks

```go
wantsCheckBalance := choice == 1
```

- if condition is similar but only round parenthesis is not used, directly the condition is used

```go
    if choice == 1 || choice == 2 {
    }
```

- All the conditional `&&` and `||` are also the same
- to check the condition or compare we can use the double equals operator `==`
- `else if` is also the same as other language

```go

	if choice == 1 {
		fmt.Println("Your Balance is", accountBalance)
	} else if choice == 2 {
  }
```

- Increment operator works the same

```go
accountBalance += depositAmount
//which equals to accountBalance = accountBalance + depositAmount
```

### Loops

- In Go lang there is only `for loop` , other than this there are no loops
- So, same as conditionals, loops does not have round brackets to wrap around the condition

```go
	for i := 0; i < 200; i++ {}
```

- they are directly declared with the `:=` syntax for the variable declaration

### Infinite for Loop

- if we define a for loop without any condition then that loop can work as an infinite loop

```go
	for {
		fmt.Println("===============")
		fmt.Println("What do you want to do?")
	}
```

- one way to get out of the infinite loop in this case would be to use `return` statement but with that whole program will be stopped and anything added after that will not be executed

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/a272ab31-7a10-4c5c-9805-a365981bae11/b976a23f-0349-4141-b0e3-cc6570c76b51/image.png)

- So to get out of the loops we can use the `break` statement, which will stop the loop and start executing next line of the loop
- We can use `continue` statement to break out of any particular condition and get back to the main loops, this helps when working with conditions inside a loop, so that way we can stop any particular condition and start the loops again
- `continue` statement will stop the current iteration of the loop and start again that loop

### Switch statement

- It is same as the other languages, nothing specific to GO in this

```go
switch choice {
		case 1:
			fmt.Println("Your Balance is", accountBalance)
		case 2:
			fmt.Print("Your Deposit: ")
		case 3:
			fmt.Print("Withdrawal Amount:")
		default:
			fmt.Print("Goodbye!")
		}
```

### Writing in Files

- the `os package` provides a function called `WriteFile` that takes in the name of the package, the data to be stored in the file and file mode, which will be the file modifying permissions.

```go
func writeBalanceFiles(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
```

### Reading from a file

- sometimes when we are getting 2 return values from a function, we can then use a special variable name `_` that means we want get the value but we dont want to use it.

```go
data, _ := os.ReadFile("balance.txt")
```

- here the data will come with a `byte` type, so to handle that, we can only convert that data into a string
- we need the data string in float number, so for that we can use the `strconv` package, that gives us different methods to handle string operations
- we can use `ParseFloat` method, which will convert the string into floating number, with that we will need to provide the `string value` and `bitSize` which will be 32 or 64

```go
func readFile() float64 {
	data, _ := os.ReadFile("balance.txt")
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}
```

- byte values cannot be directly converted into any type other then string and string value directly cannot be converted into any other type, that will require the use of `strconv` package

### Error handling

- In GO, there is a special type as `error` that can be used to give out custom errors which comes from the built in `errors package`
- In GO, we can use `nil` to check the null value
- almost all packages provide an error as return value with a main value, we use the error to to check if there is `nil` value
- So, if the error is not nil that means there is an error and in that case we can give out some error response

```go
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

```

```go
	var accountBalance, err = readFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("===============")
	}
```
