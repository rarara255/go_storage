package storage

type Rerposetory[T any] struct{
	items []T
}

func NewRepository[T any]() *Rerposetory[T]{
	return &Rerposetory[T]{
		items: make([]T, 0),
	}
}

func (r *Rerposetory[T]) Add(item T){
	r.items = append(r.items, item)
}

func (r *Rerposetory[T]) GetAll() []T{
	return r.items
}