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
	// Read input
	reader := bufio.NewScanner(&input)
	first := true
	// Initialize counters
	var prev int
	var count int

	// Main loop
	//Advance scanner as iterator
	for reader.Scan() {
		//Intial setup, initial previous
		if first {
			fmt.Println("First Line: " + reader.Text())
			prev, _ = strconv.Atoi(reader.Text())
			first = false
			continue
		}

		//Convert to  int for comparison
		curr, err := strconv.Atoi(reader.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Compare to previous value and increase value
		if curr > prev {
			fmt.Println(reader.Text())
			count++
		}
		// Update previous counter
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

func dayOne() {
	file, _ := ioutil.ReadFile("sweep.input")
	input := string(file)
	//sweep(*file)
	nums := toNums(input)
	window(nums)
}

func dive(lines []string) {
	cords := make(map[string]int)
	for line := range lines {
		command := strings.Split(strings.TrimSpace(lines[line]), " ")
		dir := command[0]
		mag, _ := strconv.Atoi(command[1])
		if dir == "up" {
			cords["down"] = cords["down"] - mag
			continue
		}
		cords[strings.TrimSpace(dir)] = cords[dir] + mag

	}
	fmt.Println(cords["down"] * cords["forward"])
}

func aim(lines []string) {
	cords := make(map[string]int)
	for line := range lines {
		command := strings.Split(strings.TrimSpace(lines[line]), " ")
		dir := command[0]
		mag, _ := strconv.Atoi(command[1])
		switch dir {
		case "up":
			cords["aim"] = cords["aim"] - mag
		case "down":
			cords["aim"] = cords["aim"] + mag
		case "forward":
			cords["forward"] = cords["forward"] + mag
			cords["down"] = cords["aim"]*mag + cords["down"]
		}
	}
	fmt.Println(cords["down"] * cords["forward"])
}

func dayTwo() {
	file, _ := ioutil.ReadFile("dive.input")
	input := string(file)
	lines := strings.Split(input, "\r")
	dive(lines)
	aim(lines)

}

func diag(lines []string) {
	rows := make(map[int]int)
	for line := range lines {

	}
}

func dayThree() {
	file, _ := ioutil.ReadFile("diagnostic.input")
	input := string(file)
	lines := strings.Split(input, "\r")
	diag(lines)
}
func main() {
	dayThree()
}
