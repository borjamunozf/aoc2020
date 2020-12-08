package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {

	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func part1(arr []string, hm map[int]int, target int) {
	// Create hashmap
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		e, err := strconv.Atoi(arr[i])
		check(err)
		hm[target-e] = e
	}
	//Iterate inputs:
	for i := 0; i < len(arr); i++ {
		e, err := strconv.Atoi(arr[i])
		check(err)
		if _, ok := hm[e]; ok {
			fmt.Printf("Found %s\n", arr[i])
			fmt.Println("2020 is sum of:", hm[e], e)
			fmt.Println("Multiplication: ", hm[e]*e)
			break
		}
	}
}

func part2(arr []string, hm map[int]int, target int) {

	// Create hashmap
	for i := 0; i < len(arr)-2; i++ {
		e, err := strconv.Atoi(arr[i])
		check(err)
		hm[e] = e
		for j := i + 1; j < len(arr); j++ {
			e2, err := strconv.Atoi(arr[j])
			check(err)
			hm[e2] = e2
			remainder := target - e - e2
			if _, ok := hm[remainder]; ok {
				fmt.Println("found", e, e2, remainder)
				fmt.Println(e + e2 + remainder)
				fmt.Println("multiply:", e*e2*remainder)
				break
			}
		}
	}

}

func main() {

	fmt.Println("Advent of Code: no. 1")

	//Read file
	input, err := ioutil.ReadFile("input")
	check(err)

	// Split lines in array
	arr := strings.Split(string(input), "\n")
	var hm map[int]int

	hm = make(map[int]int)

	// Target number 2020
	target := 2020

	part1(arr, hm, target)
	part2(arr, hm, target)
}
