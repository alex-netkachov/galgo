package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt countingsort4.go && go run countingsort4.go < countingsort4.txt

func main() {
	var d []uint8
	r := bufio.NewReader(os.Stdin)

	d, _ = r.ReadBytes('\n')
	n, _ := strconv.Atoi(string(d[0 : len(d)-1]))

	m := make([][]string, 100)
	for i := 0; i < 100; i++ {
		m[i] = []string{}
	}

	for i := 0; i < n; i++ {
		var v int
		var str string

		d, _ = r.ReadBytes('\n')
		for j := 0; j < len(d); j++ {
			tr := 0
			if d[len(d)-1] == uint8('\n') {
				tr = 1
			}
			if d[j] == uint8(' ') {
				v, _ = strconv.Atoi(string(d[0:j]))
				if i < n/2 {
					str = "-"
				} else {
					str = string(d[j+1 : len(d)-tr])
				}
				break
			}
		}

		m[v] = append(m[v], str)
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, a := range m {
		for _, w := range a {
			fmt.Fprint(wr, w)
			fmt.Fprint(wr, " ")
		}
	}
	fmt.Fprint(wr, "\n")
	wr.Flush()
}
