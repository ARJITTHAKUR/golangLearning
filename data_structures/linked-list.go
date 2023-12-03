package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}
type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(val int) {
	var node = Node{}
	node.property = val
	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = &node
}

func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node

	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedList *LinkedList) AddToEnd(val int) {
	var newNode = &Node{property: val}
	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = newNode
	}
}
func main() {

	var linkedList LinkedList
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	linkedList.AddToEnd(56)

	fmt.Println(linkedList.headNode.property)
	fmt.Println("last node value is", linkedList.LastNode())
}
