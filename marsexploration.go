package main

import (
	"fmt"
)

/*
 * Complete the 'marsExploration' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func marsExploration(s string) int32 {
	mistakes := 0

	for idx, c := range s {
		if (idx%3 == 1 && c != 'O') || (idx%3 != 1 && c != 'S') {
			mistakes++
		}
		fmt.Printf("%c %d\n", c, mistakes)
	}

	return int32(mistakes)
}

func main() {
	fmt.Println(marsExploration("SOSSOS"))
}
