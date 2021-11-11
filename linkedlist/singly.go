package linkedlist

import (
	"bytes"
	"errors"
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

func (l *SinglyLinkedList) Add(index int, val interface{}) error {
	if index < 0 || index > l.size {
		return errors.New("add failed, illegal index")
	}

	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = &Node{Val: val, Next: prev.Next}
	l.size++

	return nil
}

func (l *SinglyLinkedList) AddFirst(val interface{}) error {
	return l.Add(0, val)
}

func (l *SinglyLinkedList) AddLast(val interface{}) error {
	return l.Add(l.size, val)
}

func (l *SinglyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("get failed, illegal index")
	}

	cur := l.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val, nil
}

func (l *SinglyLinkedList) GetFirst() (interface{}, error) {
	return l.Get(0)
}

func (l *SinglyLinkedList) GetLast() (interface{}, error) {
	return l.Get(l.size - 1)
}

func (l *SinglyLinkedList) Set(index int, val interface{}) error {
	if index < 0 || index >= l.size {
		return errors.New("set failed, illegal index")
	}

	cur := l.dummyHead.Next
	for i:=0; i<index; i++ {
		cur = cur.Next
	}
	cur.Val = val

	return nil
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

func (l *SinglyLinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("remove failed, illegal index")
	}

	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	retNode := prev.Next
	prev.Next = retNode.Next
	retNode.Next = nil
	l.size--

	return retNode.Val, nil
}

func (l *SinglyLinkedList) RemoveFirst() (interface{}, error) {
	return l.Remove(0)
}

func (l *SinglyLinkedList) RemoveLast() (interface{}, error) {
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

