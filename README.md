# GO Unit Tester

Go Unit Tester is a small Golang library that provided a utility functions for unit testing.

## Usages

### Setting up:
Run the following command to install the module.

```sh
go get github.com/JohnFarmers/go-unit-tester
```

#### Example:

Let's say you have project structure like this(be sure to also create `unittester/unittester.go` file as shown below):

```
| go-example-project
|   > unittester
|       - unittester.go
|   > mathUtil
|       - mathUtil.go
|   main.go
```

If your project have a file `mathUtil.go` with the following methods:

```go
package mathUtil

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}
```

In `unittester.go` file you will have to create `inti()` function, this function will automatically be call once before the `main()` function. Be sure to import the module `"github.com/JohnFarmers/go-unit-tester"` first and then you can put your unit test code here.

```go
package unittester

import (
    test "github.com/JohnFarmers/go-unit-tester"
    mth "go-example-project/mathUtil"
)

func init() {
	// You can perform unit test like this.
	test.UnitTest(mth.Add, []interface{}{5}, []interface{}{2, 3}, false, true)
	test.UnitTest(mth.Subtract, []interface{}{2}, []interface{}{10, 8}, false, true)
	test.UnitTest(mth.Multiply, []interface{}{10}, []interface{}{5, 2}, false, true)
}
```

Lastly, in order for the `init()` function to be call, all you have to do is go to your `main.go` and import the package that the `init()` is in:

```go
package main

import (
	_ "go-example-project/unittester"
)

func main() {
    //Your code here.
}
```

Note: the `_` before the package path is needed. If you don't include the `_`, the imported package will be remove when you save the file.

Once everything is done, you can use the command to run your project and the unit test code should be running before your `main()`.

```sh
go run .
```

### Method definitions:

#### UnitTest

`UnitTest(function interface{}, expected []interface{}, params []interface{}, checkOutputTypeOnly bool, detailedPassLog bool) bool`

#### Description
Perform a unit test on the given function by checking if the real outputs is the same as the expected outputs or not and print out error if any unexpected outputs/behavior occured.

#### Parameters

| Name | Type | Description |      
| ------------- |------| ------------- |
| `function` | `interface{}` | The function to preform a unit test on. |
| `expected` | `[]interface{}` | The expected outputs of the function. |
| `params` | `[]interface{}` | The inputs of the function. |
| `checkOutputTypeOnly` | `bool` | Determine whether or not to check the type of the output only and ignore it's values. If `true`, when checking the output, it will only check if the type of the outputs match the expected outputs. If `false`, it will check both type and value. |
| `detailedPassLog` | `bool` | Determine whether or not to print the full outputs of the tested function if the test is pass. |

#### Return types: `bool`

#### Example:

```go
import (
    test "github.com/JohnFarmers/go-unit-tester"
)

// If you want to unit test this function.
func Add(a int, b int) int {
    return a + b
}

// You can do so like this.
test.UnitTest(Add, []interface{}{5}, []interface{}{2, 3}, false, true)
```

The function call above perform a unit test on function `Add(a int, b int)` with `5` as an expected output and also with `2` and `3` as an inputs.

#### UnitTestWithMultipleOutputCase

`UnitTestWithMultipleOutputCase(function interface{}, expectedOutputs [][]interface{}, params []interface{}, checkOutputTypeOnly bool, detailedPassLog bool) bool`

#### Description
Perform a unit test on the given function by checking if the real outputs is the same as the expected outputs or not and print out error if any unexpected outputs/behavior occured.

This method is for checking a function that can have multiple output cases.

#### Parameters

| Name | Type | Description |      
| ------------- |------| ------------- |
| `function` | `interface{}` | The function to preform a unit test on. |
| `expected` | `[][]interface{}` | The expected outputs of the function in each cases. |
| `params` | `[]interface{}` | The inputs of the function. |
| `checkOutputTypeOnly` | `bool` | Determine whether or not to check the type of the output only and ignore it's values. If `true`, when checking the output, it will only check if the type of the outputs match the expected outputs. If `false`, it will check both type and value. |
| `detailedPassLog` | `bool` | Determine whether or not to print the full outputs of the tested function if the test is pass. |

#### Return types: `bool`

#### Example:

```go
import (
	"math/rand/v2"
	test "github.com/JohnFarmers/go-unit-tester"
)

// If you want to unit test this function.
func RandNum() int {
    return rand.IntN(3)
}

// You can do so like this.
test.UnitTestWithMultipleOutputCase(
    RandNum,
    [][]interface{}{
        []interface{}{0}, 
        []interface{}{1}, 
        []interface{}{2},
    },
    []interface{}{},
    false,
    true,
)
```

Since the function `RandNum()` above can output either 0, 1, or 2. We can have multiple expected outputs, so that we can check for all of them.
