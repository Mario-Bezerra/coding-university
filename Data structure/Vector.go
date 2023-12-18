package vector

import "errors"

type Vector[T comparable] struct {
	Vector []T
	Size   int
}

type Comparer interface {
	Compare(other interface{}) int
}

func NewVector[T comparable]() *Vector[T] {
	return &Vector[T]{
		Vector: make([]T, 16),
		Size:   0,
	}
}

func (v *Vector[T]) size() int {
	return v.Size
}

func (v *Vector[T]) resize(newCapacity int) {
	newVector := make([]T, newCapacity)
	copy(newVector[:v.Size], v.Vector[:v.Size])
	v.Vector = newVector
}

func (v *Vector[T]) capacity() int {
	return len(v.Vector)
}

func (v *Vector[T]) is_empty() bool {
	return v.Size == 0
}

func (v *Vector[T]) item_at(index int) (T, error) {
	var result T

	if index < 0 || index >= v.Size {
		return result, errors.New("Index out of bounds")
	}
	return v.Vector[index], nil
}

func (v *Vector[T]) push(object T) {
	if v.Size == v.capacity() {
		v.resize(v.capacity() * 2)
	}
	v.Vector[v.Size] = object
	v.Size++
}

func (v *Vector[T]) insert_at(index int, item T) (T, error) {
	var result T

	if v.Size == v.capacity() {
		v.resize(v.capacity() * 2)
	}

	if index < 0 || index >= v.Size {
		return result, errors.New("Index out of bounds")
	}

	for i := index; i < v.Size; i++ {
		temp_swap := v.Vector[i]
		v.Vector[i+1] = temp_swap
	}
	v.Vector[index] = item

	v.Size++
	return v.Vector[index], nil
}

func (v *Vector[T]) prepend(item T) {
	if v.Size == v.capacity() {
		v.resize(v.capacity() * 2)
	}

	for i := 0; i < v.Size; i++ {
		temp_swap := v.Vector[i]
		v.Vector[i+1] = temp_swap
	}

	v.Size++
	v.Vector[0] = item
}

func (v *Vector[T]) pop() T {
	temp := v.Vector[v.Size]
	v.Vector = v.Vector[:v.Size-1]

	v.Size--

	return temp
}

func (v *Vector[T]) delete(index int) error {
	if index < 0 || index >= v.Size {
		return errors.New("Index out of bounds")
	}

	for i := index + 1; i < v.Size; i++ {
		temp_swap := v.Vector[i]
		v.Vector[i-1] = temp_swap
	}

	v.Size--

	return nil
}

func (v *Vector[T]) remove(item T) error {
	for i := v.Size - 1; i >= 0; i-- {
		if v.Vector[i] == item {
			v.Vector = append(v.Vector[:i], v.Vector[i+1:]...)
		}
	}

	v.Size--

	return nil
}

func (v *Vector[T]) find(item T) (T, error) {
	var result T
	for i := v.Size - 1; i >= 0; i-- {
		if v.Vector[i] == item {
			return v.Vector[i], nil
		}
	}
	return result, errors.New("Item not found")
}
