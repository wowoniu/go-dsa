package singlylinkedlist

import (
	"errors"
	"fmt"
)

//Node 单链表节点
type Node struct {
	Data interface{}
	Next *Node
}

//List 单链表
type List struct {
	Head   *Node
	Tail   *Node
	Length int32
}

//NewList 生成一个新的单链表
func NewList() *List {
	//头节点
	fenceNode := &Node{
		Data: nil,
		Next: nil,
	}
	list := &List{
		Head:   fenceNode,
		Tail:   fenceNode,
		Length: 0,
	}

	return list
}

//GetNode 获取第pos个节点
func (l *List) GetNode(pos int32) (node *Node, err error) {
	var i int32 = -1
	if pos > l.Length-1 {
		err = errors.New("无效的节点")
		return
	}
	node = l.Head
	for i < pos && node != nil {
		node = node.Next
		i = i + 1
	}

	return
}

//Insert 在链表中插入一个节点
func (l *List) Insert(pos int32, value interface{}) (err error) {
	if pos < 0 || pos > l.Length-1 {
		err = errors.New("无效的插入位置")
		return
	}
	newNode := &Node{
		Data: value,
	}
	//找到前驱节点
	prevNode, err := l.GetNode(pos - 1)
	if err != nil {
		return
	}
	posNode := prevNode.Next
	prevNode.Next = newNode
	newNode.Next = posNode
	l.Length = l.Length + 1
	return
}

//AppendBy 在链表制定位置中追加节点
func (l *List) AppendBy(pos int32, value interface{}) (err error) {
	posNode, err := l.GetNode(pos)
	if err != nil {
		return
	}
	//创建新节点
	newNode := &Node{
		Data: value,
	}
	l.Length = l.Length + 1
	//追加
	if l.Tail == posNode {
		l.Tail.Next = newNode
		l.Tail = newNode
		return
	}

	nextNode := posNode.Next
	posNode.Next = newNode
	newNode.Next = nextNode

	return

}

//Append 在链表尾部追加
func (l *List) Append(value interface{}) (err error) {
	return l.AppendBy(l.Length-1, value)
}

//Remove 移除节点
func (l *List) Remove(pos int32) (err error) {
	posNode, err := l.GetNode(pos)
	if err != nil {
		return
	}
	//遍历获取前驱节点
	node := l.Head
	var prevNode *Node
	for node.Next != posNode {
		node = node.Next
	}
	prevNode = node
	//后继节点
	prevNode.Next = posNode.Next
	l.Length = l.Length - 1
	return
}

//Travel 链表遍历
func (l *List) Travel(callback func(node *Node)) {
	fmt.Printf("链表遍历 长度:%v\n", l.Length)
	node := l.Head
	for node.Next != nil {
		node = node.Next
		if callback != nil {
			callback(node)
		} else {
			fmt.Printf("%p-%v\n", node, node.Data)
		}
	}
}

//Reverse 链表反转
func (l *List) Reverse() {
	head := l.Head
	var prevNode *Node
	currentNode := head.Next
	l.Tail = currentNode
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = nil
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
		if currentNode != nil {
			l.Head.Next = currentNode
		}
	}
}
