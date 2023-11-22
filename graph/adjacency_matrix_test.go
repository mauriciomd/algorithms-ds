package graph

import (
	"strconv"
	"strings"
	"testing"
)

func TestAdjMatrixGraph(t *testing.T) {
	g := NewAdjacencyMatrix(6, false)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 4, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 4, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(3, 5, 1)

	t.Run("Running BFS", func(t *testing.T) {
		got := strings.Join(Map(g.BFS(0, 5), strconv.Itoa), ",")
		want := "0,1,4,2,3,5"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Running DFS", func(t *testing.T) {
		got := strings.Join(Map(g.DFS(0, 5), strconv.Itoa), ",")
		want := "0,4,3,5,2,1"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestDijkstra(t *testing.T) {
	g := NewAdjacencyMatrix(5, true)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 5)
	g.AddEdge(1, 2, 7)
	g.AddEdge(1, 3, 3)
	g.AddEdge(2, 4, 1)
	g.AddEdge(3, 2, 2)

	got := strings.Join(Map(g.Dijkstra(0, 4), strconv.Itoa), ",")
	want := "4,2,0"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
