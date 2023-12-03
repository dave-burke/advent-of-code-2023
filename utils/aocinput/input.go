package aocinput

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const year = 2023

func ReadSampleAsString(day int) string {
	return readFileAsString(sampleFileName(day))
}

func ReadInputAsString(day int) string {
	downloadInputIfNeeded(day)
	return readFileAsString(inputFileName(day))
}

func readFileAsString(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func ReadSampleAsLines(day int) []string {
	return readFileAsLines(sampleFileName(day))
}

func ReadInputAsLines(day int) []string {
	downloadInputIfNeeded(day)
	return readFileAsLines(inputFileName(day))
}

func readFileAsLines(path string) []string {
	content := readFileAsString(path)
	return strings.Split(content, "\n")
}

func ReadSampleAsChannel(day int) chan string {
	return readFileAsChannel(sampleFileName(day))
}

func ReadInputAsChannel(day int) chan string {
	downloadInputIfNeeded(day)
	return readFileAsChannel(inputFileName(day))
}

func readFileAsChannel(path string) chan string {
	lines := make(chan string)

	go func() {
		file, err := os.Open(path)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("SEND: %s\n", line)
			lines <- line
		}
		log.Println("Closing lines channel")
		close(lines)
	}()

	return lines
}

func inputFileName(day int) string {
	return fmt.Sprintf("solutions/day%02d/input.txt", day)
}
func sampleFileName(day int) string {
	return fmt.Sprintf("solutions/day%02d/sample.txt", day)
}

func downloadInputIfNeeded(day int) {
	path := inputFileName(day)
	if _, err := os.Stat(path); err == nil {
		// file exists
		return
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	session := fmt.Sprintf("session=%s", readSessionFile())

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cookie", session)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("%d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Save file
	err = os.WriteFile(path, bodyBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readSessionFile() string {
	content, err := os.ReadFile("session.txt")
	if err != nil {
		log.Fatalf("Log in to adventofcode.com and then use the browser debug tools to get your session cookie and save it as 'session.txt'. (%s)", err)
	}
	return string(content)
}
