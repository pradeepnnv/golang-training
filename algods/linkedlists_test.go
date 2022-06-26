package algods

import "testing"

func TestAddNodeToLinkedList(t *testing.T) {
	list := linkedList{}
	list.addNodeAtTheBeginning(4)
	if list.length != 1 {
		t.Errorf("item not added")
	}
}
