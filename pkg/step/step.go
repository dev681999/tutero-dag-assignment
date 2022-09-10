package step

import "tutero_assignment/pkg/src/graph"

type Stepper interface {
	// Step returns a prediction for the correct node; or an error if a prediction cannot be made.
	Step(graph graph.Graph) (graph.Node, error)
}

func New() *stepper {
	//* You may mutate this instantiation if necessary; but the function signature should not change.
	return &stepper{
		visited: make(map[graph.Node]struct{}),
	}
}

type stepper struct {
	//* You may add fields to this struct.
	children []graph.Node
	visited  map[graph.Node]struct{}
}

func (s *stepper) Step(grp graph.Graph) (graph.Node, error) {
	al := grp.AdjacencyList()

	if len(s.children) == 0 {
		nodes, _ := grp.TopologicalSort()
		s.children = []graph.Node{}
		s.children = append(s.children, nodes[len(nodes)-1])
	}

	for len(s.children) > 0 {
		currNode := s.children[0]
		s.children = s.children[1:]
		// if its not in graph check next
		_, ok := al[currNode]
		if !ok {
			continue
		}

		_, ok = s.visited[currNode]
		if ok {
			continue
		}

		s.visited[currNode] = struct{}{}

		s.children = append(s.children, grp.Parents(currNode)...)
		return currNode, nil
	}

	nodes, _ := grp.TopologicalSort()
	return nodes[len(nodes)-1], nil
}
