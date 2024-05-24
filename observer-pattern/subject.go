package observer

import "fmt"

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailibilty() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.NotifyAll()
}

func (i *Item) Register(observer Observer) {
	i.observerList = append(i.observerList, observer)
}

func (i *Item) NotifyAll() {
	for _, o := range i.observerList {
		o.update(i.name)
	}
}
