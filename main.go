package main

import (
	"aoc/solutions/day01"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const year = 2023
const day = 1
const inputBaseName = "input.txt"

func main() {
	content := getInput(day)
	result := day01.Part2(content)
	fmt.Printf("Solution: %s", result)
}

func getInput(day int) string {
	content, err := readInputFile(day)
	if err != nil {
		content, err = readInputHttp(day)
		if err != nil {
			log.Fatal(err)
		}
		saveInputFile(day, content)
	}
	return string(content)
}

func readInputHttp(day int) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	session := fmt.Sprintf("session=%s", readSessionFile())

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cookie", session)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf(string(bodyBytes))
	}

	return bodyBytes, nil
}

func inputFileName(day int) string {
	return fmt.Sprintf("solutions/day%02d/%s", day, inputBaseName)
}

func readInputFile(day int) ([]byte, error) {
	path := inputFileName(day)
	content, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return content, nil
}

func saveInputFile(day int, content []byte) {
	path := inputFileName(day)
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readSessionFile() string {
	content, err := os.ReadFile("session.txt")
	if err != nil {
		log.Fatalf("Log in to adventofcode.com and then use the browser debug tools to get your session cookie. (%s)", err)
	}
	return string(content)
}
