package breadthFirstSearch

import (
	"container/list"
	"fmt"
)

type void struct{}

var member void

func search(graph map[string][]string, startElement, searchable string) bool {
	set := make(map[string]void)
	queue := list.New()
	// Set first degree
	for _, s := range graph[startElement] {
		queue.PushBack(s)
	}
	for queue.Len() > 0 {
		// Get element from queue
		next := queue.Front()
		check := fmt.Sprintf("%v", next.Value)
		queue.Remove(next)
		if _, ok := set[check]; ok {
			continue
		}
		// Set element in uncheckable set
		set[check] = member
		// Check value
		if check == searchable {
			fmt.Println(searchable)
			fmt.Println(set)
			return true
		} else {
			for _, s := range graph[check] {
				queue.PushBack(s)
			}
		}
	}
	return false
}
