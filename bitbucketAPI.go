package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func putReport(URL string, token string, report Report) error {
	body := &report

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	req, _ := http.NewRequest("PUT", URL, buf)

	req.Header.Set("Authorization", token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, e := client.Do(req)
	if e != nil {
		fmt.Println("putReport failed", e)
		return e
	}
	defer res.Body.Close()
	fmt.Println("putReport response Status:", res.Status)
	io.Copy(os.Stdout, res.Body) // Print the body to the stdout
	return nil
}

func deleteAnnotations(annotationsURL string, token string) error {
	req, _ := http.NewRequest("DELETE", annotationsURL, new(bytes.Buffer))

	req.Header.Set("Authorization", token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		fmt.Println("deleteAnnotations failed", e)
		return e
	}

	defer res.Body.Close()
	fmt.Println("deleteAnnotations response Status:", res.Status)
	io.Copy(os.Stdout, res.Body) // Print the body to the stdout
	return nil
}

func postAnnotations(annotationsURL string, token string, annotations []Annotation) error {
	fmt.Println("annotations", annotations)

	body := &BitbucketAnnotations{Annotations: annotations}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	req, _ := http.NewRequest("POST", annotationsURL, buf)

	req.Header.Set("Authorization", token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		fmt.Println("postAnnotations failed", e)
		return e
	}
	defer res.Body.Close()
	fmt.Println("postAnnotations response Status:", res.Status)

	io.Copy(os.Stdout, res.Body) // Print the body to the stdout
	return nil
}
