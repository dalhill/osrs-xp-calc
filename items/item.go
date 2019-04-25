package items

import "fmt"

type Item struct {
	Name string
	Count int
}

type ItemSlice []Item

func (itemSlice ItemSlice) GetPointersFor(sub ItemSlice) []*Item {
	var pointers []*Item
	for _, i := range itemSlice {
		if sub.Contains(i) {
			fmt.Println("Adding: ", i.Name, &i)
			pointers = append(pointers, &i)
			fmt.Println("Current pointers: ", pointers)
		}
	}
	return pointers
}

func (itemSlice ItemSlice) Contains(item Item) bool {
	for _, i := range itemSlice {
		if i.Name == item.Name {
			return true
		}
	}
	return false
}

func IndexMap(a ItemSlice, b ItemSlice) map[int]int {
	indexMap := map[int]int{}
	for i := range a {
		for j := range b {
			if a[i].Name == b[j].Name {
				indexMap[i] = j
			}
		}
	}
	return indexMap
}