package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type Package struct {
	Name string `json:"name"`
	Tags [] string `json:"keywords"`
};

func ReadPackage(directory string) (*Package, error) {
	jsonFile, err := os.Open(path.Join(directory, "package.json"));

	if err != nil {
		return nil, err;
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err;
	}

	var result Package

	json.Unmarshal([]byte(byteValue), &result)

	return &result, nil;
}