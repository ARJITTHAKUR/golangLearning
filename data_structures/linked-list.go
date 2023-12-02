package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}
type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) Add(val int) {
	var node = Node{}
	node.property = val
	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = &node
}
func main() {

	var linkedList LinkedList
	linkedList.Add(1)
	linkedList.Add(2)

	fmt.Println(linkedList.headNode.property)
}
