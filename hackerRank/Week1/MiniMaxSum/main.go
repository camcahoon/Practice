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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var maxSum int
	var minSum int
	sums := make([]int, 0)
	for i, val1 := range arr {
		for j, val2 := range arr {
			for k, val3 := range arr {
				for l, val4 := range arr {
					if i != j && i != k && i != l && j != k && j != l && k != l {
						sum := int(val1) + int(val2) + int(val3) + int(val4)
						sums = append(sums, sum)
					}
				}
			}
		}
	}
	maxSum = sums[0]
	minSum = sums[0]
	for _, value := range sums {
		if value > maxSum {
			maxSum = value
		}
	}
	for _, value := range sums {
		if value < minSum {
			minSum = value
		}
	}
	fmt.Println(minSum, maxSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
