package function

import "fmt"

// MapKeys any is an alias for interface{}
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

// Push 需要注意的是必须要一直带上type parameters, 类型是List[T]而不是List
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func W() {
	var m = map[int]string{1: "1", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))
	//可以不指明MapKeys中的K和V（如上），可以指明K和V（如下）；但是！ 类型中的泛型一定要指明！
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	fmt.Println("list:", lst.GetAll())

}
