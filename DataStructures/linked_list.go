package DataStructures

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) addLast(data string) {
	if list.head == nil {
		list.head = &Node{data, nil}
		return
	}

	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = &Node{data, nil}
}

func (list *LinkedList) addFirst(data string) {
	if list.head == nil {
		list.head = &Node{data, nil}
		return
	}

	tmp := &Node{data, list.head}
	list.head = tmp
}

func (list *LinkedList) getByName(data string) *Node {
	if list.head == nil {
		fmt.Println("Empty list")
	}

	cur := list.head
	for cur.next != nil {
		if cur.data == data {
			return cur
		}
		cur = cur.next
	}
	if cur.data == data {
		return cur
	}
	return nil
}

func (list *LinkedList) removeNode(data string) {
	if list.head == nil {
		fmt.Println("Empty list")
		return
	}

	if list.head.data == data {
		list.head = list.head.next
		return
	}

	cur := list.head
	for cur.next != nil {
		if cur.next.data == data {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
	if cur.data == data {
		cur.next = nil
	}
}

func (list *LinkedList) printList() {
	if list.head == nil {
		fmt.Println("Empty list")
	}

	cur := list.head
	for cur.next != nil {
		fmt.Println(cur)
		cur = cur.next
	}
	fmt.Println(cur)
}

func (node *Node) String() string {
	if node.next == nil {
		return fmt.Sprintf("Data: %s, Next: <nil>", node.data)
	}
	return fmt.Sprintf("Data: %s, Next: %v", node.data, *node.next)
}
