package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length int
	head   *ListItem
	tail   *ListItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	if l.length == 0 {
		return nil
	}
	return l.head
}

func (l *list) Back() *ListItem {
	if l.length == 0 {
		return nil
	}
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := ListItem{Value: v}
	switch l.Len() {
	case 0:
		l.head = &newItem
		l.tail = &newItem
	case 1:
		l.head = &newItem
		l.head.Next = l.tail
		l.tail.Prev = l.head
	default:
		oldHead := l.head
		l.head = &newItem
		l.head.Next = oldHead
		oldHead.Prev = l.head
	}
	l.length++
	return &newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := ListItem{Value: v}
	switch l.Len() {
	case 0:
		l.head = &newItem
		l.tail = &newItem
	case 1:
		l.tail = &newItem
		l.tail.Prev = l.head
		l.head.Next = l.tail
	default:
		oldTail := l.tail
		l.tail = &newItem
		l.tail.Prev = oldTail
		oldTail.Next = l.tail
	}
	l.length++
	return &newItem
}

func (l *list) Remove(i *ListItem) {
	switch l.Len() {
	case 0:
		return
	case 1:
		if l.head == i && l.tail == i {
			l.head = &ListItem{}
			l.tail = &ListItem{}
		}
	default:
		switch i {
		case l.head:
			l.head = l.head.Next
			l.head.Prev = nil
		case l.tail:
			l.tail = l.tail.Prev
			l.tail.Next = nil
		default:
			i.Next.Prev = i.Prev
			i.Prev.Next = i.Next
		}
	}
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i.Value)
	l.Remove(i)
}

func NewList() List {
	newList := new(list)
	newList.length = 0
	newList.head = new(ListItem)
	newList.tail = new(ListItem)
	return newList
}
