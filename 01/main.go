package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"aoc23/pkg/load_input"
)

const file = "input.txt"

func main() {
	input, err := load_input.Load(file)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	part_one(input)
	part_two(input)
}

func part_one(input []string) {
	var regex = []string{"[1-9]"}
	answer := get_answer(regex, input)
	fmt.Print("Part one: ")
	fmt.Println(answer)
}

func part_two(input []string) {
	var regex = []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	answer := get_answer(regex, input)
	fmt.Print("Part two: ")
	fmt.Println(answer)
}

func get_answer(regex []string, input []string) int {

	answer := 0
	for _, record := range input {
		current_min_idx := 1_000_000
		current_first := ""
		current_max_idx := -1
		current_last := ""
		for _, regex := range regex {
			numbers := find_numbers(record, regex)
			idxs := find_idxs(record, regex)
			if len(idxs) == 0 {
				continue
			}
			if idxs[0][0] < current_min_idx {
				current_min_idx = idxs[0][0]
				current_first = numbers[0]
			}
			if idxs[len(idxs)-1][0] > current_max_idx {
				current_max_idx = idxs[len(idxs)-1][0]
				current_last = numbers[0]
			}
		}
		first := convert_to_digit(current_first)
		last := convert_to_digit(current_last)
		answer += int(first*10 + last)
	}
	return answer
}

func find_numbers(input string, regex string) []string {
	re := regexp.MustCompile(regex)
	return re.FindAllString(input, -1)
}

func find_idxs(input string, regex string) [][]int {
	re := regexp.MustCompile(regex)
	return re.FindAllStringIndex(input, -1)
}

func convert_to_digit(number_str string) int {
	numbersMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	if digit_str, ok := numbersMap[number_str]; ok {
		var digit, _ = strconv.ParseInt(digit_str, 10, 64)
		return int(digit)
	}
	var number, _ = strconv.ParseInt(number_str, 10, 64)
	return int(number)

}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
