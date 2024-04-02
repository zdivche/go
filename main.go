// Статистика текстового файла

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func mostFreqLetter(text string) (rune, int) {
	m := make(map[rune]int)
	for _, v := range text {
		m[v]++
	}
	var maxCount int
	var maxRune rune
	for k, v := range m {
		if v > maxCount {
			maxCount = v
			maxRune = k
		}
	}
	return maxRune, maxCount
}

func mostFreqWord(text string) (string, int) {
	m := make(map[string]int)
	words := strings.Fields(text)
	for _, v := range words {
		m[v]++
	}
	var maxCount int
	var maxWord string
	for k, v := range m {
		if v > maxCount {
			maxCount = v
			maxWord = k
		}
	}
	return maxWord, maxCount
}


func counterChars(text string) int {
	count := 0
	for i := 0; i < len(text); i++ {
		count++
	}
	return count
}
func Print(text string) {
	fmt.Println("Количество слов: ", counterWords(string(text)), "\n")
	fmt.Print("Количество символов: ", counterChars(string(text)), "\n")

	mostFreqLetter, freq := mostFreqLetter(string(text))
	fmt.Println("Самая частая буква: ", mostFreqLetter, "Количество вхождений: ", freq, "\n")
	mostFreqWord, freq := mostFreqWord(string(text))
	fmt.Println("Самое частое слово: ", mostFreqWord, "Количество вхождений: ", freq, "\n")
}
func counterWords(text string) int {
	count := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			count++
		}
	}
	return count
}

func main() {
	
	text, err := ioutil.ReadFile("sometext.txt")
	if err != nil {
		fmt.Println(err, "\n")
	}

	Print(string(text))
	
	
}
