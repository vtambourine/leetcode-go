package main

import (
	"container/list"
	"fmt"
)

type FirstUnique struct {
	index       map[int]int
	queue       *list.List
	firstUnique *list.Element
}

func Constructor(nums []int) FirstUnique {
	queue := FirstUnique{
		index: make(map[int]int),
		queue: &list.List{},
	}

	for _, n := range nums {
		queue.Add(n)
	}

	return queue
}

func (this *FirstUnique) ShowFirstUnique() int {
	if this.firstUnique == nil {
		return -1
	}
	return this.firstUnique.Value.(int)
}

func (this *FirstUnique) Add(value int) {
	this.queue.PushBack(value)
	this.index[value] += 1

	node := this.firstUnique
	if node == nil {
		node = this.queue.Front()
	}
	for node != nil {
		if this.index[node.Value.(int)] == 1 {
			this.firstUnique = node
			break
		}
		this.firstUnique = nil
		node = node.Next()
	}
}

func (this *FirstUnique) Print() {
	node := this.queue.Front()
	for node != nil {
		fmt.Printf("%v-", node.Value)
		node = node.Next()
	}
	fmt.Println()
	fmt.Println(this.index)
}

func main() {
	fmt.Println("First Unique Number")
	//
	//obj := Constructor([]int{2, 3, 5})
	//fmt.Println(obj.ShowFirstUnique())
	//obj.Add(5)
	//fmt.Println(obj.ShowFirstUnique())
	//obj.Add(2)
	//fmt.Println(obj.ShowFirstUnique())
	//obj.Add(3)

	//obj := Constructor([]int{7,7,7,7,7,7})
	//fmt.Println(obj.ShowFirstUnique())
	//obj.Add(7)
	//obj.Add(3)
	//obj.Add(3)
	//obj.Add(7)
	//obj.Add(17)
	//fmt.Println(obj.ShowFirstUnique())

	obj := Constructor([]int{809})
	fmt.Println(obj.ShowFirstUnique())
	obj.Add(809)
	fmt.Println(obj.ShowFirstUnique())

	obj.Print()
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
