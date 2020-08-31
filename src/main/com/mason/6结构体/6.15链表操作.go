package main

import (
	"fmt"
	"strconv"
)

func main() {
	testShowSingleList()

	testHeadInsertSingleList()

	testTailInsertSingleList()
}

// 单链表
type SingleNode struct {
	Value int
	Next  *SingleNode
}

// 遍历
func showSingleList(head *SingleNode) {
	var s string
	for head != nil {
		s += strconv.Itoa(head.Value) + "-->"
		head = head.Next
	}
	s += "nil"
	fmt.Println(s)
}

func testShowSingleList() {
	var head = new(SingleNode)
	head.Value = 1
	var node1 = new(SingleNode)
	node1.Value = 2
	head.Next = node1
	var node2 = new(SingleNode)
	node2.Value = 3
	node1.Next = node2
	showSingleList(head)
}

func testHeadInsertSingleList() {
	var head = new(SingleNode)
	head.Value = 0
	var tail *SingleNode
	tail = head //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		var node = SingleNode{Value: i}
		node.Next = tail //将新插入的node的next指向头结点
		tail = &node     //重新赋值头结点
	}
	showSingleList(tail) //遍历结果
}

func testTailInsertSingleList() {
	var head = new(SingleNode)
	head.Value = 0
	var tail *SingleNode
	tail = head //tail用于记录最末尾的结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		var node = SingleNode{Value: i}
		(*tail).Next = &node
		tail = &node
	}
	showSingleList(head) //遍历结果
}
