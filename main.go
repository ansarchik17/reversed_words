package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Ansar is the best"
	result := ReverseWords(str)
	fmt.Println(result)
}

func ReverseWords(str string) string {
	str = strings.TrimSpace(str)

	words := strings.Fields(str)
	i := 0
	j := len(words) - 1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}
	out := strings.Join(words, " ")
	return out
}
