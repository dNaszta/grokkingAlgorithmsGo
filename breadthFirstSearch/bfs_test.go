package breadthFirstSearch

import (
	"fmt"
	"github.com/dominikbraun/graph"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	g := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}

	if !search(g, "you", "thom") {
		t.Errorf("It doesn't work well")
	}
}

func TestBFS(t *testing.T) {
	g := graph.New(graph.IntHash, graph.Acyclic())

	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_ = g.AddVertex(3)
	_ = g.AddVertex(4)
	_ = g.AddVertex(5)
	_ = g.AddVertex(6)
	_ = g.AddVertex(7)
	_ = g.AddVertex(8)

	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(1, 3)
	_ = g.AddEdge(1, 4)
	_ = g.AddEdge(3, 5)
	_ = g.AddEdge(3, 6)
	_ = g.AddEdge(2, 6)
	_ = g.AddEdge(4, 7)
	_ = g.AddEdge(4, 8)

	startElement := 1
	seachable := 4
	_ = graph.BFS(g, startElement, func(value int) bool {
		fmt.Println(value)
		if value == seachable {
			fmt.Println("Found!")
			return true
		}
		return false
	})
}
