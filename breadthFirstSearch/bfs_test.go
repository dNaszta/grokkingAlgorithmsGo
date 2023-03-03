package breadthFirstSearch

import (
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}

	if !search(graph, "you", "thom") {
		t.Errorf("It doesn't work well")
	}
}
