package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	// ClientID : client id
	ClientID = "0kZtO5NQebBCXVpOEkmcIF9t5pGXrkwX"

	// ClientSecret : client scret key
	ClientSecret = "A9l4iSeHT4dkcpmJ3VDHA5Wg6jMgZo4CvZM05ppJKbg9HPldwsmuIfRUDLhU1sEa"

	// ConfigFilePath : config file path
	ConfigFilePath = `.backlog.config`
)

// WriteConfig : save to config file
func WriteConfig(model Config) error {
	file, err := os.Create(ConfigFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonBytes, err := json.Marshal(model)
	if err != nil {
		return err
	}

	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "    ")
	file.Write(([]byte)(out.String()))

	return nil
}

// ReadConfig : read config file
func ReadConfig() (Config, error) {
	var model Config

	bytes, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		return model, err
	}
	if err := json.Unmarshal(bytes, &model); err != nil {
		return model, err
	}

	return model, nil
}
