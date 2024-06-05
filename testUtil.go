package unittester

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

func UnitTest(function interface{}, expected []interface{}, params []interface{}, checkOutputTypeOnly bool) bool {
	fnType := reflect.TypeOf(function)
	if fnType.Kind() != reflect.Func {
		fmt.Println("\033[31mFAIL:", function, "is not a function.\033[0m")
		return false
	}

	args := []reflect.Value{}
	for _, param := range params {
		args = append(args, reflect.ValueOf(param))
	}

	functionName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	inputLength := fnType.NumIn()

	if len(params) != inputLength {
		fmt.Println("\033[31mFAIL: The amount of given inputs is", len(params), "which doesn't match the amount of input of function", functionName, "which is", inputLength, "\033[0m")
		return false
	}

	for i := 0; i < inputLength; i++ {
		realInputType := fnType.In(i)
		inputType := reflect.TypeOf(params[i])
		if realInputType != inputType {
			fmt.Println("\033[31mFAIL:", functionName, "input type index of", i, "is", realInputType, "but the input that has been given is", inputType, "\033[0m")
			return false
		}
	}

	f := reflect.ValueOf(function)
	results := f.Call(args)

	outputLength := len(results)

	if len(expected) != outputLength {
		fmt.Println("\033[31mFAIL: The amount of outputs of function", functionName, "is", outputLength, "which doesn't match the amount of expected output which is", len(expected), "\033[0m")
		return false
	}

	if outputLength >= 1 {
		lastOutput := results[outputLength-1]
		if lastOutput.Interface() != nil && reflect.TypeOf(lastOutput.Interface()).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			fmt.Println("\033[31mFAIL:", functionName, "returned an error:", lastOutput.Interface())
			return false
		}
	}

	for i, value := range results {
		if reflect.TypeOf(value.Interface()) != reflect.TypeOf(expected[i]) {
			fmt.Println("\033[31mFAIL:", functionName, "output index of", i, "return type is", reflect.TypeOf(value.Interface()), "but expected type of", reflect.TypeOf(expected[i]), "\033[0m")
			return false
		}
		if checkOutputTypeOnly {
			continue
		}
		if value.Interface() != expected[i] {
			fmt.Println("\033[31mFAIL:", functionName, "output index of", i, "returned", value.Interface(), "but expected", expected[i], "\033[0m")
			return false
		}
	}

	fmt.Println("\033[32mPASS: " + functionName + " function outputs " + formatValuesAsStr(results) + " with " + formatValuesAsStr(args) + " as an arguments and run successfully.\033[0m")
	return true
}

func UnitTestWithMultipleOutputCase(function interface{}, expectedOutputs [][]interface{}, params []interface{}, checkOutputTypeOnly bool) bool {
	fnType := reflect.TypeOf(function)
	if fnType.Kind() != reflect.Func {
		fmt.Println("\033[31mFAIL:", function, "is not a function.\033[0m")
		return false
	}

	args := []reflect.Value{}
	for _, param := range params {
		args = append(args, reflect.ValueOf(param))
	}

	functionName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	inputLength := fnType.NumIn()

	if len(params) != inputLength {
		fmt.Println("\033[31mFAIL: The amount of given inputs is", len(params), "which doesn't match the amount of input of function", functionName, "which is", inputLength, "\033[0m")
		return false
	}

	for i := 0; i < inputLength; i++ {
		realInputType := fnType.In(i)
		inputType := reflect.TypeOf(params[i])
		if realInputType != inputType {
			fmt.Println("\033[31mFAIL:", functionName, "input type index of", i, "is", realInputType, "but the input that has been given is", inputType, "\033[0m")
			return false
		}
	}

	f := reflect.ValueOf(function)
	results := f.Call(args)

	numCase := len(expectedOutputs)
	outputLength := len(results)
	isSuccessful := false
	errorMessages := [][]interface{}{}

out:
	for i := 0; i < numCase; i++ {
		expected := expectedOutputs[i]
		caseN := strconv.Itoa(i+1) + ":"

		if len(expected) != outputLength {
			errorMessages = append(errorMessages, []interface{}{"\t\033[31mOutput case", caseN, "The amount of outputs of function", functionName, "is", outputLength, "which doesn't match the amount of expected output which is", len(expected), "\033[0m"})
			continue out
		}

		if outputLength >= 1 {
			lastOutput := results[outputLength-1]
			if lastOutput.Interface() != nil && reflect.TypeOf(lastOutput.Interface()).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
				errorMessages = append(errorMessages, []interface{}{"\t\033[31mOutput case", caseN, functionName, "returned an error:", lastOutput.Interface(), "\033[0m"})
				continue out
			}
		}

		for i, value := range results {
			if reflect.TypeOf(value.Interface()) != reflect.TypeOf(expected[i]) {
				fmt.Println("\t\033[31mOutput case", functionName, "output index of", i, "return type is", reflect.TypeOf(value.Interface()), "but expected type of", reflect.TypeOf(expected[i]), "\033[0m")
				return false
			}
			if checkOutputTypeOnly {
				continue
			}
			if value.Interface() != expected[i] {
				errorMessages = append(errorMessages, []interface{}{"\t\033[31mOutput case", caseN, functionName, "output index of", i, "returned", value.Interface(), "but expected", expected[i], "\033[0m"})
				continue out
			}
		}
		isSuccessful = true
	}

	if isSuccessful {
		fmt.Println("\033[32mPASS: " + functionName + " function outputs " + formatValuesAsStr(results) + " with " + formatValuesAsStr(args) + " as an arguments and run successfully.\033[0m")
	} else {
		fmt.Println("\033[31mFAIL:", functionName, "error logs: \n[\033[0m")
		printFn := reflect.ValueOf(fmt.Println)
		for _, msg := range errorMessages {
			reflectValues := []reflect.Value{}
			for _, a := range msg {
				reflectValues = append(reflectValues, reflect.ValueOf(a))
			}
			printFn.Call(reflectValues)
		}
		fmt.Println("\033[31m]\033[0m")
	}

	return isSuccessful
}

func formatValuesAsStr(values []reflect.Value) string {
	if len(values) <= 0 {
		return "nothing"
	}

	var sb strings.Builder
	maxIndex := len(values) - 1

	for i, value := range values {
		jsonBytes, _ := json.Marshal(value.Interface())
		str := string(jsonBytes)
		sb.WriteString(str)
		if i < maxIndex && maxIndex != 0 {
			sb.WriteString(", ")
		}
	}

	return sb.String()
}
