package main

import (
	"strings"
)

func assert(condition bool, msg ...string) {
	if !condition {
		panic("assertion failed: " + strings.Join(msg, " "))
	}
}

type Registry struct {
	Nodes map[string]*Node
}

func (r *Registry) Init() {
	r.Nodes = make(map[string]*Node)
}

func (r *Registry) CreateNode(id string, data ...string) *Node {
	if len(data) == 0 {
		data = []string{""}
	}
	node := &Node{ID: id, Data: data[0]}
	assert(r.Nodes != nil, "registry is nil")
	assert(r.Nodes[id] == nil, "node already exists")
	node.Registry = r
	r.Nodes[id] = node
	return node
}

type Node struct {
	ID       string
	Data     string
	Registry *Registry
}

func (n *Node) Pull() {
	results := []string{}
	for _, node := range n.Registry.Nodes {
		if node.ID != n.ID {
			results = append(results, node.Data)
		}
	}

	occurances := make(map[string]int)
	for _, result := range results {
		occurances[result]++
	}

	winner := ""
	for _, result := range results {
		if occurances[result] > occurances[winner] {
			winner = result
		}
	}

	n.Data = winner
}

func (n *Node) Push() {
	for _, node := range n.Registry.Nodes {
		if node.ID != n.ID {
			node.Data = n.Data
		}
	}
}

func (n *Node) Edit(data string) {
	if n.Data == data {
		return
	}

	n.Data = data
	n.Push()
}

func main() {
	// init registry
	registry := &Registry{}
	registry.Init()

	// create node 1
	registry.CreateNode("1", "Hello, World!")
	assert(len(registry.Nodes) == 1, "registry has 1 node")
	assert(registry.Nodes["1"].ID == "1", "node 1 has id 1")
	assert(registry.Nodes["1"].Data == "Hello, World!", "node 1 has data: Hello, World!")

	// edit node 1
	registry.Nodes["1"].Edit("Bye, World!")
	assert(registry.Nodes["1"].Data == "Bye, World!", "node 1 has data: Bye, World!")

	// create node 2
	node2 := registry.CreateNode("2")
	node2.Pull()
	assert(len(registry.Nodes) == 2, "registry has 2 nodes")
	assert(registry.Nodes["2"].ID == "2", "node 2 has id 2")
	assert(registry.Nodes["2"].Data == "Bye, World!", "node 2 has data: Bye, World!")

	// edit node 2
	registry.Nodes["2"].Edit("Hello, World!")
	for _, node := range registry.Nodes {
		assert(node.Data == "Hello, World!", "node "+node.ID+" has data: Hello, World!")
	}
}
