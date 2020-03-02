package main

import (
	"fmt"
	"os"
)

func checkEnvVar(enVarName string, missingEnvVars []string) []string {
	if os.Getenv(enVarName) != "" {
		fmt.Println("ENV VAR: ", enVarName, ": ", os.Getenv(enVarName))
		return missingEnvVars
	}
	fmt.Println("WARNING ENV VARIABLE NOT SET: ", enVarName)
	return append(missingEnvVars, enVarName)
}

func areInputParamsValid() bool {
	var missingEnvVars []string
	missingEnvVars = checkEnvVar("BITBUCKET_SERVER_URL", missingEnvVars)
	missingEnvVars = checkEnvVar("BITBUCKET_SERVER_TOKEN", missingEnvVars)
	missingEnvVars = checkEnvVar("PROJECT_ID", missingEnvVars)
	missingEnvVars = checkEnvVar("BITRISE_GIT_COMMIT", missingEnvVars)
	missingEnvVars = checkEnvVar("REPORT_NAME", missingEnvVars)
	missingEnvVars = checkEnvVar("BITRISE_SOURCE_DIR", missingEnvVars)

	//if list of missingEnvVars is empty arething is good to go
	if len(missingEnvVars) != 0 {
		fmt.Println("Some ENV VARIABLES are not set", missingEnvVars)
		return false
	}
	return true
}
