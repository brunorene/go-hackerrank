package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var sameChars = regexp.MustCompile(`aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz`)

/*
 * Complete the 'superReducedString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func superReducedString(s string) string {
	for sameChars.FindString(s) != "" {
		s = sameChars.ReplaceAllString(s, "")
	}

	if s == "" {
		return "Empty String"
	}

	return s
}

func mainSuper() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := superReducedString(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}
