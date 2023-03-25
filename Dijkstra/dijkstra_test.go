package Dijkstra

import (
	"github.com/dominikbraun/graph"
	"reflect"
	"testing"
)

// no cycle, weighted graph algorithm, no-negative-weight
// Bellman-Ford algorithm able to use negative-weight too

func TestDijekstra(t *testing.T) {
	start := "start"
	graph := map[string]map[string]int{
		start: {"a": 6, "b": 2},
		"a":   {"fin": 1},
		"b":   {"a": 3, "fin": 5},
		"fin": {},
	}

	result := FindLowestCosts(graph, start)
	wanted := map[string]int{start: 0, "a": 5, "b": 2, "fin": 6}
	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("dijekstra failed, result: %v, wanted: %v", result, wanted)
	}
}

func TestShortestPath(t *testing.T) {
	g := graph.New(graph.StringHash, graph.Weighted(), graph.Directed())

	g.AddVertex("start")
	g.AddVertex("a")
	g.AddVertex("b")
	g.AddVertex("fin")

	g.AddEdge("start", "a", graph.EdgeWeight(6))
	g.AddEdge("start", "b", graph.EdgeWeight(2))
	g.AddEdge("a", "fin", graph.EdgeWeight(1))
	g.AddEdge("b", "a", graph.EdgeWeight(3))
	g.AddEdge("b", "fin", graph.EdgeWeight(5))

	path, _ := graph.ShortestPath(g, "start", "fin")

	wanted := []string{"start", "b", "a", "fin"}
	if !reflect.DeepEqual(path, wanted) {
		t.Errorf("shortest path failed, result: %v, wanted: %v", path, wanted)
	}
}
