// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println(ParsePhone("     700-128-08-07"))
	fmt.Println(Anagram("abcd", "cdab"))
	nums := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(FindEvens(nums))
	fmt.Println(SliceProduct(nums))
	fmt.Println(Unique(nums))
	cars := map[string]int{
		"audi":   5,
		"bmw":    3,
		"merc":   8,
		"toyota": 12,
	}
	fmt.Println(InvertMap(cars))
	fmt.Println(TopCharacters("1112223334444", 3))

}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	phone = strings.Replace(phone, " ", "", -1)
	phone = strings.Replace(phone, "-", "", -1)
	var result = "("
	for pos, char := range phone {
		result = result + string(char)
		if pos == 2 {
			result = result + ") "
		}
		if pos == 5 {
			result = result + "-"
		}
	}
	return result
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	var result = s2
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {

				result = strings.Replace(result, string(s2[j]), "", -1)
			}
		}
	}
	if result == "" {
		return true
	} else {
		return false
	}
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	result := []int{}
	for _, e := range e {
		if e%2 == 0 {
			result = append(result, e)
		}
	}
	return result
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	var result = 0
	for _, e := range e {
		result = result + e
	}
	return result
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	result := []int{}
	for _, e := range e {
		if !contains(result, e) {
			result = append(result, e)
		}
	}
	return result
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	// TODO
	m := make(map[int]string)
	for k, v := range kv {
		m[v] = k
	}
	return m
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	// TODO
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	for r, cnt := range counts {
		if cnt <= k {
			delete(counts, r)
		}
	}
	return counts
}
