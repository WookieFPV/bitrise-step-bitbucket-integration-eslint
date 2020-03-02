package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("bitrise-step-bitbucket-integration:")
	if !areInputParamsValid() {
		os.Exit(1)
		return
	}
	var errEslint = runEslint()
	if errEslint != nil {
		fmt.Println(errEslint)
		os.Exit(3)
		return
	}
	var err = reportEslintErrors()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
		return
	}
	os.Exit(0)
}
