package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'marsExploration' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func marsExploration(s string) int32 {
	// Write your code here
	var changes int32
	sArr := strings.Split(s, "")
	for i, v := range sArr {
		if i%3 == 0 || i%3 == 2 {
			if v != "S" {
				changes++
			}
		} else {
			if v != "O" {
				changes++
			}
		}
	}
	return changes
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	s := readLine(reader)

	result := marsExploration(s)

	//fmt.Fprintf(writer, "%d\n", result)
	fmt.Println(result)

	//writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
