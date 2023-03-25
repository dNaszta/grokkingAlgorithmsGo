package Dijkstra

import (
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
