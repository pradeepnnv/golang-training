package algods

type node struct {
	val  int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) addNodeAtTheBeginning(i int) {
	n := &node{i, nil}
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
