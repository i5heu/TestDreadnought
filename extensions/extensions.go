// extensions.go is licensed under the MIT License, described in the LICENSE file in the same folder as this file.
// Changes to extensions.go do not need to be published or fall under the AGLP-3.0 license.

package extensions

import (
	"github.com/fatih/color"
	"github.com/robertkrimen/otto"
)

func SetUpExtensions(vm *otto.Otto, testCaseParentFolder, configDir string) {

}

func Log(a ...interface{}) (n int, err error) {
	clr := color.New(color.FgCyan)
	inter := []interface{}{"	Extention: "}
	args := append(inter, a...)
	return clr.Println(args...)
}
