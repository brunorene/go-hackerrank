package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'insertionSort2' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY arr
 */

func print(arr []int32) {
	out := fmt.Sprint(arr)
	fmt.Println(out[1 : len(out)-1])
}

func insertion(n int32, arr []int32) []int32 {
	last := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		if arr[i] > last {
			arr[i+1] = arr[i]

			if i == 0 {
				arr[i] = last
			}

			continue
		}

		arr[i+1] = last

		break
	}

	return arr
}

func insertionSort2(n int32, arr []int32) {
	for i := 2; i <= int(n); i++ {
		slice := arr[:i]

		sorted := insertion(int32(len(slice)), slice)

		arr = append(sorted, arr[i:]...)

		print(arr)
	}
}

func main() {
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

	insertionSort2(n, arr)
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
