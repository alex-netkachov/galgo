Simple Tree
---

```go
type INode interface {
    Nodes() []INode
}

type NodeInt struct {
  key int
  nodes []INode
}

func (this *NodeInt) Key() int {
    return this.key
}

func (this *NodeInt) Nodes() []INode {
    return this.nodes
}
```
