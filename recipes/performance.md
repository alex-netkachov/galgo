Performance
===

Memory
---

### Read Memory Usage

See [runtime.MemStats](https://golang.org/pkg/runtime/#MemStats) for more details.

```go
import "runtime"

...

  var m runtime.MemStats
  
  ...
  
  runtime.ReadMemStats(&m)
  Println("HeapAlloc:", m.HeapAlloc)
```
