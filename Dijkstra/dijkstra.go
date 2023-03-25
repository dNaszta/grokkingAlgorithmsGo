package Dijkstra

import "math"

func findLowestCostNode(costs map[string]int, processed map[string]bool) string {
	lowestCost := math.MaxInt32
	lowestCostNode := ""
	for k, v := range costs {
		if v < lowestCost && !processed[k] {
			lowestCost = v
			lowestCostNode = k
		}
	}
	return lowestCostNode
}

func FindLowestCosts(graph map[string]map[string]int, start string) map[string]int {
	costs := make(map[string]int)
	parents := make(map[string]*string)
	for k := range graph {
		costs[k] = math.MaxInt32
		parents[k] = nil
	}
	costs[start] = 0

	processed := make(map[string]bool)

	for len(processed) < len(graph) {
		node := findLowestCostNode(costs, processed)
		neighbors := graph[node]
		for k, v := range neighbors {
			newCost := costs[node] + v
			if costs[k] > newCost {
				costs[k] = newCost
				parents[k] = &node
			}
		}
		processed[node] = true
	}
	return costs
}
