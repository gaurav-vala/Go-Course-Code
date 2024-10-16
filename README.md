- We can split code in multiple files, then split those files in multiple packages
- then we can use those packages in out projects
- It seems similar to making components, splitting code into multiple files to handle code in better way, that way we can create our own packages too

### Creating Packages

- we can just create a new `.go` file
- every go file must be part of a package
- whenever we declare a file in the same package as other file then we can just directly use the functions in any file of the same package
- e.g., we create a file with a function, to call that function in a different file we can have both the file in same package, and then directly call the function
- Imports are not passed directly through packages
- so we have to import other packages that are being used in the function of the new go file
- If you create multiple files into a package, then you can use that package into different projects, that way you can have a package that has code that you use again and again then that code goes into a package and that package can be used in another project for the same purpose

### Custom Packages

- To create a custom package we will have to create a new folder for that package
- the folder name should be the same as the package
- the file name can be different, does not need to be same as the package

```go
package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
```

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/a272ab31-7a10-4c5c-9805-a365981bae11/7b280edd-5f4e-45d3-9959-6eae56c62e71/image.png)

- Here I have a package named fileops for which i have folder names fileops which contains a file
- that file can be named anything, on top of the file i declared the package name which my file is part of

### Using packages

- to use this custom package, you have to import it, the process is same as other inbuilt packages
- but in this case you have to point out to the path same as the `go.mod` file

```go
package main

import (
	"fmt"
	"example.com/bank/fileops"
)
```

- here the `example.com/bank` comes from the `go.mod` file which was created when creating the project

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/a272ab31-7a10-4c5c-9805-a365981bae11/8f4f214d-ed69-4a49-abbc-6f3c660af85c/image.png)

### Exporting Identifies - Variables, function, etc…

- if you are using some function from a different package, and you have created that package, you will have to export those functions
- To export something from a package, that as to be defined in a way so that the name is starting from an uppercase

```go
func WriteFloatToFile(value float64, fileName string) {}
func GetFLoatFromFile(fileName string) (float64, error) {}
```

- this works as same as `export` keyword in JavaScript, where we define a function or variable with export keyword and explicitly, define that, that variable or functions needs to be exported and can be used in other files
- So the naming of the function or variable can define if something is exported or not

```go
package main

import (
	"fmt"

	"example.com/bank/fileops"
)
fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
```
