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
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 */

func matchingStrings(stringsSlice []string, queries []string) []int32 {
	// Write your code here
	matches := make([]int32, len(queries))
	for i, val1 := range queries {
		for _, val2 := range stringsSlice {
			if val1 == val2 {
				matches[i]++
			}
		}
	}
	return matches
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stringsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var stringsSlice []string

	for i := 0; i < int(stringsCount); i++ {
		stringsItem := readLine(reader)
		stringsSlice = append(stringsSlice, stringsItem)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []string

	for i := 0; i < int(queriesCount); i++ {
		queriesItem := readLine(reader)
		queries = append(queries, queriesItem)
	}

	res := matchingStrings(stringsSlice, queries)

	for i, resItem := range res {
		//fmt.Fprintf(writer, "%d", resItem)
		fmt.Printf("%d", resItem)

		if i != len(res)-1 {
			//fmt.Fprintf(writer, "\n")
			fmt.Print("\n")
		}
	}

	//fmt.Fprintf(writer, "\n")
	fmt.Print("\n")

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
