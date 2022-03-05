package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func sweep(input os.File) {
	reader := bufio.NewScanner(&input)
	first := true
	var prev int
	var count int
	for reader.Scan() {
		if first {
			fmt.Println("First Line: " + reader.Text())
			prev, _ = strconv.Atoi(reader.Text())
			first = false
			continue
		}

		curr, err := strconv.Atoi(reader.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if curr > prev {
			fmt.Println(reader.Text())
			count++
		}
		prev = curr

	}
	print(count)
}

func window(lines []int) {
	count := 0
	for i := 0; i < len(lines)-3; i++ {
		sum1 := lines[i] + lines[i+1] + lines[i+2]
		sum2 := lines[i+1] + lines[i+2] + lines[i+3]

		if sum1 < sum2 {
			count += 1
		}

	}
	fmt.Println(count)
}

func toNums(input string) []int {

	lines := strings.Split(input, "\n")
	nums := make([]int, len(lines))
	for i, raw := range lines {
		num, _ := strconv.Atoi(strings.TrimSuffix(raw, "\r"))
		nums[i] = num

	}
	return nums
}

func main() {
	file, _ := ioutil.ReadFile("sweep.input")
	input := string(file)
	//sweep(*file)
	nums := toNums(input)
	window(nums)
}
