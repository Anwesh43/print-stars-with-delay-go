package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
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

func printStars(n int64, ch chan bool) {
	fmt.Println("Starting to print", n, "stars")

	for i := 0; i < int(n); i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
			time.Sleep(400 * time.Millisecond)
		}
		fmt.Println()
	}
	fmt.Println("Done printing", n, "stars")

	ch <- true
}

func printNStars(nums []int64, ch chan bool) {

	for _, n := range nums {
		boolch := make(chan bool)
		go printStars(n, boolch)
		<-boolch
		close(boolch)
	}

	ch <- true
	close(ch)
}

func main() {
	ch := make(chan []int64)
	go getStarsInput(ch)
	nums := <-ch
	starCh := make(chan bool)
	go printNStars(nums, starCh)
	<-starCh
}
