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
 * Complete the 'flippingBits' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER n as parameter.
 */

func flippingBits(n int64) int64 {
	// Write your code here
	binary := strconv.FormatInt(n, 2)
	binarySlice := strings.Split(binary, "")
	leadingZeros := make([]string, 32-len(binarySlice))
	for i := range leadingZeros {
		leadingZeros[i] = "0"
	}
	fullBinary := append(leadingZeros, binarySlice...)
	for i, v := range fullBinary {
		if v == "0" {
			fullBinary[i] = "1"
		} else {
			fullBinary[i] = "0"
		}
	}
	newBinary := strings.Join(fullBinary, "")
	newInt, err := strconv.ParseInt(newBinary, 2, 64)
	checkError(err)
	return newInt
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		result := flippingBits(n)

		//fmt.Fprintf(writer, "%d\n", result)
		fmt.Printf("%d\n", result)
	}

	//writer.Flush()
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
