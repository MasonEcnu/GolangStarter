package main

import (
	"container/list"
	"fmt"
)

// 列表是一种非连续的存储容器
// 列表==链表
// 实现原理是双链表
// 可以分别在头尾插入元素
func main() {
	// 两种声明方式效果相同
	var firstList list.List
	//secondList := list.New()
	fmt.Println(firstList)
	//  *list.Element
	ele := firstList.PushBack(1)
	firstList.InsertAfter(3, ele)
	firstList.PushFront(2)
	for i := firstList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println(firstList.Len())
	//firstList.Remove(ele)

	// 常用方法
	// InsertAfter(v interface {}, mark * Element) * Element 	在 mark 点之后插入元素，mark 点由其他插入函数提供
	// InsertBefore(v interface {}, mark * Element) *Element 	在 mark 点之前插入元素，mark 点由其他插入函数提供
	// PushBackList(other *List) 	添加 other 列表元素到尾部
	// PushFrontList(other *List) 	添加 other 列表元素到头部
}
