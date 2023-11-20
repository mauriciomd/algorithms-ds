package graph

import (
	"fmt"

	"github.com/mauriciomd/algorithms-ds/queue"
	"github.com/mauriciomd/algorithms-ds/stack"
)

type AdjacencyMatrix struct {
	graph      [][]int
	isDirected bool
}

func NewAdjacencyMatrix(size int, isDirected bool) *AdjacencyMatrix {
	matrix := make([][]int, size)

	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	return &AdjacencyMatrix{
		graph:      matrix,
		isDirected: isDirected,
	}
}

func (g *AdjacencyMatrix) AddEdge(i, j, weight int) {
	g.graph[i][j] = weight

	if !g.isDirected {
		g.graph[j][i] = weight
	}
}

func (g *AdjacencyMatrix) RemoveEdge(i, j int) {
	g.graph[i][j] = 0

	if !g.isDirected {
		g.graph[j][i] = 0
	}
}

func (g *AdjacencyMatrix) Print() {
	for i := 0; i < len(g.graph); i++ {
		fmt.Println(g.graph[i])
	}
}

func (g *AdjacencyMatrix) BFS(from, to int) []int {
	path := []int{}
	q := queue.New[int]()
	q.Enqueue(from)

	visited := make(map[int]bool)

	for !q.IsEmpty() {
		node := q.Dequeue()

		if visited[node] {
			continue
		}

		path = append(path, node)
		visited[node] = true

		for i := 0; i < len(g.graph); i++ {
			if g.graph[node][i] == 1 {
				q.Enqueue(i)
			}
		}
	}

	return path
}

func (g *AdjacencyMatrix) DFS(from, to int) []int {
	path := []int{}
	seen := make([]bool, len(g.graph))
	s := stack.New[int]()

	s.Push(from)

	for !s.IsEmpty() {
		node := s.Pop()

		if seen[node] {
			continue
		}

		path = append(path, node)
		seen[node] = true

		for i := 0; i < len(seen); i++ {
			if g.graph[node][i] == 1 {
				s.Push(i)
			}
		}
	}

	return path
}
