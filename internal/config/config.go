package config

import (
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/robertkrimen/otto"
)

// LoadScript reads and executes a JavaScript file in the given VM.
func LoadScript(vm *otto.Otto, filePath string) error {
	script, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading %s file: %w", filePath, err)
	}

	if _, err := vm.Run(script); err != nil {
		return fmt.Errorf("Error executing %s script: %w", filePath, err)
	}

	return nil
}

// GetSettings retrieves the global settings from the VM.
func GetSettings(vm *otto.Otto) (map[string]interface{}, error) {
	settingsValue, err := vm.Get("Settings")
	if err != nil {
		return nil, fmt.Errorf("Error getting Settings: %w", err)
	}

	settings, err := settingsValue.Export()
	if err != nil {
		return nil, fmt.Errorf("Error exporting Settings: %w", err)
	}

	if settings == nil {
		color.Red("	Warning: Settings not found")
		return nil, nil
	}

	return settings.(map[string]interface{}), nil
}

// GetHeaders retrieves headers from the settings in the provided Otto VM.
func GetHeaders(vm *otto.Otto) (map[string]string, error) {
	settings, err := GetSettings(vm)
	if err != nil {
		return nil, err
	}

	headers, ok := settings["headers"]
	if !ok {
		color.Red("	Warning: Headers not found in Settings")
		return make(map[string]string), nil
	}

	headersMap, ok := headers.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("headers is not a map[string]interface{}")
	}

	headersStringMap := make(map[string]string)
	for key, value := range headersMap {
		strValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("header value for key %s is not a string", key)
		}
		headersStringMap[key] = strValue
	}

	return headersStringMap, nil
}

func GetBaseUrl(vm *otto.Otto) (string, error) {
	settings, err := GetSettings(vm)
	if err != nil {
		return "", err
	}

	if _, ok := settings["baseUrl"]; !ok {
		return "", nil
	}

	baseUrl := settings["baseUrl"].(string)
	return baseUrl, nil
}
