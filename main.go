package main

import (
	"fmt"

	"github.com/mauriciomd/algorithms-ds/lists"
)

func main() {
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
