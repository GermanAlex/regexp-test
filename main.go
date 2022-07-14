package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	outputFile, err := os.OpenFile("./out.txt", os.O_RDWR|os.O_CREATE, 0777)
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
		writer.Write([]byte(getStrResult(scanner.Text() + "\n")))
		writer.Flush()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func getStrResult(str string) string {
	var result string
	var res int
	re := regexp.MustCompile(`^[0-9]*[-+*\/][0-9]*[=][?][\n]$`)
	/*Выражения для разбора*/
	numOneR := regexp.MustCompile(`^[0-9]*`)
	operatorR := regexp.MustCompile(`[-+*\/]`)
	numTwoR := regexp.MustCompile(`[-+*\/][0-9]*`)
	resultR := regexp.MustCompile(`[?]`)
	submatch := re.FindStringSubmatch(str)
	if submatch != nil {
		for _, s := range submatch {
			fmt.Println(s)
			fmt.Println(numOneR.FindString(s))
			numOne, err := strconv.Atoi(numOneR.FindString(s))
			if err != nil {
				panic(err)
			}
			numTwo, err := strconv.Atoi(numTwoR.FindString(s)[1:])
			if err != nil {
				panic(err)
			}
			operator := operatorR.FindString(s)
			switch operator {
			case "+":
				res = numOne + numTwo
			case "-":
				res = numOne - numTwo
			case "*":
				res = numOne * numTwo
			case "/":
				res = numOne / numTwo
			}
			result = resultR.ReplaceAllString(s, strconv.Itoa(res))
		}
	} else {
		result = ""
	}
	return result
}
