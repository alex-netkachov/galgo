package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/library-fine

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	da, _ := strconv.Atoi(s.Text())
	s.Scan()
	ma, _ := strconv.Atoi(s.Text())
	s.Scan()
	ya, _ := strconv.Atoi(s.Text())

	s.Scan()
	de, _ := strconv.Atoi(s.Text())
	s.Scan()
	me, _ := strconv.Atoi(s.Text())
	s.Scan()
	ye, _ := strconv.Atoi(s.Text())

	fine := 0
	if ya > ye {
		fine = 10000
	} else if ya == ye {
		if ma > me {
			fine = (ma - me) * 500
		} else if ma == me {
			if da > de {
				fine = (da - de) * 15
			}
		}
	}
	fmt.Printf("%v\n", fine)
}
