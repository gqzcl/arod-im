package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	con := Constructor()
	con.AddNum(1)
	con.AddNum(2)
	fmt.Println(con.FindMedian())
	con.AddNum(3)
	fmt.Println(con.FindMedian())
}

func sumOddLengthSubarrays(arr []int) int {
	res, lens := 0, len(arr)
	for i := 1; i < lens; i++ {
		arr[i] += arr[i-1]
	}
	for i := 1; i <= lens; i += 2 {
		a, b := 0, i
		res += arr[b-1]
		for b < lens {
			res += arr[b] - arr[a]
			a++
			b++
		}
	}
	return res
}

type MedianFinder struct {
	queMin, queMax hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}
func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(maxQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}
func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
