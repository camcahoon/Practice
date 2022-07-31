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
 * Complete the 'twoArrays' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY A
 *  3. INTEGER_ARRAY B
 */

func twoArrays(k int32, A []int32, B []int32) string {
	// Write your code here
	/*
		    permutationsA := permute(A)
			permutationsB := permute(B)
			var possible bool
			for _, a := range permutationsA {
				for _, b := range permutationsB {
					for i := range a {
						if a[i]+b[i] >= k {
							possible = true
						} else {
							possible = false
							break
						}
					}
					if possible {
						break
					}
				}
				if possible {
					break
				}
			}
			if possible {
				return "YES"
			} else {
				return "NO"
			}
	*/
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})
	for i := range A {
		if A[i]+B[i] < k {
			return "NO"
		}
	}
	return "YES"
}

/*
func permute(arr []int32) [][]int32 {
	permutations := make([][]int32, 0)
	c := make([]int32, len(arr))
	permutations = append(permutations, arr)

	i := 1
	for i < len(arr) {
		if c[i] < int32(i) {
			tempArr := make([]int32, len(arr))
			copy(tempArr, arr)
			if i%2 == 0 {
				tempArr[0], tempArr[i] = tempArr[i], tempArr[0]
			} else {
				tempArr[c[i]], tempArr[i] = tempArr[i], tempArr[c[i]]
			}
			var in bool
			for _, v := range permutations {
				for i, v2 := range v {
					if v2 == tempArr[i] {
						in = true
					} else {
						in = false
						break
					}
				}
				if in {
					break
				}
			}
			if !in {
				permutations = append(permutations, tempArr)
			}
			c[i]++
			i = 1
			arr = tempArr
			fmt.Println(arr)
		} else {
			c[i] = 0
			i++
		}
	}
	return permutations
}
*/

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
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		k := int32(kTemp)

		ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var A []int32

		for i := 0; i < int(n); i++ {
			AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			A = append(A, AItem)
		}

		BTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var B []int32

		for i := 0; i < int(n); i++ {
			BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
			checkError(err)
			BItem := int32(BItemTemp)
			B = append(B, BItem)
		}

		result := twoArrays(k, A, B)

		//fmt.Fprintf(writer, "%s\n", result)
		fmt.Println(result)
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
