package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getESLintPath() (bool, string) {
	isExist := true
	path, err := exec.LookPath("eslint")
	if err != nil {
		path, err = exec.LookPath("node_modules/eslint/bin/eslint.js")
		if err != nil {
			fmt.Println("First of all, install eslint")
			isExist = false
		}
	}
	return isExist, path
}

func runESLint(eslintPath string, srcDirectory string) (bool, error) {
	nodePath, err := exec.LookPath("node")
	fmt.Println("running eslint...", nodePath, " ", eslintPath, " ", srcDirectory)
	output, err := exec.Command(nodePath, eslintPath, srcDirectory, "--f", "json", "--o", "lint.json").CombinedOutput()
	fmt.Printf("Output: %s\n", string(output[:]))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		if strings.Contains(err.Error(), "1") { // Linting worked but had errors
			return false, nil
		}
		// Linting was not possible
		return false, err
	}
	return true, nil
}

func runEslint() (bool, error) {
	srcDirectory := os.Getenv("BITRISE_SOURCE_DIR") + "/"
	isExist, eslintPath := getESLintPath()
	if !isExist {
		return false, errors.New("ESLint was not found")
	}
	return runESLint(eslintPath, srcDirectory)
}
