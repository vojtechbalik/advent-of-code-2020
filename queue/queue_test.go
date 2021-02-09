package queue

import "testing"

func TestEmpty(t *testing.T) {
	q := New()
	if q.Len() != 0 {
		t.Fatal("Initial size not 0")
	}
	if item, ok := q.Front(); item != nil || ok {
		t.Fatal("Invalid Front return on empty queue")
	}
	if item, ok := q.Pop(); item != nil || ok {
		t.Fatal("Invalid Pop return on empty queue")
	}
	if q.Len() != 0 {
		t.Fatal("Size altered after non-modifying operations")
	}
}

func TestFewItemsNoExtend(t *testing.T) {
	q := New(3)
	q.Push(0)
	q.Push(1)
	q.Push(2)
	if q.Len() != 3 {
		t.Fatal("Size does not reflect correct number of items")
	}
	q.Pop()
	if item, ok := q.Pop(); item != 1 || !ok {
		t.Fatal("Invalid pop return", item, "expected", 1)
	}
	q.Push(3)
	q.Push(4)
	if q.Len() != 3 {
		t.Fatal("Size not correct after push/pop operations")
	}
	if item, ok := q.Pop(); item != 2 || !ok {
		t.Fatal("Invalid pop return", item, "expected", 2)
	}
}

func TestFewItemsExpand(t *testing.T) {
	q := New(2)
	q.Push(0)
	q.Push(1)
	if q.Len() != 2 {
		t.Fatal("Size does not reflect correct number of items")
	}
	if item, ok := q.Pop(); item != 0 || !ok {
		t.Fatal("Invalid pop return", item, "expected:", 0)
	}
	q.Push(3)
	q.Push(4)
	if q.Len() != 3 {
		t.Fatal("Size not correct after push/pop operations")
	}
	if item, ok := q.Pop(); item != 1 || !ok {
		t.Fatal("Invalid pop return", item, "expected:", 1)
	}
}
