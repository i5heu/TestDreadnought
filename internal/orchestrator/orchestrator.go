package orchestrator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/i5heu/TestDreadnought/internal/vm"
)

type TestCaseFileInfo struct {
	RelativeParentFolder string
	Filename             string
	Passed               bool
	Err                  error
}

func RunTests(configDir string, subSet string) error {
	testList, err := GetListOfTestCases(configDir, subSet)
	if err != nil {
		return err
	}

	globalScript := filepath.Join(configDir, "globalScript.js")

	c := color.New(color.FgMagenta)

	for i, testCase := range testList {
		testCaseFile := filepath.Join(configDir, testCase.RelativeParentFolder, testCase.Filename)
		testCaseParentFolder := filepath.Join(configDir, testCase.RelativeParentFolder)

		c.Printf("\n\nRunning test case: %s\n", testCaseFile)

		ottoVM, err := vm.InitializeVM(globalScript, testCaseParentFolder, configDir)
		if err != nil {
			testList[i].Passed = false
			testList[i].Err = fmt.Errorf("error initializing VM: %w", err)
			return fmt.Errorf("error initializing VM: %w", err)
		}

		c.Println("	GlobalScript was run now Executing test case...")

		err = vm.ExecuteConfigScript(ottoVM, testCaseFile)
		if err != nil {
			testList[i].Passed = false
			color.Red("	%s", err)
			testList[i].Err = err
			continue
		}

		testList[i].Passed = true
		color.Green("	-- Passed --")
	}

	SumTestResults(testList)
	return nil
}

func SumTestResults(testList []TestCaseFileInfo) {
	passedByPath := make(map[string][]string)
	failedByPath := make(map[string][]string)

	for _, test := range testList {
		if test.Passed {
			passedByPath[test.RelativeParentFolder] = append(passedByPath[test.RelativeParentFolder], test.Filename)
		} else {
			failedByPath[test.RelativeParentFolder] = append(failedByPath[test.RelativeParentFolder], test.Filename)
		}
	}

	fmt.Println("\n\n\n####### Test results: #######")
	fmt.Printf("Total tests: %d\n", len(testList))

	green := color.New(color.FgGreen)

	green.Println("\nPassed tests by path:")
	for path, files := range passedByPath {
		green.Printf("%s:\n", path)
		for _, file := range files {
			green.Printf("  - %s\n", file)
		}
	}

	red := color.New(color.FgRed)

	red.Println("\nFailed tests by path:")

	if len(failedByPath) == 0 {
		green.Println("  No failed tests")
	}
	for path, files := range failedByPath {
		red.Printf("%s:\n", path)
		for _, file := range files {
			red.Printf("  - %s\n", file)
		}
	}

	fmt.Println("")

	if len(failedByPath) > 0 {
		color.Red("!!! %d tests failed !!!", len(failedByPath))
	} else {
		color.Green("All tests passed")
	}
}

// GetListOfTestCases recursively gets all .js files in the directory and stores them with the relative folder path and filename.
func GetListOfTestCases(globalScript string, subSet string) ([]TestCaseFileInfo, error) {
	var files []TestCaseFileInfo

	rootPath := globalScript
	if subSet != "" {
		rootPath = filepath.Join(globalScript, subSet)
	}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".js") {
			relPath, err := filepath.Rel(globalScript, path)
			if err != nil {
				return err
			}

			if info.Name() == "globalScript.js" {
				return nil
			}

			files = append(files, TestCaseFileInfo{
				RelativeParentFolder: filepath.Dir(relPath),
				Filename:             info.Name(),
			})
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking the path %s: %w", rootPath, err)
	}

	return files, nil
}
