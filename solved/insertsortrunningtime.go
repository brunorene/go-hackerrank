package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func singleInsertion(n int32, arr []int32) (out []int32, shifts int32) {
	last := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		if arr[i] > last {
			arr[i+1] = arr[i]

			if i == 0 {
				arr[i] = last
			}

			shifts++

			continue
		}

		arr[i+1] = last

		break
	}

	return arr, shifts
}

func insertionOther(n int32, arr []int32) (sum int32) {
	for i := 2; i <= int(n); i++ {
		slice := arr[:i]

		sorted, shifts := singleInsertion(int32(len(slice)), slice)

		sum += shifts

		arr = append(sorted, arr[i:]...)
	}

	return sum
}

/*
 * Complete the 'runningTime' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func runningTime(arr []int32) int32 {
	return insertionOther(int32(len(arr)), arr)
}

func mainRunningTime() {
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

	result := runningTime(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}
