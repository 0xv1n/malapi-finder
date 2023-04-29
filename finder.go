package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

const pattern = `(?s)LPVOID\s+(\w+)\s*\(\s*HANDLE\s+\w+\s*,\s*LPVOID\s+\w+\s*,\s*SIZE_T\s+\w+\s*,\s*DWORD\s+\w+\s*,\s*DWORD\s+\w+\s*\)`

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <source_code_file>", os.Args[0])
	}

	sourceCodeFile := os.Args[1]

	data, err := ioutil.ReadFile(sourceCodeFile)
	if err != nil {
		log.Fatal(err)
	}

	// Remove comments
	dataWithoutComments := regexp.MustCompile(`(?:\/\/[^\r\n]*[\r\n]?)|(?:\/\*[\s\S]*?\*\/)`).ReplaceAllString(string(data), "")

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(dataWithoutComments, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			functionName := match[1]
			fmt.Printf("Found matching function signature: %s\n", functionName)

			parameterList := strings.Split(match[0], "(")[1]
			parameterList = strings.Split(parameterList, ")")[0]
			parameters := strings.Split(parameterList, ",")

			var parsedParameters []string
			for _, param := range parameters {
				parsedParameters = append(parsedParameters, strings.TrimSpace(param))
			}
			fmt.Printf("Parsed parameters: %v\n", parsedParameters)
		}
	} else {
		fmt.Println("No matching function signature found.")
	}
}
