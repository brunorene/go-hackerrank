package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var (
	numbers      = regexp.MustCompile(`[0-9]`)
	lowercase    = regexp.MustCompile(`[a-z]`)
	uppercase    = regexp.MustCompile(`[A-Z]`)
	specialChars = regexp.MustCompile(`[!@#$%^&*()\-+]`)
)

func minimumNumber(n int32, password string) int32 {
	more := int32(0)

	if numbers.FindStringSubmatch(password) == nil {
		fmt.Println("numbers")
		more++
	}

	if lowercase.FindStringSubmatch(password) == nil {
		fmt.Println("lower")
		more++
	}

	if uppercase.FindStringSubmatch(password) == nil {
		fmt.Println("upper")
		more++
	}

	if specialChars.FindStringSubmatch(password) == nil {
		fmt.Println("special")
		more++
	}

	total := more + n

	fmt.Println(total)

	if total >= 6 {
		return more
	}

	return more + (6 - total)
}

func main() {
	answer := minimumNumber(7, "AUzs-nV")

	fmt.Println(answer)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
