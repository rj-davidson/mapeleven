package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadImageIfNeeded(url string, filename string) (string, error) {
	// Prepend the public directory to the filename
	filename = "public/" + filename

	// Check if the file already exists
	if _, err := os.Stat(filename); err == nil {
		fmt.Printf("File already exists: %s\n", filename)
		return filename[7:], nil // Remove the public directory from the filename
	}

	// Create the directory if it doesn't exist
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
	}

	// Download the image
	fmt.Printf("local --> Downloading image: %s\n", filename)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Printf("Downloaded image: %s\n", filename)

	// Remove the public directory from the filename
	filename = filename[7:]

	return filename, nil
}
