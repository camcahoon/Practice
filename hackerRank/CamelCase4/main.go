package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func camelCase(arr []string) []string {
	var returnStrings []string
	for _, v := range arr {
		sections := strings.Split(v, ";")
		if sections[0] == "S" {
			returnStrings = append(returnStrings, split(sections[2], sections[1]))
		} else {
			returnStrings = append(returnStrings, combine(sections[2], sections[1]))
		}
	}
	return returnStrings
}

func split(str string, version string) string {
	separators := make([]string, 0)
	words := make([]string, 0)
	for _, v := range str {
		if unicode.IsUpper(v) {
			separators = append(separators, string(v))
		}
	}
	for _, v := range separators {
		splitWords := strings.SplitN(str, v, 2)
		words = append(words, strings.ToLower(splitWords[0]))
		splitWords[0] = v
		str = strings.Join(splitWords, "")
	}
	words = append(words, strings.ToLower(str))
	if version == "M" {
		//before := strings.Split(words[len(words)-1], "(")
		//words[len(words)-1] = strings.ToLower(before[0])
		before, _, _ := strings.Cut(words[len(words)-1], "(")
		words[len(words)-1] = strings.ToLower(before)
	}
	if version == "C" {
		words = words[1:]
	}
	return strings.Join(words, " ")
}

func combine(str string, version string) string {
	strSlice := strings.Split(str, " ")
	c := []cases.Caser{
		cases.Title(language.English),
	}
	if version == "C" {
		for i, v := range strSlice {
			//strSlice[i] = strings.ToTitle(v)
			strSlice[i] = c[0].String(v)
		}
	} else {
		for i := 1; i < len(strSlice); i++ {
			//strSlice[i] = strings.ToTitle(strSlice[i])
			strSlice[i] = c[0].String(strSlice[i])
		}
	}
	if version == "M" {
		strSlice = append(strSlice, "()")
	}
	return strings.Join(strSlice, "")
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var stringsSlice []string

	//for {
	//    stringsItem := readLine(reader)
	//    if stringsItem == "" {
	//        break
	//    }
	//    stringsSlice = append(stringsSlice, stringsItem)
	//}
	for i := 0; i < 6; i++ {
		stringsItem := readLine(reader)
		stringsSlice = append(stringsSlice, stringsItem)
	}

	res := camelCase(stringsSlice)

	for i, resItem := range res {
		//fmt.Fprintf(writer, "%s", resItem)
		fmt.Printf("%s", resItem)

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

//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
