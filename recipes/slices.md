Slices
===

Basic
---

### Reverse

To new array:

```go
func revd(a []int) []int {
  n := len(a)
  r := make([]int, n)
  copy(r, a)
  for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
    r[i], r[j] = r[j], r[i]
  }
  return r
}
```

In the same array:

```go
func rev(a []int) {
  for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
  }
}
```

### Merge

Merges two slices into one:

```go
func merge(a, b []int) []int {
  r, ir, ia, ib := make([]int, len(a) + len(b)), 0, 0, 0
  for {
    if ia == len(a) && ib == len(b) {
      break
    }
    v := 0
    if ia == len(a) {
      v = b[ib]
      ib++
    } else if ib == len(b) {
      v = a[ia]
      ia++
    } else {
      va, vb := a[ia], b[ib]
      if va < vb {
        v = va
        ia++
      } else {
        v = vb
        ib++
      }
    }
    r[ir] = v
    ir++
  }
  return r
}
```

### Max/Min

```go
func max(a []int) (int, int) {
  m := a[0]
  mi := 0
  for i := 1; i < len(a); i++ {
    if a[i] > m {
      m = a[i]
      mi = i
    }
  }
  return m, mi
}

func min(a []int) (int, int) {
  m := a[0]
  mi := 0
  for i := 1; i < len(a); i++ {
    if a[i] < m {
      m = a[i]
      mi = i
    }
  }
  return m, mi
}
```
