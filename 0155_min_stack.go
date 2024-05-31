package main

import "fmt"

type MinStack struct {
	len  int
	nums []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		len:  0,
		nums: []int{},
		mins: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if this.len == 0 {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, min(this.mins[this.len-1], val))
	}

	this.len++
	this.nums = append(this.nums, val)
}

func (this *MinStack) Pop() {
	if this.len == 0 {
		panic("invalid operation")
	}

	this.len--

	if this.len == 0 {
		this.nums = []int{}
		this.mins = []int{}
	} else {
		this.nums = this.nums[:this.len]
		this.mins = this.mins[:this.len]
	}
}

func (this *MinStack) Top() int {
	if this.len == 0 {
		panic("invalid operation")
	}

	return this.nums[this.len-1]
}

func (this *MinStack) GetMin() int {
	if this.len == 0 {
		panic("invalid operation")
	}

	return this.mins[this.len-1]
}

func main() {
	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	num := minStack.GetMin()
	fmt.Println(num)

	minStack.Pop()

	num = minStack.Top()
	fmt.Println(num)

	num = minStack.GetMin()
	fmt.Println(num)
}
