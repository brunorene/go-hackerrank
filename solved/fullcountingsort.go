package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'countSort' function below.
 *
 * The function accepts 2D_STRING_ARRAY arr as parameter.
 */

func countSort(arr [][]string) {
	maxNum := 100

	counter := make(map[int][]string, maxNum)

	for idx, tuple := range arr {
		num, err := strconv.Atoi(tuple[0])
		checkError(err)

		if (idx / (len(arr) / 2)) == 0 {
			counter[num] = append(counter[num], "-")
		} else {
			counter[num] = append(counter[num], tuple[1])
		}
	}

	result := make([]string, 0, len(counter))

	for idx := range maxNum {
		vals, exists := counter[idx]

		if !exists {
			continue
		}

		result = append(result, strings.Join(vals, " "))
	}

	fmt.Println(strings.Join(result, " "))
}

func mainFCS() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]string
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []string
		for _, arrRowItem := range arrRowTemp {
			arrRow = append(arrRow, arrRowItem)
		}

		if len(arrRow) != 2 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	countSort(arr)
}
