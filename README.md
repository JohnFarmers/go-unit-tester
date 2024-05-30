# GO Unit Tester

Go Unit Tester is a small Golang library that provided a utility functions for unit testing.

## Usages

### Setting up:
Drag and drop this folder into your project directory and you can import this package and call it's function.

In some case you might want to run the test before running your `main()`, You can do so by using the template in `unittest.go` file.

#### Example:

In `unittester.go` file there is a `inti()` function, this function will automatically be call once before the `main()` function. You can put your unit test here.

For example, let's say you have project structure like this(assuming that you already imported this package):

```
| go-example-project
|   > go-unittester
|       - testUtil.go
|       - unittester.go
|       - README.md
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

In `unittester.go` you can test them like this:

```go
package unittester

import (
    mth "go-example-project/mathUtil"
)

func init() {
	// You can perform unit test by calling UnitTest function like this.
	UnitTest(mth.Add, []interface{}{5}, []interface{}{2, 3}, false)
    UnitTest(mth.Subtract, []interface{}{2}, []interface{}{10, 8}, false)
    UnitTest(mth.Multiply, []interface{}{10}, []interface{}{5, 2}, false)
}
```

Lastly, in order for the `init()` function to be call, all you have to do is go to your `main.go` and import the package that the `init()` is in:

```go
package main

import (
	_ "go-example-project/go-unit-tester"
)

func main() {
    //Your code here.
}
```

Note: the `_` before the package path is needed. If you don't include the `_`, the imported package will be remove when you save the file.

Once everything is done, you can use te command to run your project and the unit test code should be running before your `main()`.

```sh
go run .
```

### Method definitions:

#### UnitTest

`UnitTest(function interface{}, expected []interface{}, params []interface{}, checkOutputTypeOnly bool) bool`

#### Description
Perform a unit test on the given function by checking if the real outputs is the same as the expected outputs or not and print out error if any unexpected outputs/behavior occured.

#### Parameters

| Name | Type | Description |      
| ------------- |------| ------------- |
| `function` | `interface{}` | The function to preform a unit test on. |
| `expected` | `[]interface{}` | The expected outputs of the function. |
| `params` | `[]interface{}` | The inputs of the function. |
| `checkOutputTypeOnly` | `bool` | Determine whether or not to check the type of the output only and ignore it's values. If `true`, when checking the output, it will only check if the type of the outputs match the expected outputs. If `false`, it will check both type and value. |

#### Return types: `bool`

#### Example:

```go
// If you want to unit test this function.
func Add(a int, b int) int {
    return a + b
}

// You can do so like this.
UnitTest(Add, []interface{}{5}, []interface{}{2, 3}, false)
```

The function call above perform a unit test on function `Add(a int, b int)` with `5` as an expected output and also with `2` and `3` as an inputs.

#### UnitTestWithMultipleOutputCase

`UnitTestWithMultipleOutputCase(function interface{}, expectedOutputs [][]interface{}, params []interface{}, checkOutputTypeOnly bool) bool`

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

#### Return types: `bool`

#### Example:

```go
import (
	"math/rand/v2"
)

// If you want to unit test this function.
func RandNum() int {
    return rand.IntN(3)
}

// You can do so like this.
UnitTestWithMultipleOutputCase(
    RandNum,
    [][]interface{}{
        []interface{}{0}, 
        []interface{}{1}, 
        []interface{}{2},
    },
    []interface{}{},
    false,
)
```

Since the function `RandNum()` above can output either 0, 1, or 2. We can have multiple expected outputs, so that we can check for all of it.