// extensions.go is licensed under the MIT License, described in the LICENSE file in the same folder as this file.
// Changes to extensions.go do not need to be published or fall under the AGLP-3.0 license.

package extensions

import (
	"path"
	"plugin"

	"github.com/robertkrimen/otto"
)

type ExtensionV1 struct {
}

func (*ExtensionV1) SetUpExtensions(vm *otto.Otto, testCaseParentFolder, configDir string) {
	panic("SetUpExtensions not implemented in Plugin")
}

func LoadExtensions(vm *otto.Otto, testCaseParentFolder, configDir string) error {
	extentionsFolder := path.Join(configDir, "extensions/out/extensions.so")
	plugin, err := plugin.Open(extentionsFolder)
	if err != nil {
		return err
	}

	setUp, err := plugin.Lookup("SetUpExtensions")
	if err != nil {
		return err
	}

	setUp.(func(*otto.Otto, string, string))(vm, testCaseParentFolder, configDir)

	return nil
}
