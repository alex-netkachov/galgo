I/O
===

Reading
---

### Array of ints

```go
import (
  "bufio"
  "strconv"
)

func ReadInts(s *bufio.Scanner) []int {
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[j], _ = strconv.Atoi(s.Text())
	}
	return a
}
```

### Read int

```go
import (
  "bufio"
  "strconv"
)

func ReadInt(s *bufio.Scanner) int {
	s.Split(bufio.ScanWords)
	
	s.Scan()
	i, _ := strconv.Atoi(s.Text())

	return i
}
```

### Read long string

The problem with the scanner is that it cannot read long strings
so the reader should be used instead ([link](http://stackoverflow.com/questions/21124327)).

```go
import (
  "bufio"
  "strings"
)

func ReadLongString(r bufio.Reader) string {
  s := r.ReadBytes('\n')
  return strings.Trim(s, "\r\n")
}
```
