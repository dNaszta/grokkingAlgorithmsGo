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
	g := graph.New(graph.StringHash, graph.Acyclic(), graph.Directed())

	_ = g.AddVertex("you")
	_ = g.AddVertex("alice")
	_ = g.AddVertex("bob")
	_ = g.AddVertex("claire")
	_ = g.AddVertex("anuj")
	_ = g.AddVertex("peggy")
	_ = g.AddVertex("thom")
	_ = g.AddVertex("jonny")

	_ = g.AddEdge("you", "alice")
	_ = g.AddEdge("you", "bob")
	_ = g.AddEdge("you", "claire")
	_ = g.AddEdge("bob", "anuj")
	_ = g.AddEdge("bob", "peggy")
	_ = g.AddEdge("alice", "peggy")
	_ = g.AddEdge("claire", "thom")
	_ = g.AddEdge("claire", "jonny")

	startElement := "you"
	seachable := "thom"
	_ = graph.BFS(g, startElement, func(value string) bool {
		fmt.Println(value)
		if value == seachable {
			fmt.Println("Found!")
			return true
		}
		return false
	})
}
