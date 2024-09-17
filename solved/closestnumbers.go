package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'closestNumbers' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func closestNumbers(arr []int32) []int32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	var result []int32

	minDist := int32(math.MaxInt32)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == minDist {
			result = append(result, arr[i], arr[i+1])

			continue
		}

		if arr[i+1]-arr[i] < minDist {
			result = nil

			result = append(result, arr[i], arr[i+1])

			minDist = arr[i+1] - arr[i]

			continue
		}
	}

	return result
}

func mainCN() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := closestNumbers(arr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}
