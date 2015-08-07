package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/morgan-and-a-string

func main() {
	r := bufio.NewReader(os.Stdin)
	d, _ := r.ReadBytes('\n')
	n, _ := strconv.Atoi(strings.Trim(string(d), "\r\n"))

	for i := 0; i < n; i++ {
		d, _ = r.ReadBytes('\n')
		a := strings.Trim(string(d), "\r\n")
		d, _ = r.ReadBytes('\n')
		b := strings.Trim(string(d), "\r\n")
		Println(getstr(a, b))
	}
}

func lt(a, b string) bool {
	return (a + b) < (b + b)
}

func totk(tk string, a []string, i int, b []string, j int, f *bytes.Buffer) (int, int, bool) {
	for {
		if i < len(a) && j < len(b) {
			if lt(a[i], b[j]) {
				f.WriteString(a[i])
				i++
			} else if lt(b[j], a[i]) {
				f.WriteString(b[j])
				j++
			} else {
				return i, j, false
			}
		} else {
			f.WriteString(strings.Join(a[i:], ""))
			f.WriteString(strings.Join(b[j:], ""))
			return len(a), len(b), true
		}
	}
}

// Returns token v+[<v]+, true or "", false.
func tk(v string, s []string, i int) (string, int, bool) {
	j := i
	for j < len(s) && s[j] == v {
		j++
	}
	for j < len(s) && lt(s[j], v) {
		j++
	}

	if j-i < 2 {
		return "", i, false
	}
	return strings.Join(s[i:j], ""), j, true
}

func tkns(v string, s []string, i int) ([]string, int) {
	r := make([]string, 0)
	d := false
	st := ""
	for {
		st, i, d = tk(v, s, i)
		if !d {
			break
		}
		r = append(r, st)
	}
	return r, i
}

func mergetokens(tk string, a []string, b []string, f *bytes.Buffer) {
	i := 0
	j := 0
	d := false
	for {
		i, j, d = totk(tk, a, i, b, j, f)

		if d {
			break
		}

		// a[i] == b[j]

		if i == len(a)-1 && j == len(b)-1 {
			f.WriteString(a[i])
			f.WriteString(b[j])
			break
		}

		ntk := a[i]

		var sa []string
		sa, i = tkns(ntk, a, i)

		var sb []string
		sb, j = tkns(ntk, b, j)

		if 0 == len(sa) && 0 == len(sb) {
			for i < len(a) && a[i] == ntk {
				f.WriteString(a[i])
				i++
			}
			for j < len(b) && b[j] == ntk {
				f.WriteString(b[j])
				j++
			}
		}

		mergetokens(ntk, sa, sb, f)
	}
}

func getstr(as, bs string) string {
	a := make([]string, len(as))
	for i := 0; i < len(a); i++ {
		a[i] = as[i : i+1]
	}

	b := make([]string, len(bs))
	for i := 0; i < len(b); i++ {
		b[i] = bs[i : i+1]
	}

	var f bytes.Buffer

	// body
	i := 0
	j := 0
	d := false
	for {
		i, j, d = totk("", a, i, b, j, &f)

		if d {
			break
		}

		ntk := a[i]

		var sa []string
		sa, i = tkns(ntk, a, i)

		var sb []string
		sb, j = tkns(ntk, b, j)

		if 0 == len(sa) && 0 == len(sb) {
			for i < len(a) && a[i] == ntk {
				f.WriteString(a[i])
				i++
			}
			for j < len(b) && b[j] == ntk {
				f.WriteString(b[j])
				j++
			}
		}

		mergetokens(ntk, sa, sb, &f)
	}

	return f.String()
}
