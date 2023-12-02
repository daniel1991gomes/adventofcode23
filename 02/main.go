package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc23/pkg/load_input"
)

const file = "input.txt"
const test_file = "test.txt"

func main() {
	run(test_file)
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
	dice_count := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	part_one_answer := 0
	part_two_answer := 0

	for _, line := range input {
		game_rounds_split := strings.Split(line, ": ")
		game_str, rounds_str := game_rounds_split[0], game_rounds_split[1]
		game_id, err := strconv.Atoi(strings.Split(game_str, " ")[1])
		if err != nil {
			log.Fatalf("failed to convert game to int: %s", err)
		}
		rounds := strings.Split(rounds_str, "; ")
		all_dices_valid := true

		min_dice_cont := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, round := range rounds {
			dices := strings.Split(round, ", ")
			for _, dice := range dices {
				number, err := strconv.Atoi(strings.Split(dice, " ")[0])
				if err != nil {
					log.Fatalf("failed to convert dice to int: %s", err)
				}
				color := strings.Split(dice, " ")[1]
				if number > dice_count[color] {
					all_dices_valid = false
				}
				if number > min_dice_cont[color] {
					min_dice_cont[color] = number
				}
			}
		}
		if all_dices_valid {
			part_one_answer += game_id
		}
		power := min_dice_cont["red"] * min_dice_cont["green"] * min_dice_cont["blue"]
		part_two_answer += power
	}
	fmt.Println("####################")
	fmt.Print("Part one: ")
	fmt.Println(part_one_answer)
	fmt.Print("Part two: ")
	fmt.Println(part_two_answer)
	fmt.Println("####################")
}
