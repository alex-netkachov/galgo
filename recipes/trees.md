Simple Tree
---

```go
type INode interface {
	Nodes() []INode
	Append(INode)
	Remove(INode)
}

type NodeInt struct {
	id    int
	key   int
	nodes []INode
	sum   int
}

func (this *NodeInt) ID() int {
	return this.id
}

func (this *NodeInt) Key() int {
	return this.key
}

func (this *NodeInt) Nodes() []INode {
	return this.nodes
}

func (this *NodeInt) Append(node INode) {
	this.nodes = append(this.nodes, node)
}

func (this *NodeInt) Remove(node INode) {
	n, _ := node.(*NodeInt)
	for k, v := range this.nodes {
		if v == n {
			this.nodes = append(this.nodes[0:k], this.nodes[k+1:]...)
			return
		}
	}
}
```

DFS
---

```go
func DFS(node INode, callback func(INode)) {
	for _, n := range node.Nodes() {
		DFS(n, callback)
	}
	callback(node)
}
```
