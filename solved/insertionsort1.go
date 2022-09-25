package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func print(arr []int32) {
	out := fmt.Sprint(arr)
	fmt.Println(out[1 : len(out)-1])
}

/*
 * Complete the 'insertionSort1' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY arr
 */

func insertionSort1(n int32, arr []int32) {
	last := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		if arr[i] > last {
			arr[i+1] = arr[i]
			print(arr)

			if i == 0 {
				arr[i] = last
				print(arr)
			}

			continue
		}

		arr[i+1] = last
		print(arr)

		break
	}
}

func mainIns1() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

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

	insertionSort1(n, arr)
}
