package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"

	"aoc23/pkg/load_input"
)

const file = "input.txt"
const test_file = "test.txt"

type Card struct {
	count   int
	matches int
}

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

	// part one
	part_one := 0
	list_of_cards := []Card{}

	for _, line := range input {
		match_count, score := score_card(line)
		part_one += score

		// part two
		list_of_cards = append(list_of_cards, Card{count: 1, matches: match_count})
	}

	fmt.Println("Part one:", part_one)

	// part two
	part_two := 0
	for c, card := range list_of_cards {
		for i := 1; i <= card.matches; i++ {
			if c+i >= len(list_of_cards) {
				break
			}
			next_card := list_of_cards[c+i]
			new_next_card := Card{count: next_card.count + card.count, matches: next_card.matches}
			list_of_cards[c+i] = new_next_card
		}
		part_two += card.count
	}
	fmt.Println("part two:", part_two)

}

func score_card(line string) (int, int) {
	card := strings.Split(line, ": ")[1]
	all_numbers := strings.Split(card, " | ")
	winning_numbers := all_numbers[0]
	our_numbers := all_numbers[1]

	match_count := 0

	for _, number := range strings.Split(our_numbers, " ") {
		if number == "" {
			continue
		}
		is_winning := is_winning_number(winning_numbers, number)
		// fmt.Print("Number: ")
		// fmt.Print(number)
		// fmt.Print(" Winning: ")
		// fmt.Println(is_winning)
		if is_winning {
			match_count += 1
		}
	}
	score := int(math.Pow(2, float64(match_count-1)))
	return match_count, score
}

func is_winning_number(winning_numbers string, number string) bool {
	re := regexp.MustCompile("\\b(?:" + number + ")\\b")
	match := re.FindString(winning_numbers)
	if match != "" {
		return true
	}
	return false
}
