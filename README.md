<p align="center">
  <img src=".media/logo.webp"  width="300">
</p>

# TestDreadnought

Highly adoptable test Framework for large and complex projects, that need a dependency-conscious, secure and fast testing solution.  

TestDreadnought is built for you to mold precisely to your project's demands, avoiding the one-size-fits-all approach. It sidesteps dependency issues and the security risks of excessive third-party packages, particularly from JS npm. With no forced structure or excessive rules, you get a clean, efficient, and secure testing environment that’s exactly suited to tackle your unique challenges.

## Install
> We will soon provide a docker image for easier installation. Which will also enable have the extensions in the same folder as the tests. 

Since you need to build the extensions in Go, you need to have Go installed on your machine. Make sure you have the `$GOPATH/bin` in your `$PATH` environment variable - See [instructions](https://stackoverflow.com/a/21012349).  
Best is to clone the repo, insert your extensions and build and install it via `go install`.

We advise to use a git submodule for the extensions folder, so you can easily update the extensions without having to update the whole TestDreadnought repository while reducing the risk of accidentally pushed sensitive data.

```bash
git clone https://github.com/i5heu/TestDreadnought.git
# Optional: Add the extensions folder as a submodule
cd TestDreadnought
go install
```
Now you can use `TestDreadnought` in your terminal.
```bash
TestDreadnought <test-root-directory> <optional: subset path relative to config-directory>
```

If you pulled the repos for updated, you need to run `go install` again to update the binary.

## ClI Usage
There are 2 CLI options, the test-root-directory and the optional "subset path" that is relative to the test-root-directory.  
If you provide no arguments, TestDreadnought will show this message:

```bash
$ TestDreadnought
TestDreadnought Usage: TestDreadnought <test-root-directory> <optional: subset path relative to config-directory>
```


## Setup
There must be a `globalScript.js` file in the root directory of your test folder.  
This script is called before any `.js` file in the test root directory and its subdirectories.  
It is best used to define global variables and functions that are used in multiple tests.  

We suggest you put into here a variable ´globalSettings´ that contains, for example, the base URL of your API.

```js
// testRootDirectory/globalScript.js
var globalSettings = { // This is a global variable that can be used in all tests
    baseUrl: "https://example.com", // This is the base URL of your API 
    headers: {
        "Origin": "https://this.is.a.example.com"
    }
};

var ThisIsTest = function () {
    console.log("This is a test"); // This will be callable in all Tests.
}

console.log("Global settings loaded"); // This will be printed in the console before executing an test
```

## Writing your first test

TestDreadnought uses ES5 JavaScript as language for writing tests.  
We choose JS as it is widely used and allows for more familiarity and easier onboarding.  
As you will see, JS is preferably only used as a kind of routing for data and simple logic that is test case relevant, all requests in the default HTTP requests functions are made in golang and callable via fake JS functions, aka extensions.  

Lets write your first test, that will Get a request to the base URL of your API.
```js
// testRootDirectory/test/helloWorld.js - note that there is no forced structure
// TestDreadnought it will test all .js files in all directories of the test folder,
// unless you specify to test only a subset

// This is the global variable from the globalScript.js file
Settings = globalSettings; 

// This will make a GET request to the base URL and the path /helloWorld using the headers defined in the Settings variable
result = Get("/helloWorld");

// This will print the result of the request to the console
console.log("Cache-Control:", result.header["Cache-Control"]);  

// if no error is thrown the test is successful
```

Note that the Settings variable is used by the Get function and all default client functions to get the base URL and headers.  
The structure for this looks like this:
```js
Settings = {
	baseUrl: "https://example.com",
	headers: {
		"Origin": "https://this.is.a.example.com"
	}
}
```

For more examples checkout the `test_example` folder in this repository.

## Building your own extensions  

Extensions are a way to add custom functionality to TestDreadnought.  
They are meant for more complex steps and time sensitive measurements, like performance testing.  
We use Go for extensions since for a lot of tests JS is not precise enough or capable enough in a elegant way and without a lot of third party packages, which are a security risk.  

To build an extension you need to create a `.go` file in the `extensions` folder, using the package name `extensions`.
After that you need to add your new custom function to the `SetUpExtensions` function of the `extensions.go` file. 
Pls note that for the `extensions.go` file the MIT license applies. So you can change it without having to open source it.  

```go
// extensions/helloWorld.go
package extensions

import	"github.com/robertkrimen/otto"

func exampleHelloWorld(vm *otto.Otto, testCaseParentFolder, configDir string) {
	vm.Set("ExampleHelloWorld", func(call otto.FunctionCall) otto.Value {
		incomingValue := call.Argument(0).String()

		Log("helloWorld", incomingValue)
		back := "Hello World Back!"

		value, err := vm.ToValue(back)
		if err != nil {
			panic(err)
		}

		return value
	})
}
```

Now you need the to add the `exampleHelloWorld` function to the `SetUpExtensions` function of the `extensions.go` file.
```go
// extensions/extensions.go

func SetUpExtensions(vm *otto.Otto, testCaseParentFolder, configDir string) {
	exampleHelloWorld(vm, testCaseParentFolder, configDir)
}
```

Now you can call the `ExampleHelloWorld` function in all your tests.
```js
// testRootDirectory/test/helloWorld.js
console.log(ExampleHelloWorld("FooBar"))
```

## Demo Video

https://github.com/i5heu/TestDreadnought/assets/22565269/e8b5398a-e990-44e4-937c-d08638409c51
