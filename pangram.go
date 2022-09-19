package main

import (
	"fmt"
	"strings"
)

func pangrams(s string) string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	upper := strings.ToUpper(s)

	for _, c := range upper {
		letters = strings.ReplaceAll(letters, string(c), "")
	}

	if letters == "" {
		return "pangram"
	}

	return "not pangram"
}

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the next prize"))
}
