package iterator

import "fmt"

// паттерн который позволяет последовательно проходиться по элементам составных обьектов без раскрытия внутреннего представления

// Итератор (интерфейс)
type Iterator interface {
	Next() interface{}
	HasNext() bool
}

// Конкретный итератор для среза
type SliceIterator struct {
	data  []interface{}
	index int
}

func NewSliceIterator(data []interface{}) *SliceIterator {
	return &SliceIterator{data: data, index: 0}
}

func (i *SliceIterator) HasNext() bool {
	return i.index < len(i.data)
}

func (i *SliceIterator) Next() interface{} {
	if !i.HasNext() {
		return nil
	}
	value := i.data[i.index]
	i.index++
	return value
}

// Коллекция
type Collection struct {
	items []interface{}
}

func (c *Collection) Add(item interface{}) {
	c.items = append(c.items, item)
}

func (c *Collection) Iterator() Iterator {
	return NewSliceIterator(c.items)
}

// Использование
func main() {
	collection := &Collection{}
	collection.Add("первый")
	collection.Add("второй")
	collection.Add("третий")

	iterator := collection.Iterator()

	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Println(item)
	}
}
