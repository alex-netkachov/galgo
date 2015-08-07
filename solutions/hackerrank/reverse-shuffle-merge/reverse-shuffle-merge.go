package main

import (
	"bytes"
	. "fmt"
	"math/rand"
	"time"
)

// https://www.hackerrank.com/challenges/reverse-shuffle-merge

const z byte = 122
const a byte = 97

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	s := ""

	// test("aa", "a")
	// test("eggegg", "egg")
	// test("bdabaceadaedaaaeaecdeadababdbeaeeacacaba", "aaaaaabaaceededecbdb")
	// test("djjcddjggbiigjhfghehhbgdigjicafgjcehhfgifadihiajgciagicdahcbajjbhifjiaajigdgdfhdiijjgaiejgegbbiigida", "aaaaabccigicgjihidfiejfijgidgbhhehgfhjgiibggjddjjd")
	// s = "aaaaaaaaaaaaaabbbbbbbbbbbbcccccccccccccddddddddddddddeeeeeeeeeeeeeezzzzzzzzzzz"
	// test(string(merge(rev([]byte(s)), shuffle([]byte(s)))), s)

	s = ""
	Scanln(&s)
	d := []byte(s)
	Println(string(find(d)))

	//s := "aaaaaaaaaaaaaabbbbbbbbbbbbcccccccccccccddddddddddddddeeeeeeeeeeeeeezzzzzzzzzzz"
	//test(string(merge(rev([]byte(s)), shuffle([]byte(s)))), s)

	// for i := 0; i < 10000; i++ {
	//  s := "aaaaaaaaaaaaaabbbbbbbbbbbbcccccccccccccddddddddddddddeeeeeeeeeeeeeezzzzzzzzzzz"
	//  d := merge(rev([]byte(s)), shuffle([]byte(s)))
	//  r := string(find(d))
	//  if r != s {
	//    Println(string(d), ">", r)
	//  }
	// }
}

func test(s, r string) {
	//Println(s)
	a := string(find([]byte(s)))
	//Println(a, r)
	if a == r {
		Println("> PASS")
	} else {
		Println(s)
		Println(a)
		Println(r)
		Println("> FAIL")

	}
}

func merge(b1 []byte, b2 []byte) []byte {
	r := make([]byte, len(b1)+len(b2))
	i := []int{0, 0}
	b := [][]byte{b1, b2}
	j := 0
	for {
		s := rand.Intn(2)
		if i[s] == len(b[s]) {
			s = (s + 1) % 2
		}
		if i[s] == len(b[s]) {
			break
		}
		r[j] = b[s][i[s]]
		i[s]++
		j++
	}
	return r
}

func shuffle(b []byte) []byte {
	i := rand.Perm(len(b))
	r := make([]byte, len(b))
	for k, v := range b {
		r[i[k]] = v
	}
	return r
}

func rev(a []byte) []byte {
	r := make([]byte, len(a))
	for i, j := 0, len(a)-1; i < len(a); i, j = i+1, j-1 {
		r[i] = a[j]
	}
	return r
}

type selection struct {
	t byte
	m int
	v []int
	s []int
	r []int
}

func makeSelection(t byte, m int, v []int, n int) selection {
	s := selection{t, m, v, make([]int, n), make([]int, n)}
	s.min()
	return s
}

func (this *selection) val(i int) int {
	return this.v[this.s[i]]
}

func (this *selection) save() {
	copy(this.r, this.s)
}

func (this *selection) load() {
	copy(this.s, this.r)
}

func (this *selection) next() bool {
	// Example:
	// [ 0, 1, 2, 3, 4]
	//   x  x  x        0 1 2
	//   x  x     x     0 1 3
	//   x     x  x     0 2 3
	//      x  x  x     1 2 3
	//   x  x        x  0 1 4
	//   x     x     x  0 2 4
	//      x  x     x  1 2 4
	//   x        x  x  0 3 4
	//      x     x  x  1 3 4
	//         x  x  x  2 3 4

	// find the node that can be moved
	s := this.s
	last := len(s) - 1
	max := len(this.v) - 1
	j := -1

	// everyhting except last
	for i := 0; i < last; i++ {
		if s[i+1]-s[i] == 1 {
			continue
		}
		j = i
		break
	}

	// last
	if j == -1 {
		if s[last] == max {
			return false
		} else {
			j = last
		}
	}

	// build the next permutation by increasing the selected and
	// resetting the items before the selected
	s[j]++
	for i := 0; i < j; i++ {
		s[i] = i
	}
	return true
}

