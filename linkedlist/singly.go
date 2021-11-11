package linkedlist

import (
	"bytes"
	"fmt"
)

type SinglyLinkedList struct {
	dummyHead *Node
	size int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{&Node{}, 0}
}

func (l *SinglyLinkedList) GetSize() int {
	return l.size
}

func (l *SinglyLinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *SinglyLinkedList) Add(index int, val interface{}) bool {
	if index < 0 || index > l.size {
		return false
	}

	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = &Node{Val: val, Next: prev.Next}
	l.size++

	return true
}

func (l *SinglyLinkedList) AddFirst(val interface{}) bool {
	return l.Add(0, val)
}

func (l *SinglyLinkedList) AddLast(val interface{}) bool {
	return l.Add(l.size, val)
}

func (l *SinglyLinkedList) Get(index int) (interface{}, bool) {
	if index < 0 || index >= l.size {
		return nil, false
	}

	cur := l.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val, true
}

func (l *SinglyLinkedList) GetFirst() (interface{}, bool) {
	return l.Get(0)
}

func (l *SinglyLinkedList) GetLast() (interface{}, bool) {
	return l.Get(l.size - 1)
}

func (l *SinglyLinkedList) Set(index int, val interface{}) bool {
	if index < 0 || index >= l.size {
		return false
	}

	cur := l.dummyHead.Next
	for i:=0; i<index; i++ {
		cur = cur.Next
	}
	cur.Val = val

	return true
}

func (l *SinglyLinkedList) Contains(val interface{}) bool {
	cur := l.dummyHead.Next
	for cur != nil {
		if cur.Val == val {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (l *SinglyLinkedList) Remove(index int) (interface{}, bool) {
	if index < 0 || index >= l.size {
		return nil, false
	}

	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	retNode := prev.Next
	prev.Next = retNode.Next
	retNode.Next = nil
	l.size--

	return retNode.Val, true
}

func (l *SinglyLinkedList) RemoveFirst() (interface{}, bool) {
	return l.Remove(0)
}

func (l *SinglyLinkedList) RemoveLast() (interface{}, bool) {
	return l.Remove(l.size - 1)
}

func (l *SinglyLinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := l.dummyHead.Next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.Val) + "->")
		cur = cur.Next
	}

	buffer.WriteString("nil")
	return buffer.String()
}

