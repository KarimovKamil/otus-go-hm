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
	len   int
	first *ListItem
	last  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	newFirst := ListItem{v, l.first, nil}

	if l.first != nil {
		l.first.Prev = &newFirst
	}
	if l.last == nil {
		l.last = &newFirst
	}

	l.first = &newFirst
	l.len++
	return &newFirst
}

func (l *list) PushBack(v interface{}) *ListItem {
	newLast := ListItem{v, nil, l.last}

	if l.last != nil {
		l.last.Next = &newLast
	}
	if l.first == nil {
		l.first = &newLast
	}

	l.last = &newLast
	l.len++
	return &newLast
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.first = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.last = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next

		if i.Next != nil {
			i.Next.Prev = i.Prev
		} else {
			l.last = i.Prev
		}

		i.Prev = nil
		l.first.Prev = i
		i.Next = l.first
		l.first = i
	}
}

func NewList() List {
	return new(list)
}
