package main

import (
	fmtLog "github.com/i5heu/TestDreadnought/pkg"
	"github.com/robertkrimen/otto"
)

func SetUpExtensions(vm *otto.Otto, testCaseParentFolder, configDir string) {
	exampleHelloWorld(vm, testCaseParentFolder, configDir)
}

func exampleHelloWorld(vm *otto.Otto, testCaseParentFolder, configDir string) {
	vm.Set("ExampleHelloWorld", func(call otto.FunctionCall) otto.Value {
		incomingValue := call.Argument(0).String()

		fmtLog.Log("helloWorld", incomingValue)
		back := "Hello World Back!"

		value, err := vm.ToValue(back)
		if err != nil {
			panic(err)
		}

		return value
	})
}
