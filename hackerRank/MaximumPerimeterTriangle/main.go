package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'maximumPerimeterTriangle' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY sticks as parameter.
 */

func maximumPerimeterTriangle(sticks []int32) []int32 {
	// Write your code here
	possibleTriangles := make([][]int32, 0)
	var largestPerimeter int
	var lpTriangle []int32
	for i, v1 := range sticks {
		for j, v2 := range sticks {
			for k, v3 := range sticks {
				if i != j && i != k && j != k {
					possibleTriangles = append(possibleTriangles, []int32{v1, v2, v3})
				}
			}
		}
	}
	for i, v := range possibleTriangles {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		if v[0]+v[1] <= v[2] {
			possibleTriangles[i] = []int32{-1}
		}
	}
	for _, triangle := range possibleTriangles {
		var sum int
		for _, side := range triangle {
			sum += int(side)
		}
		if sum > largestPerimeter {
			largestPerimeter = sum
			lpTriangle = triangle
		} else if sum == largestPerimeter {
			if triangle[2] > lpTriangle[2] {
				lpTriangle = triangle
			} else if triangle[2] == lpTriangle[2] {
				if triangle[0] > lpTriangle[0] {
					lpTriangle = triangle
				}
			}
		}
	}
	if largestPerimeter > 0 {
		return lpTriangle
	} else {
		return []int32{-1}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	sticksTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var sticks []int32

	for i := 0; i < int(n); i++ {
		sticksItemTemp, err := strconv.ParseInt(sticksTemp[i], 10, 64)
		checkError(err)
		sticksItem := int32(sticksItemTemp)
		sticks = append(sticks, sticksItem)
	}

	result := maximumPerimeterTriangle(sticks)

	for i, resultItem := range result {
		//fmt.Fprintf(writer, "%d", resultItem)
		fmt.Print(resultItem)

		if i != len(result)-1 {
			//fmt.Fprintf(writer, " ")
			fmt.Print(" ")
		}
	}

	//fmt.Fprintf(writer, "\n")
	fmt.Println()

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
