package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	resp, err := http.Get("https://yandex.ru")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Get the response body as a string
	dataInBytes, err := ioutil.ReadAll(resp.Body)
	pageContent := string(dataInBytes)

	// Find a substr
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		return
	}

	// The start index of the title is the index of the first
	// character, the < symbol. We don't want to include
	// <title> as part of the final value, so let's offset
	// the index by the number of characers in <title>
	titleStartIndex += 7

	// Find the index of the closing tag
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleEndIndex == -1 {
		fmt.Println("No closing tag for title found.")
		return
	}

	// (Optional)
	// Copy the substring in to a separate variable so the
	// variables with the full document data can be garbage collected
	pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])
	elapsed := time.Since(start)

	log.Printf("\n%s", elapsed)
	fmt.Printf("Page title: %s\n", pageTitle) // Page title: Яндекс
}
