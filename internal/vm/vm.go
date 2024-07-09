package vm

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/i5heu/TestDreadnought/extensions"
	"github.com/i5heu/TestDreadnought/internal/client"
	"github.com/i5heu/TestDreadnought/internal/config"
	"github.com/robertkrimen/otto"
)

func InitializeVM(globalScript, testCaseParentFolder, configDir string) (*otto.Otto, error) {
	vm := otto.New()

	// Will be used to print console.log messages in color
	SetupConsoleLog(vm)

	// Load and execute the global configuration file
	if err := config.LoadScript(vm, globalScript); err != nil {
		return nil, err
	}

	// Set up the clients for the VM - so that the VM can make HTTP requests
	SetUpClients(vm, testCaseParentFolder)

	// Set up the test function
	SetUpTestFunction(vm, testCaseParentFolder, configDir)

	// Set up extensions
	extensions.SetUpExtensions(vm, testCaseParentFolder, configDir)

	return vm, nil
}

func SetUpClients(vm *otto.Otto, testCaseParentFolder string) {
	c := color.New(color.FgCyan)
	cr := color.New(color.FgRed)

	// Define the post function
	vm.Set("Post", func(call otto.FunctionCall) otto.Value {
		url, _ := call.Argument(0).ToString()
		data, _ := call.Argument(1).Export()

		// Get the headers
		headers, err := config.GetHeaders(vm)
		if err != nil {
			cr.Println("	", err)
			value, _ := otto.ToValue(fmt.Sprintf("Error: %s", err.Error()))
			return value
		}

		// Get the base URL
		baseUrl, err := config.GetBaseUrl(vm)
		if err != nil {
			cr.Println("	", err)
			value, _ := otto.ToValue(fmt.Sprintf("Error: %s", err.Error()))
			return value
		}

		resp, body, err := client.PostRequest(baseUrl, url, headers, data)
		if err != nil {
			c.Println("	POST request failed", err)
			panic(vm.MakeCustomError("Error To Js", fmt.Sprintf("%s", err.Error())))
		}

		c.Println("	POST request successful")

		// Create an otto.Object and set properties
		obj, _ := vm.Object(`({})`)
		_ = obj.Set("response", resp)
		_ = obj.Set("body", body)

		return obj.Value()
	})

}

func SetUpTestFunction(vm *otto.Otto, testCaseParentFolder, configDir string) {
	// c := color.New(color.FgHiCyan)

	// Define the resultIsLikeFile function
	vm.Set("ResultIsLikeFile", func(call otto.FunctionCall) otto.Value {
		response, _ := call.Argument(0).ToString()
		filePathTestCaseRel, _ := call.Argument(1).ToString()
		filePath := filepath.Join(testCaseParentFolder, filePathTestCaseRel)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			msg := fmt.Sprintf("Expected result file does not exist: %s", filePath)
			color.Red("	" + msg)
			value, _ := otto.ToValue(msg)
			return value
		}

		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			msg := fmt.Sprintf("Error reading file %s: %s", filePath, err.Error())
			color.Red("	" + msg)
			value, _ := otto.ToValue(msg)
			return value
		}

		if string(fileContent) == response {
			msg := "Response matches the file content"
			color.HiCyan("	" + msg)
			value, _ := otto.ToValue(msg)
			return value
		} else {
			msg := "Response does not match the file content"
			color.Red("	" + msg)
			value, _ := otto.ToValue(msg)
			return value
		}
	})

	// Define the resultIsLikeFile function
	vm.Set("ResultIsLikeGlobalFile", func(call otto.FunctionCall) otto.Value {
		response, _ := call.Argument(0).ToString()
		filePathGlobalResultRel, _ := call.Argument(1).ToString()
		filePath := filepath.Join(configDir, "globalTestFiles/results", filePathGlobalResultRel)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			msg := fmt.Sprintf("Expected result file does not exist: %s", filePath)
			color.Red("	" + msg)
			value, _ := otto.ToValue(msg)
			return value
		}

		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			msg := fmt.Sprintf("Error reading file %s: %s", filePath, err.Error())
			color.Red("	" + msg)
			value, _ := otto.ToValue(msg)
			return value
		}

		if string(fileContent) == response {
			msg := "Response matches the file content"
			color.HiCyan("	" + msg)
			value, _ := otto.ToValue(msg)
			return value
		} else {
			msg := "Response does not match the file content"
			color.Red("	" + msg)
			value, _ := otto.ToValue(msg)
			return value
		}
	})
}

// ExecuteConfigScript loads and executes the JavaScript configuration file.
func ExecuteConfigScript(vm *otto.Otto, filePath string) error {
	return config.LoadScript(vm, filePath)
}

// Helper function to format console output arguments
func formatForConsole(argumentList []otto.Value) string {
	var output []string
	for _, argument := range argumentList {
		output = append(output, fmt.Sprintf("%v", argument))
	}
	return strings.Join(output, " ")
}

// SetupConsoleLog sets up the console.log function in the Otto VM.
func SetupConsoleLog(vm *otto.Otto) {
	c := color.New(color.FgHiBlue)

	console := map[string]interface{}{
		"log": func(call otto.FunctionCall) otto.Value {
			c.Println("	console.log:", formatForConsole(call.ArgumentList))
			return otto.UndefinedValue()
		},
	}
	vm.Set("console", console)
}
