package main

import (
	"fmt"
)

/*BinaryNode 二叉树节点
*
 */
type BinaryNode struct {
	Parent     *BinaryNode
	LeftChild  *BinaryNode
	RightChild *BinaryNode
	Height     int32
	Data       interface{}
}

func (b *BinaryNode) InsertAsLC() {

}

func (b *BinaryNode) InsertAsRC() {

}

//层次遍历
func (b *BinaryNode) TravLevel() {

}

//中序遍历
func (b *BinaryNode) TravIn() {

}

//后序遍历
func (b *BinaryNode) TravPost() {

}

func main() {
	fmt.Println("dfdfd")
}
