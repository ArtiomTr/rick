package main

import (
	"os"
	"rick/pkg/core"

	"github.com/sirupsen/logrus"
)

func main() {
	workingDir, _ := os.Getwd()
	pkg, err := core.ReadPackage(workingDir)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info(pkg.Tags)
	
}
