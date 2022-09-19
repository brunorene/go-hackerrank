package main

import (
	"fmt"
	"regexp"
)

/*
 * Complete the 'hackerrankInString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

var hasHackerrankMatch = regexp.MustCompile(".*h.*a.*c.*k.*e.*r.*r.*a.*n.*k.*")

func hackerrankInString(s string) string {
	if hasHackerrankMatch.FindString(s) != "" {
		return "YES"
	}

	return "NO"
}

func main() {
	result := hackerrankInString("hackerworld")

	fmt.Printf("%s\n", result)
}
