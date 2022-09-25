package main

import (
	"fmt"
	"sort"
)

type pair [2]rune

func uniqueChars(s string) (result string) {
	chars := make(map[rune]struct{})

	for _, c := range s {
		chars[c] = struct{}{}
	}

	for k := range chars {
		result += string(k)
	}

	return
}

func removeOthers(s string, p pair) string {
	var result string

	for _, c := range s {
		if c == p[0] || c == p[1] {
			result += string(c)
		}
	}

	return result
}

func combinations(c string) (pairs []pair) {
	pairMap := make(map[string]struct{})

	for _, c1 := range c {
		for _, c2 := range c {
			if c1 != c2 {
				pair := []rune{c1, c2}
				sort.Slice(pair, func(i, j int) bool {
					return pair[i] < pair[j]
				})
				pairMap[string(pair)] = struct{}{}
			}
		}
	}

	for k := range pairMap {
		pairs = append(pairs, pair{rune(k[0]), rune(k[1])})
	}

	return
}

func valid(s string, p pair) bool {
	index := 0

	if rune(s[0]) == p[1] {
		index = 1
	}

	for _, c := range s {
		if c != p[index] {
			return false
		}

		index = (index + 1) % 2
	}

	return true
}

func alternate(s string) int32 {
	combs := combinations(uniqueChars(s))

	var max int

	for _, k := range combs {
		candidate := removeOthers(s, k)

		if valid(candidate, k) && len(candidate) > max {
			max = len(candidate)
		}
	}

	return int32(max)
}

func main2Chars() {
	fmt.Println(alternate("beabeefeab"))
}
