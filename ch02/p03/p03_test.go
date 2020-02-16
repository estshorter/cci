package p03

import (
	"ctci/container/list"
	"fmt"
	"testing"
)

func equal(l1 *list.List, l2 *list.List) bool {
	if l1.Len() != l2.Len() {
		fmt.Printf("wrong length: got %v, want %v\n", l1.Len(), l2.Len())
		return false
	}

	e1 := l1.Front()
	e2 := l2.Front()
	for i := 0; i < l1.Len(); i++ {
		if e1.Value != e2.Value {
			fmt.Printf("wrong value: got %v, want %v\n", e1.Value, e2.Value)
			return false
		}
		e1 = e1.Next
		e2 = e2.Next
	}
	return true
}

func TestDeleteElement(t *testing.T) {
	l1 := list.New()
	l1.PushBack(3)
	l1.PushBack(5)
	e3 := l1.PushBack(1)
	l1.PushBack(100)
	l1.PushBack(10)

	l2 := list.New()
	l2.PushBack(3)
	l2.PushBack(5)
	l2.PushBack(100)
	l2.PushBack(10)

	type args struct {
		l    *list.List
		mark *list.Element
	}
	tests := []struct {
		name string
		args args
		want *list.List
	}{
		{"0", args{l1, e3}, l2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteElement(tt.args.l, tt.args.mark); !equal(got, tt.want) {
				t.Errorf("DeleteElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
