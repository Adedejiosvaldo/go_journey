package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func NewCustomerMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

// Generics
func foo[V any](val V) {
	fmt.Printf("%T of %v", val, val)
}
func main() {
	m1 := NewCustomerMap[int, string]()
	m1.Insert(1, "2")
	fmt.Printf("%+v", m1)

	m2 := NewCustomerMap[string, int]()
	m2.Insert("1", 2)
	fmt.Printf("%+v\n", m2)

	foo[int](3)
	foo[string]("Hi")

}
