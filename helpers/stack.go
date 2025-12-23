package helpers

type Stack[T any] struct {
	list []T
}

func (o *Stack[T]) Len() int {
	return len(o.list)
}

func (o *Stack[T]) IsEmpty() bool {
	return o.Len() == 0
}

func (o *Stack[T]) Push(item T) {
	if o.list == nil {
		o.list = make([]T, 0)
	}

	o.list = append(o.list, item)
}

func (o *Stack[T]) Pop() T {
	if o.list == nil {
		var def T
		return def
	}

	r := o.list[len(o.list)-1]

	o.list = o.list[:len(o.list)-1]

	return r

}

func (o *Stack[T]) Reverse() {
	o.list = ReverseArray(o.list)
}
