package main

import "strings"

func assert(condition bool, msg ...string) {
	if !condition {
		panic("assertion failed: " + strings.Join(msg, " "))
	}
}

type Registry struct {
	Nodes map[string]*Node
}

func (r *Registry) CreateNode(id string, data ...string) *Node {
	node := &Node{ID: id, Data: data[0]}
	assert(r.Nodes[id] == nil, "node already exists")
	r.Nodes[id] = node
	return node
}

type Node struct {
	ID   string
	Data string
}

func main() {
	registry := &Registry{}

	registry.CreateNode("1", "Hello, World!")

	assert(len(registry.Nodes) == 1)
	assert(registry.Nodes["1"].ID == "1")
	assert(registry.Nodes["1"].Data == "Hello, World!")
}
