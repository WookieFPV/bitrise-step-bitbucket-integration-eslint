package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Mapping for eslint Severity
func getSeverityString(severityInteger int) string {
	if severityInteger == 0 {
		return "LOW"
	} else if severityInteger == 1 {
		return "MEDIUM"
	} else if severityInteger == 2 {
		return "HIGH"
	} else {
		fmt.Println("invalid severityInteger")
		return "LOW"
	}
}

func getStatsFromIssues(issues []Issues) ([]Annotation, int, int) {
	var totalErrorCount = 0
	var totalWarningCount = 0
	var annotations []Annotation
	// The filePath is absolute, but Bitbucket Server requires it to be relative to the git repository
	basePath := os.Getenv("BITRISE_SOURCE_DIR") + "/"
	for i := 0; i < len(issues); i++ {
		totalErrorCount += issues[i].ErrorCount
		totalWarningCount += issues[i].WarningCount
		for x := 0; x < len(issues[i].Messages); x++ {
			var anno = issues[i].Messages[x]
			var path string
			path = strings.ReplaceAll(issues[i].FilePath, basePath, "")
			annotations = append(annotations, Annotation{Path: path, Line: anno.Line, Message: anno.Message, Severity: getSeverityString(anno.Severity)})
			fmt.Println("Messages :", issues[i].Messages[x])
		}
	}
	return annotations, totalErrorCount, totalWarningCount
}

func createReport(totalErrorCount int, totalWarningCount int) Report {
	var report Report
	report.Title = "ESLint Report"
	report.Vendor = "WookieFPV"
	report.LogoURL = "https://upload.wikimedia.org/wikipedia/en/e/e3/ESLint_logo.svg"
	report.Data = append(report.Data, EslintReportData{Title: "warnings", Value: totalWarningCount})
	report.Data = append(report.Data, EslintReportData{Title: "errors", Value: totalErrorCount})
	report.Link = "https://www.youtube.com/watch?v=oHg5SJYRHA0"
	report.Details = "Hello World, here could be alot of text"

	if totalErrorCount == 0 {
		report.Result = "PASS"
	} else {
		report.Result = "FAIL"
	}
	return report
}

func reportEslintErrors() error {
	lintFile := os.Getenv("BITRISE_SOURCE_DIR") + "/lint.json"
	jsonFile, err := os.Open(lintFile)
	if err != nil {
		fmt.Println("os.Open(lintFile) failed", err, lintFile)
		return err
	}
	fmt.Println("Successfully Opened lint file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var issues []Issues

	json.Unmarshal(byteValue, &issues)

	annotations, totalWarningCount, totalErrorCount := getStatsFromIssues(issues)

	// COMPUTED VALUES
	var token string = "Bearer " + os.Getenv("BITBUCKET_SERVER_TOKEN")
	var url string = os.Getenv("BITBUCKET_SERVER_URL") + "rest/insights/1.0/projects/" + os.Getenv("PROJECT_ID") + "/repos/" + os.Getenv("BITRISEIO_GIT_REPOSITORY_SLUG") + "/commits/" + os.Getenv("BITRISE_GIT_COMMIT") + "/reports/" + os.Getenv("REPORT_NAME")
	var annotationsURL string = url + "/annotations"
	report := createReport(totalErrorCount, totalWarningCount)

	fmt.Println("token", token)
	fmt.Println("url", url)

	err = putReport(url, token, report)
	if err != nil {
		return err
	}
	err = deleteAnnotations(annotationsURL, token)
	if err != nil {
		return err
	}
	err = postAnnotations(annotationsURL, token, annotations)
	if err != nil {
		return err
	}
	return nil
}
