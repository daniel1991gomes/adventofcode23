package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"aoc23/pkg/load_input"
)

const file = "input.txt"
const test_file = "test.txt"

func main() {
	fmt.Println("-------- TEST INPUT ----------")
	run(test_file)
	fmt.Println("-------- PUZZLE INPUT --------")
	run(file)
}

func run(filename string) {
	input, err := load_input.Load(filename)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	get_answer(input)
}

func get_answer(input []string) {
	input_arr := [][]string{}
	numbers_arr := [][]string{}
	numbers_idx_arr := [][][]int{}

	stars_locs_arr := [][][]int{}

	for line_idx, line := range input {
		line_star_locs_arr := [][]int{}
		line_arr := []string{}

		// identify numbers in the line
		numbers := find_chars(line, "[0-9]+")
		numbers_arr = append(numbers_arr, numbers)
		numbers_idxs := find_idxs(line, "[0-9]+")
		numbers_idx_arr = append(numbers_idx_arr, numbers_idxs)

		// identify star chars in the line
		stars_idxs := find_idxs(line, "[*]")
		for _, star_idx := range stars_idxs {
			star_loc := [][]int{{line_idx, star_idx[0]}}
			line_star_locs_arr = append(line_star_locs_arr, star_loc...)
		}

		stars_locs_arr = append(stars_locs_arr, line_star_locs_arr)

		for _, c := range line {
			line_arr = append(line_arr, string(c))
		}

		input_arr = append(input_arr, line_arr)
	}

	// part one
	fmt.Println("---------- PART ONE ----------")
	part_one := 0
	nr_matches := 0
	for line_idx, numbers_idxs := range numbers_idx_arr {
		for n, number_idx := range numbers_idxs {

			surrounding_idxs := find_surrounding_idxs(number_idx, line_idx, len(input), len(input[0]))
			should_be_added := false
			for _, idx := range surrounding_idxs {
				line_idx_to_check := idx[0]
				char_index_to_check := idx[1]
				is_special_char := is_special_char(input_arr[line_idx_to_check][char_index_to_check])
				if is_special_char {
					should_be_added = true
					break
				}
			}

			if should_be_added {
				nr_matches += 1
				part_one += convert_to_int(numbers_arr[line_idx][n])
			}

		}
	}
	fmt.Print("Part one: matches: ")
	fmt.Print(nr_matches)
	fmt.Print(" // sum: ")
	fmt.Println(part_one)

	// part two

	fmt.Println("---------- PART TWO ----------")
	part_two := 0

	for line_idx, star_locs := range stars_locs_arr {
		if len(star_locs) == 0 {
			continue
		}
		for _, star_loc := range star_locs {
			surrounding_locs := find_surrounding_locs(star_loc, line_idx, len(input), len(input[0]))
			nearby_numbers := find_nearby_numbers(surrounding_locs, numbers_idx_arr, numbers_arr)
			if len(nearby_numbers) == 2 {
				res := nearby_numbers[0] * nearby_numbers[1]
				part_two += res

			}
		}
	}

	fmt.Print("Part two: ")
	fmt.Println(part_two)

}

func find_surrounding_locs(star_loc []int, line_idx int, max_rows int, max_cols int) [][]int {
	r := star_loc[0]
	c := star_loc[1]
	idxs := [][]int{}
	for i := -1; i <= 1; i++ {
		line_idx_to_check := r + i
		if line_idx_to_check < 0 || line_idx_to_check > max_rows-1 {
			continue
		}
		for j := -1; j <= 1; j++ {
			char_index_to_check := c + j
			if char_index_to_check < 0 || char_index_to_check > max_cols-1 {
				continue
			}
			idxs = append(idxs, []int{line_idx_to_check, char_index_to_check})
		}
	}
	return idxs
}

func find_nearby_numbers(star_locs [][]int, numbers_idx_arr [][][]int, numbers_arr [][]string) []int {
	nearby_numbers := []int{}
	uniqueMap := make(map[int]struct{})
	for _, star_loc := range star_locs {
		line_idx := star_loc[0]
		char_idx := star_loc[1]

		for i, number_idx := range numbers_idx_arr[line_idx] {
			if char_idx >= number_idx[0] && char_idx < number_idx[1] {
				number := convert_to_int(numbers_arr[line_idx][i])
				if _, ok := uniqueMap[number]; !ok {
					nearby_numbers = append(nearby_numbers, number)
					uniqueMap[number] = struct{}{}
				}
			}
		}
		// fmt.Println()
	}
	return nearby_numbers
}

func find_surrounding_idxs(number_idx []int, line_idx int, max_rows int, max_cols int) [][]int {
	starts := number_idx[0]
	ends := number_idx[1]
	idxs := [][]int{}
	for i := -1; i <= 1; i++ {
		line_idx_to_check := line_idx + i
		if line_idx_to_check < 0 || line_idx_to_check > max_rows-1 {
			continue
		}
		for j := starts - 1; j <= ends; j++ {
			char_index_to_check := j
			if char_index_to_check < 0 || char_index_to_check > max_cols-1 {
				continue
			}
			idxs = append(idxs, []int{line_idx_to_check, char_index_to_check})
		}
	}
	return idxs
}

func is_special_char(char string) bool {
	special_chars := []string{
		"#",
		"-",
		"*",
		"/",
		"@",
		"%",
		"+",
		"$",
		"=",
		"&",
	}
	for _, c := range special_chars {
		if char == c {
			return true
		}
	}
	return false
}

func convert_to_int(number_str string) int {
	number, err := strconv.Atoi(number_str)
	if err != nil {
		log.Fatalf("failed to convert string to int: %s", err)
	}
	return number
}

func find_chars(input string, regex string) []string {
	re := regexp.MustCompile(regex)
	return re.FindAllString(input, -1)
}

func find_idxs(input string, regex string) [][]int {
	re := regexp.MustCompile(regex)
	return re.FindAllStringIndex(input, -1)
}

func is_in_slice(value int, slice []int) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}
