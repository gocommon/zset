package zset

import (
	"fmt"
	"testing"
)

func TestZset(t *testing.T) {
	z := NewZset()

	z.Add(3.4, "sh600600")
	z.Add(3.3, "sh600600")
	z.Add(7.7, "sh600601")
	z.Add(6.6, "sh600602")

	z.Add(5.5, "sh600603")
	z.Add(2.2, "sh600604")
	z.Add(8.8, "sh600605")
	z.Add(9.9, "sh600606")
	z.Add(23.23, "sh600607")
	z.Add(6.6, "sh600608")
	z.Add(6.1, "sh6006010")
	z.Add(6.62, "sh6006011")
	z.Add(6.63, "sh6006012")
	z.Add(6.62, "sh600602")
	z.Add(6.67, "sh600616")
	z.Add(6.68, "sh600615")
	z.Add(6.69, "sh600618")
	z.Add(-6.64, "sh6006013")
	z.Add(6.65, "sh600614")
	z.Add(6.66, "sh600615")
	z.Add(6.61, "sh6006012")
	z.Add(-6.1, "sh600602")

	z.Add(-6.63, "sh6006013")

	t.Log("list len", z.Len())

	// list := z.RevRange(0, -1)
	list := z.Range(0, 3)

	if len(list) == 0 {
		t.Fatal("len(list) == 0 ")
	}

	for i := range list {

		t.Log(list[i].Score(), list[i].Member())
	}

	if !checkList(list) {
		t.Fatal("checkList")
	}
}

func checkList(list []*Obj) bool {
	hasMap := map[string]bool{}
	for i := range list {
		if i > 0 && list[i].Score() < list[i-1].Score() {
			fmt.Println("<", list[i].Score(), list[i-1].Score())
			return false
		}

		_, has := hasMap[list[i].Member()]
		if has {
			fmt.Println("has", list[i].Member())
			return false
		}
		hasMap[list[i].Member()] = true
	}

	return true
}
