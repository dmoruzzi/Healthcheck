package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Check for the correct number of command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: healthcheck.exe [url] [logPath]")
		os.Exit(1)
	}

	// Extract command-line arguments
	url := os.Args[1]
	logPath := os.Args[2]

	// Open the log file in append mode
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %s", err)
	}
	defer logFile.Close()

	// Create a new logger instance that writes to the log file
	logger := log.New(logFile, "", log.Ldate|log.Ltime)

	// Fetch the URL
	response, err := http.Get(url)
	if err != nil {
		logError(logger, url, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		logSuccess(logger, url, response.Status)
	} else {
		logError(logger, url, fmt.Errorf("HTTP error: %s", response.Status))
	}
}

func logError(logger *log.Logger, url string, err error) {
	logger.SetPrefix("[ERROR] ")
	logger.Printf("%s: %s", url, err)
}

func logSuccess(logger *log.Logger, url, status string) {
	logger.SetPrefix("[OK] ")
	logger.Printf("%s: %s\n", url, status)
}
