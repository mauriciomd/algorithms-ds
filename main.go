package main

import (
	"fmt"

	"github.com/mauriciomd/algorithms-ds/graph"
	"github.com/mauriciomd/algorithms-ds/lists"
	"github.com/mauriciomd/algorithms-ds/queue"
)

func testSinglyLinkedList() {
	fmt.Println("Using Singly Linked List")
	singlyLinkedList := lists.NewSinglyLinkedList[int]()

	singlyLinkedList.Append(1)
	singlyLinkedList.Append(3)
	singlyLinkedList.Append(6)
	singlyLinkedList.Preppend(25)
	singlyLinkedList.Append(11)
	singlyLinkedList.Append(45)

	singlyLinkedList.PrintAllElements()
	singlyLinkedList.Remove(2)

	fmt.Println()
	singlyLinkedList.PrintAllElements()
}

func testDoublyLinkedList() {
	fmt.Println("Using Doubly Linked List")
	doublyLinkedList := lists.NewDoublyLinkedList[int]()
	doublyLinkedList.Append(1)
	doublyLinkedList.Append(3)
	doublyLinkedList.Append(6)
	doublyLinkedList.Preppend(25)
	doublyLinkedList.Append(11)
	doublyLinkedList.Append(45)
	doublyLinkedList.PrintAllElements()
	doublyLinkedList.Remove(2)
	doublyLinkedList.InsertAt(99, 1)
	fmt.Println()
	doublyLinkedList.PrintAllElements()
	fmt.Println()
	doublyLinkedList.PrintAllReversedElements()
}

func testQueue() {
	fmt.Println("Using Queue")
	q := queue.New[int]()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.PrintQueue()
	fmt.Println("Dequeue", q.Dequeue())
	fmt.Println("Dequeue", q.Dequeue())
	fmt.Println("Dequeue", q.Dequeue())
}

func testAdjacencyMatrixGraph() {
	fmt.Println("Using Graph - Adjacency Matrix")

	// Source: freecodecamp.com
	g := graph.NewAdjacencyMatrix(6, false)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 4, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 4, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(3, 5, 1)

	g.Print()
}

func main() {
	testSinglyLinkedList()
	fmt.Println()
	testDoublyLinkedList()
	fmt.Println()
	testQueue()
	fmt.Println()
	testAdjacencyMatrixGraph()
	fmt.Println()
}
