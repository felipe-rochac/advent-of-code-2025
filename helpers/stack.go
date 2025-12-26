package helpers

type Stack[T any] struct {
	list []T
}

func FromArray[T any](arr []T) Stack[T] {
	return Stack[T]{
		list: arr,
	}
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

// push array to the top of the stack
func (o *Stack[T]) PushArray(arr []T) {
	for _, i := range arr {
		o.list = append(o.list, i)
	}
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
