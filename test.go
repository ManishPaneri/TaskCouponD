package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the GetEmpRatingToHikeTotalNumber function below.
func GetEmpRatingToHikeTotalNumber(n int32, arr []int32) int64 {
	var i, j, k, l int32
	var t int64
	minR := make([]int32, n)
	minL := make([]int32, n)
	for i = 0; i < n; i++ {
		minR[i] = 1
		minL[i] = 1
	}

	for j = n - 2; j >= 0; j-- {

		if arr[j+1] < arr[j] {

			minR[j] = 1 + minR[j+1]
		}
	}

	for k = 1; k < n; k++ {

		if arr[k-1] < arr[k] {
			minL[k] = 1 + minL[k-1]
		}
	}

	t = 0
	for l = 0; l < n; l++ {

		if minR[l] > minL[l] {
			t = t + int64(minR[l])
		} else {
			t = t + int64(minL[l])

		}

	}

	return t
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := GetEmpRatingToHikeTotalNumber(n, arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
		fmt.Println(err)
	}
}
