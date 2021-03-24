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
	var err = reportEslintErrors()
	if err != nil {
		fmt.Println("reportEslintErrors:", err)
		os.Exit(2)
		return
	}
}
