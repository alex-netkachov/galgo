package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

// go fmt pairs.go && go run pairs.go < pairs.txt

type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	n := ReadInt(s)
	k := ReadInt64(s)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = ReadInt64(s)
	}
	sort.Sort(int64arr(a))

	c := 0

	i, j := 0, 0
	for {
		if j == n {
			break
		}

		if a[j]-a[i] < k {
			j++
			continue
		}

		if a[j]-a[i] == k {
			jt := j
			for jt < n && a[jt]-a[i] == k {
				jt++
			}
			jt--

			it := i
			for it < n && a[jt]-a[it] == k {
				it++
			}
			it--

			c += (jt - j + 1) * (it - i + 1)
			i = it + 1
			j = jt + 1
			continue
		}

		if a[j]-a[i] > k {
			i++
			continue
		}
	}

	Println(c)
}

func ReadInt(s *bufio.Scanner) int {
	s.Scan()
	i, _ := strconv.Atoi(s.Text())
	return i
}

func ReadInt64(s *bufio.Scanner) int64 {
	s.Scan()
	i, _ := strconv.ParseInt(s.Text(), 10, 64)
	return i
}
