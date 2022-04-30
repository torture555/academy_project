package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var tabooWords []string          // initialization slice with taboo words
	sc := bufio.NewScanner(os.Stdin) // initialization scanner for input

	for sc.Scan() { // while input values

		var inputValue = sc.Text()

		if inputValue == "exit" { //exit from programm
			fmt.Println("Bye!")
			break
		} else if strings.HasSuffix(inputValue, ".txt") {

			tabooWords = nil // clear old slice with taboo words

			file, err := os.Open(inputValue) // open file with taboo words
			if err != nil {
				fmt.Println(err)
				break
			}

			scanTxt := bufio.NewScanner(file) //scaning file and append taboo words in slice
			for scanTxt.Scan() {
				if strings.HasPrefix(scanTxt.Text(), ">") {
					continue
				}
				tabooWords = append(tabooWords, strings.ToLower(scanTxt.Text()))
			}

		} else {

			outputPrint := inputValue
			tmpStr := outputPrint
			tmpStr = strings.ToLower(tmpStr)

			for _, tabooWord := range tabooWords {
				tmpStr = strings.ReplaceAll(tmpStr, strings.ToLower(tabooWord), getAsteriskWord(tabooWord))
			}

			if strings.Contains(tmpStr, "*") {
				for i, letter := range tmpStr {
					if letter == rune(42) {
						outputPrint = outputPrint[:i] + "*" + outputPrint[i+1:]
					}
				}
			}

			fmt.Println(outputPrint)
		}
	}
}

func getAsteriskWord(str string) string {
	outputStr := ""
	for i := 0; i < len(str); i++ {
		outputStr += "*"
	}
	return outputStr
}
