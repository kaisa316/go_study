package datastructure

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head Node
}

//头部 插入一个元素
func (list LinkedList) Lpush(node Node) LinkedList {
	//边界值考虑
	if list.head == nil {
		list.head = node
	}
	return list
}
