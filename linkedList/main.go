package main

import (
	"fmt"

	"github.com/wowoniu/go-dsa/linkedList/singlylinkedlist"
)

func main() {
	fmt.Println("单链表")
	list := singlylinkedlist.NewList()
	list.Append(23)
	list.AppendBy(0, 24)
	list.Append(25)
	list.Insert(1, "insert VALUE1")
	list.Insert(1, "insert VALUE2")
	list.Insert(1, "insert VALUE3")
	//链表遍历
	list.Travel(func(node *singlylinkedlist.Node) {
		fmt.Printf("%p-%v\n", node, node.Data)
	})
	//链表反转
	list.Reverse()

	fmt.Println(list)
	list.Travel(nil)

	list.Remove(5)
	list.Travel(nil)

	// fmt.Println(list.Head)
}
