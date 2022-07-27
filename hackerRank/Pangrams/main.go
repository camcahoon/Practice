package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(str string) string {
	// Write your code here
	var a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z int
	strArr := strings.Split(str, "")
	for _, letter := range strArr {
		switch {
		case letter == "a" || letter == "A":
			a++
		case letter == "b" || letter == "B":
			b++
		case letter == "c" || letter == "C":
			c++
		case letter == "d" || letter == "D":
			d++
		case letter == "e" || letter == "E":
			e++
		case letter == "f" || letter == "F":
			f++
		case letter == "g" || letter == "G":
			g++
		case letter == "h" || letter == "H":
			h++
		case letter == "i" || letter == "I":
			i++
		case letter == "j" || letter == "J":
			j++
		case letter == "k" || letter == "K":
			k++
		case letter == "l" || letter == "L":
			l++
		case letter == "m" || letter == "M":
			m++
		case letter == "n" || letter == "N":
			n++
		case letter == "o" || letter == "O":
			o++
		case letter == "p" || letter == "P":
			p++
		case letter == "q" || letter == "Q":
			q++
		case letter == "r" || letter == "R":
			r++
		case letter == "s" || letter == "S":
			s++
		case letter == "t" || letter == "T":
			t++
		case letter == "u" || letter == "U":
			u++
		case letter == "v" || letter == "V":
			v++
		case letter == "w" || letter == "W":
			w++
		case letter == "x" || letter == "X":
			x++
		case letter == "y" || letter == "Y":
			y++
		case letter == "z" || letter == "Z":
			z++
		}
	}
	switch {
	case a <= 0:
		return "not pangram"
	case b <= 0:
		return "not pangram"
	case c <= 0:
		return "not pangram"
	case d <= 0:
		return "not pangram"
	case e <= 0:
		return "not pangram"
	case f <= 0:
		return "not pangram"
	case g <= 0:
		return "not pangram"
	case h <= 0:
		return "not pangram"
	case i <= 0:
		return "not pangram"
	case j <= 0:
		return "not pangram"
	case k <= 0:
		return "not pangram"
	case l <= 0:
		return "not pangram"
	case m <= 0:
		return "not pangram"
	case n <= 0:
		return "not pangram"
	case o <= 0:
		return "not pangram"
	case p <= 0:
		return "not pangram"
	case q <= 0:
		return "not pangram"
	case r <= 0:
		return "not pangram"
	case s <= 0:
		return "not pangram"
	case t <= 0:
		return "not pangram"
	case u <= 0:
		return "not pangram"
	case v <= 0:
		return "not pangram"
	case w <= 0:
		return "not pangram"
	case x <= 0:
		return "not pangram"
	case y <= 0:
		return "not pangram"
	case z <= 0:
		return "not pangram"
	default:
		return "pangram"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := pangrams(s)

	//fmt.Fprintf(writer, "%s\n", result)
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
