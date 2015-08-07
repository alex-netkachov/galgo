package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/funny-string

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanLines)

	for i := 0; i < t; i++ {
		s.Scan()
		strS := s.Text()
		funny := true
		var strRB bytes.Buffer
		for j := 0; j < len(strS); j++ {
			strRB.WriteByte(strS[len(strS)-1-j])
		}
		strR := strRB.String()

		for j := 1; j < len(strS); j++ {
			funny = funny && math.Abs(float64(strS[j])-float64(strS[j-1])) == math.Abs(float64(strR[j])-float64(strR[j-1]))
		}
		if funny {
			fmt.Printf("Funny\n")
		} else {
			fmt.Printf("Not Funny\n")
		}
	}
}
