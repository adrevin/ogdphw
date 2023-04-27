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
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return &list{}
}

func (l *list) PushFront(v interface{}) *ListItem {
	next := l.front
	if next == nil {
		next = l.back
	}

	l.front = &ListItem{Value: v, Next: next}

	if next != nil {
		next.Prev = l.front
	}

	if l.back == nil && l.front.Next != nil {
		l.back = l.front.Next
		l.back.Prev = l.front
	}

	if l.back != nil && l.back.Prev == nil {
		l.back.Prev = l.front
	}

	l.len++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	prev := l.back
	if prev == nil {
		prev = l.front
	}

	l.back = &ListItem{Value: v, Prev: prev}

	if prev != nil {
		prev.Next = l.back
	}

	if l.front == nil && l.back.Prev != nil {
		l.front = l.back.Prev
		l.front.Next = l.back
	}

	if l.front != nil && l.front.Next == nil {
		l.front.Next = l.back
	}

	l.len++
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		panic("null ListItem value of received")
	}
	// if it is front
	if i == l.front {
		// if next exists and it is no back
		if i.Next != nil && i.Next.Next != nil {
			l.front = i.Next
			i.Next.Prev = l.front
			l.front.Prev = nil
		} else {
			l.front = nil
		}
		l.len--
		return
	}
	// if it is back
	if i == l.back {
		// if prev exists and it is no front
		if i.Prev != nil && i.Prev.Prev != nil {
			l.back = i.Prev
			i.Prev.Next = l.back
			l.back.Next = nil
		} else {
			l.back = nil
		}
		l.len--
		return
	}

	i.Next.Prev = i.Prev
	i.Prev.Next = i.Next
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil {
		panic("null ListItem value of received")
	}
	if l.front == i {
		return
	}
	// only back item
	if l.front == nil && i == l.back {
		l.front = i
		l.front.Next = nil
		l.back = nil
		return
	}

	l.Remove(i)
	i.Next = l.front
	l.front.Prev = i

	i.Prev = nil
	l.front = i

	l.len++
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}
