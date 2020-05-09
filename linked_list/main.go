package main

import (
	"fmt"
	"strconv"
)

// node implementation
type node struct {
	data string
	next *node
}

func (n node) String() string {
	return "node: " + n.data
}

//linkedList implementation

type linkedList struct {
	name        string
	head        *node
	currentNode *node
}

func (l linkedList) String() string {
	return "linkedList: " + l.name + "\n" + "with total " + strconv.Itoa(l.Length()) + " nodes\n\n" + "head node:\n" + l.head.data
}

func (l *linkedList) Add(data string) {

	node := node{
		data: data,
		next: nil,
	}

	if l.head == nil {
		l.head = &node
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		l.currentNode = &node
		currentNode.next = &node
	}
}

func (l linkedList) Length() int {
	length := 0
	if l.head == nil {
		return length
	} else {
		currentNode := l.head
		length++
		for currentNode.next != nil {
			length++
			currentNode = currentNode.next
		}
		return length
	}
}

func (l linkedList) traverseNodes() {
	if l.head == nil {
		return
	}
	currentNode := l.head
	fmt.Println(currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Println(currentNode)
	}
}

//reset the pointr of current node to head node
func (l *linkedList) resetCurrentNode() {
	if l.head == nil {
		return
	}
	l.currentNode = l.head
}

//set the current node pointer of linkedList to next node
func (l *linkedList) next() {
	if l.head == nil {
		return
	}
	if l.currentNode.next == nil {
		fmt.Println("This is the last node!")
	} else {
		l.currentNode = l.currentNode.next
	}
}
func main() {
	list := linkedList{
		name:        "my first linkedList",
		head:        nil,
		currentNode: nil,
	}
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")
	//list.traverseNodes()
	list.resetCurrentNode()
	fmt.Println(list.currentNode)
	list.next()
	fmt.Println(list.currentNode)
	list.next()
	fmt.Println(list.currentNode)
	list.next()
	fmt.Println(list.currentNode)
	list.next()
	fmt.Println(list.currentNode)

}
