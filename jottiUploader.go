package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
Jotti Uploader - Tool to upload files to https://virusscan.jotti.org
Jotti is an alternative to VirusTotal
by cyclone
https://github.com/cyclone-github/jottiUploader
change log:
v2023-11-10.1800; initial version
v2023-11-11.1800; cleaned up code; github release
*/

// global variables
var (
	jottiUploadURL   = "https://virusscan.jotti.org/en-US/submit-file"
	jottiChecksumURL = "https://virusscan.jotti.org/en-US/search/hash/%s"
)

func versionFunc() {
	fmt.Fprintln(os.Stderr, "Jotti Uploader v2023-11-11.1800")
	fmt.Fprintln(os.Stderr, "https://github.com/cyclone-github/jottiUploader")
}

// help function
func helpFunc() {
	versionFunc()
	str := "\nExample Usage:\n" +
		"\n./jottiUploader {file_to_scan}\n" +
		"\n./jottiUploader -help\n" +
		"\n./jottiUploader -version\n"
	fmt.Fprintln(os.Stderr, str)
	os.Exit(0)
}

// calculate SHA1 checksum of file
func calculateSHA1Checksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// upload file to Jotti
func uploadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("sample-file[]", filePath)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", jottiUploadURL, body)
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response status: %d", response.StatusCode)
	}

	return "", nil
}

// check if SHA1 checksum exists on Jotti
func checkJottiSearch(checksum string) (bool, string, error) {
	searchURL := fmt.Sprintf(jottiChecksumURL, checksum)
	response, err := http.Get(searchURL)
	if err != nil {
		return false, "", err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return false, "", err
		}
		body := string(bodyBytes)

		if strings.Contains(body, "Too many requests") {
			// rate limit detected, exit
			fmt.Println("Rate limited by Jotti. Please try again in a few minutes.")
			os.Exit(0)
		}

		// search for "Hash not found" string
		if strings.Contains(body, "Hash not found") {
			return false, searchURL, nil
		}
		return true, searchURL, nil
	}

	return false, "", fmt.Errorf("unexpected response status: %d", response.StatusCode)
}

func main() {
	help := flag.Bool("help", false, "Prints help:")
	version := flag.Bool("version", false, "Program Version:")
	cyclone := flag.Bool("cyclone", false, "")
	flag.Parse()
	if *version {
		versionFunc()
		os.Exit(0)
	}
	if *cyclone {
		fmt.Fprintln(os.Stderr, "Coded by cyclone ;)")
		os.Exit(0)
	}

	// check for file in cli
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./jottiUploader <file_to_scan>")
	}
	if *help {
		helpFunc()
	}

	// loop over each file
	for _, filePath := range os.Args[1:] {
		// calculate SHA1 checksum of file
		checksum, err := calculateSHA1Checksum(filePath)
		if err != nil {
			log.Printf("Error calculating SHA1 checksum for %s: %v\n", filePath, err)
			continue
		}
		fmt.Printf("SHA1 Checksum: %s\n", checksum)

		// check if SHA1 checksum is on Jotti
		found, jottiURL, err := checkJottiSearch(checksum)
		if err != nil {
			log.Printf("Error checking Jotti's malware scan: %v\n", err)
			continue
		}

		if found {
			fmt.Printf("File %s found on Jotti:\n%s\n", filePath, jottiURL)
			continue // skip to next file if found
		}

		fmt.Printf("Uploading %s: ", filePath)
		_, err = uploadFile(filePath)
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}
		fmt.Print("OK\n", fmt.Sprintf(jottiChecksumURL, checksum+"\n"))

		// wait for nth sec
		time.Sleep(1000 * time.Millisecond)
	}
}

// end code
