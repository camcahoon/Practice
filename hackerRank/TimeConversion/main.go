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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	strSlice := strings.Split(s, ":")
	sec := strSlice[2]
	secSlice := strings.Split(sec, "")
	if secSlice[2] == "P" {
		if strSlice[0] != "12" {
			hour, err := strconv.ParseInt(strSlice[0], 10, 64)
			checkError(err)
			strSlice[0] = strconv.FormatInt(hour+12, 10)
		}
	} else {
		if strSlice[0] == "12" {
			strSlice[0] = "00"
		}
	}
	secSlice = secSlice[:2]
	sec = strings.Join(secSlice, "")
	strSlice[2] = sec
	return strings.Join(strSlice, ":")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
		panic(err)
	}
}
