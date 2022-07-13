package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	outputFile, err := os.OpenFile("out.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	regexpMath(inputFile, outputFile)

}

func regexpMath(input *os.File, output *os.File) {
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	for scanner.Scan() {
		writer.Write([]byte(getStrResult(scanner.Text()) + "\n"))
		writer.Flush()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func getStrResult(str string) string {
	var result string
	re := regexp.MustCompile(`^[0-9]*[-+*\/][0-9]*[=][?]$`)
	submatch := re.FindStringSubmatch(str)
	fmt.Println(submatch)
	if submatch != nil {
		for _, s := range submatch {
			result = s
		}
	} else {
		result = re.ReplaceAllString(str, "")
	}
	fmt.Println(result)
	return result
}
