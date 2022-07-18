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

// Len() int
func (l *list) Len() int {
	return l.length
}

// Front() *ListItem
func (l *list) Front() *ListItem {
	if l.length == 0 {
		return nil
	}
	return l.head.Next
}

// Back() *ListItem
func (l *list) Back() *ListItem {
	if l.length == 0 {
		return nil
	}
	return l.tail.Prev
}

// PushFront(v interface{}) *ListItem
func (l *list) PushFront(v interface{}) *ListItem {
	newItem := ListItem{Value: v.(int), Next: l.head.Next, Prev: l.head}
	l.head.Next.Prev = &newItem
	l.head.Next = &newItem
	l.length++
	return &newItem
}

// PushBack(v interface{}) *ListItem
func (l *list) PushBack(v interface{}) *ListItem {
	newItem := ListItem{Value: v.(int), Next: l.tail, Prev: l.tail.Prev}
	l.tail.Prev.Next = &newItem
	l.tail.Prev = &newItem
	l.length++
	return &newItem
}

// Remove(i *ListItem)
func (l *list) Remove(i *ListItem) {
	i.Next.Prev = i.Prev
	i.Prev.Next = i.Next
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i)
	l.Remove(i)
}

func NewList() List {
	newList := new(list)
	newList.length = 0
	newList.head = new(ListItem)
	newList.tail = new(ListItem)
	newList.head.Next = newList.tail
	newList.tail.Prev = newList.head
	return newList
}