func (this *selection) prev() bool {
	// inverse to next
	s := this.s
	first := 0
	min := 0
	j := -1

	// everyhting except first
	for i := len(s) - 1; i > first; i-- {
		if s[i-1]-s[i] == -1 {
			continue
		}
		j = i
	}

	// first
	if j == -1 {
		if s[first] == min {
			return false
		} else {
			j = first
		}
	}

	// build the prev permutation by decreasing the selected and
	// maximizing the items after the selected
	s[j]--
	for i, k := j-1, s[j]-1; i > -1; i, k = i-1, k-1 {
		s[i] = k
	}
	return true
}

func (this *selection) min() {
	for i := 0; i < len(this.s); i++ {
		this.s[i] = i
	}
}

func (this *selection) max() {
	for i := 0; i < len(this.s); i++ {
		this.s[len(this.s)-1-i] = len(this.v) - 1 - i
	}
}

func (this *selection) shift() bool {
	if this.s[len(this.s)-1] == len(this.v)-1 {
		return false
	}
	for i := 0; i < len(this.s); i++ {
		this.s[i]++
	}
	return true
}

func (this *selection) shiftFromIndex(i int) bool {
	if this.s[len(this.s)-1] == len(this.v)-1 {
		return false
	}
	for ; i < len(this.s); i++ {
		this.s[i]++
	}
	return true
}

func (this *selection) unshift() bool {
	if this.s[0] == 0 {
		return false
	}
	for i := 0; i < len(this.s); i++ {
		this.s[i]--
	}
	return true
}

func (this *selection) unshiftIndex(i int) bool {
	if i == 0 {
		if this.s[0] == 0 {
			return false
		}
	} else {
		if this.s[i] == this.s[i-1]+1 {
			return false
		}
	}
	this.s[i]--
	return true
}

func (this *selection) mask() []byte {
	m := make([]byte, this.m)
	this.apply(m)
	return m
}

func (this *selection) apply(m []byte) {
	v := this.v
	for _, i := range this.s {
		m[v[i]] = this.t
	}
}

func (this *selection) print() {
	s1 := make([]byte, this.m)
	s2 := make([]byte, this.m)
	for i := 0; i < this.m; i++ {
		s1[i] = byte('.')
		s2[i] = byte(' ')
	}
	for i := 0; i < len(this.v); i++ {
		s1[this.v[i]] = byte('*')
	}
	for i := 0; i < len(this.s); i++ {
		s2[this.v[this.s[i]]] = byte('^')
	}
	Println(string(this.t), ":", string(s1))
	Println("  :", string(s2))
}

func printm(m []*selection) {
	for _, s := range m {
		s.print()
	}
}

func find(t []byte) []byte {
	d := rev(t)
	m := make([]*selection, 0)
	for q := a; q <= z; q++ {
		v := make([]int, 0)
		for i := 0; i < len(d); i++ {
			if d[i] == q {
				v = append(v, i)
			}
		}
		if len(v) > 0 {
			s := makeSelection(q, len(d), v, len(v)/2)
			m = append(m, &s)
			//Println("For", string(q))
			//s.print()
		}
	}

	mz := make([]byte, len(d))
	ms := make([]byte, len(d))
	getrow := func() []byte {
		copy(ms, mz)
		for _, mi := range m {
			mi.apply(ms)
		}
		s := make([]byte, len(d)/2)
		j := 0
		for _, v := range ms {
			if v != 0 {
				s[j] = v
				j++
			}
		}
		return s
	}

	//Println(string(d))
	//printm(m)

	// brutforce
	// min := getrow()
	// var rec func(i int)
	// rec = func(i int) {
	//  if i < len(m) {
	//    do := true
	//    for do {
	//      rec(i + 1)
	//      do = m[i].next()
	//    }
	//    m[i].min()
	//  } else {
	//    s := getrow()
	//    if bytes.Compare(min, s) == 1 {
	//      min = s
	//    }
	//  }
	// }
	// rec(0)
	// return min

	// topological steps
	min := getrow()
	for {
		moved := false
		if bytes.Compare(min, getrow()) == 0 {
			moved = false
		}
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i].s); j++ {
				p := m[i].val(j)
				for k := len(m) - 1; k > i; k-- {
					//print("\033[H\033[2J")
					//Println("printing....")
					for n := 0; n < len(m[k].s); n++ {
						for m[k].val(n) <= p {
							m[k].save()
							if m[k].shiftFromIndex(n) {
								s2 := getrow()
								if bytes.Compare(min, s2) == -1 {
									m[k].load()
									break
								}
								min = s2
							} else {
								break
							}
						}
					}
					//printm(m)
					//time.Sleep(50 * time.Millisecond)
				}
			}
		}

		min = getrow()

		if !moved {
			break
		}
	}

	return min
}
