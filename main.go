package main

import (
	"fmt"
	"os"
)

func main() {
	if !areInputParamsValid() {
		os.Exit(1)
		return
	}
	var errEslint = runEslint()
	if errEslint != nil {
		fmt.Println("runEslint failed: ", errEslint)
		os.Exit(3)
		return
	}
	var err = reportEslintErrors()
	if err != nil {
		fmt.Println("reportEslintErrors:", err)
		os.Exit(2)
		return
	}
	os.Exit(0)
}
