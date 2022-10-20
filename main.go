package main

import (
	"bufio"
	"os"
	"strconv"
)

func mapWordsToNumber(words []string) []int64 {
	nums := make([]int64, 0)
	for _, word := range words {
		num, err := strconv.ParseInt(word, 10, 64)
		if err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

func getStarsInput(ch chan []int64) {
	words := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "QUIT" {
			break
		}
		words = append(words, text)
	}
	ch <- mapWordsToNumber(words)
}

func main() {
	ch := make(chan []int64)
	go getStarsInput(ch)
	nums := <-ch

}
