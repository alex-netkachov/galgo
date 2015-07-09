IO
==

Reading
---

### Array of ints

```go
import (
  "bufio"
  "strconv"
)

func ReadInts(s bufio.Scanner) []int {
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

func ReadInt(s bufio.Scanner) int {
	s.Split(bufio.ScanWords)
	
  s.Scan()
	i, _ := strconv.Atoi(s.Text())

	return i
}
```
