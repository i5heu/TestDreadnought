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
