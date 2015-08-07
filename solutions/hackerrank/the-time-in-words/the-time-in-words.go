package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/the-time-in-words

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	h, _ := strconv.Atoi(s.Text())

	s.Scan()
	m, _ := strconv.Atoi(s.Text())

	fmt.Printf("%s\n", timeToStr(h, m))
}

func timeToStr(h int, m int) string {
	switch {
	case m == 0:
		return fmt.Sprintf("%s o' clock", numberToStr(h))
	case m == 1:
		return fmt.Sprintf("one minute past %s", numberToStr(h))
	case m == 15:
		return fmt.Sprintf("quarter past %s", numberToStr(h))
	case 1 < m && m < 15, 15 < m && m < 30:
		return fmt.Sprintf("%s minutes past %s", numberToStr(m), numberToStr(h))
	case m == 30:
		return fmt.Sprintf("half past %s", numberToStr(h))
	case 30 < m && m < 45, 45 < m && m < 59:
		return fmt.Sprintf("%s minutes to %s", numberToStr(60-m), numberToStr(nextHour(h)))
	case m == 45:
		return fmt.Sprintf("quarter to %s", numberToStr(nextHour(h)))
	case m == 59:
		return fmt.Sprintf("one minute to %s", numberToStr(nextHour(h)))
	}
	return ""
}

func numberToStr(n int) string {
	switch {
	case n == 1:
		return "one"
	case n == 2:
		return "two"
	case n == 3:
		return "three"
	case n == 4:
		return "four"
	case n == 5:
		return "five"
	case n == 6:
		return "six"
	case n == 7:
		return "seven"
	case n == 8:
		return "eight"
	case n == 9:
		return "nine"
	case n == 10:
		return "ten"
	case n == 11:
		return "eleven"
	case n == 12:
		return "twelve"
	case n == 13:
		return "thirteen"
	case n == 15:
		return "fifteen"
	case n == 14, n == 16, n == 17, n == 19:
		return fmt.Sprintf("%steen", numberToStr(n-10))
	case n == 18:
		return "eighteen"
	case n == 20:
		return "twenty"
	case 20 < n && n < 30:
		return fmt.Sprintf("twenty %s", numberToStr(n-20))
	case n == 30:
		return "thirty"
	case 30 < n && n < 40:
		return fmt.Sprintf("thirty %s", numberToStr(n-30))
	case n == 40:
		return "fourty"
	case 40 < n && n < 50:
		return fmt.Sprintf("fourty %s", numberToStr(n-40))
	case n == 50:
		return "fifty"
	case 50 < n && n < 60:
		return fmt.Sprintf("fifty %s", numberToStr(n-50))
	}
	return ""
}

func nextHour(n int) int {
	if n == 12 {
		return 1
	}
	return n + 1
}
