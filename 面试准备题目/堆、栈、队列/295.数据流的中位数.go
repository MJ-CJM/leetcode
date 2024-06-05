package main

import "container/heap"

type MinHeap []int

func(hp MinHeap) Len() int {return len(hp)}
func(hp MinHeap) Less(i, j int)bool {return hp[i]<hp[j]}
func(hp MinHeap) Swap(i, j int) {hp[i],hp[j] = hp[j], hp[i]}

func(hp *MinHeap) Push(x interface{} ) {
	*hp = append(*hp,x.(int))
}

func(hp *MinHeap) Pop() interface{} {
	n := len(*hp)
	x := (*hp)[n-1]
	*hp = (*hp)[:n-1]
	return x
}

type MaxHeap []int

func(hp MaxHeap) Len() int {return len(hp)}
func(hp MaxHeap) Less(i, j int)bool {return hp[i]>hp[j]}
func(hp MaxHeap) Swap(i, j int) {hp[i],hp[j] = hp[j], hp[i]}

func(hp *MaxHeap) Push(x interface{} ) {
	*hp = append(*hp,x.(int))
}

func(hp *MaxHeap) Pop() interface{} {
	n := len(*hp)
	x := (*hp)[n-1]
	*hp = (*hp)[:n-1]
	return x
}

type MedianFinder struct {
	//维护后半个
	minhp *MinHeap
	//维护前半个
	maxhp *MaxHeap
}


func Constructor() MedianFinder {
	return MedianFinder{&MinHeap{},&MaxHeap{}}
}


func (this *MedianFinder) AddNum(num int)  {
	//如果是空的
	if this.minhp.Len()==0{
		heap.Push(this.minhp,num)
		return
	}
	//如果不为空
	//后半段的更多
	if this.minhp.Len()>this.maxhp.Len(){
		heap.Push(this.minhp,num)
		heap.Push(this.maxhp,heap.Pop(this.minhp).(int))
	}else{
		//一样多
		heap.Push(this.maxhp,num)
		heap.Push(this.minhp,heap.Pop(this.maxhp).(int))
	}
	// fmt.Println(this.maxhp)
	// fmt.Println(this.minhp)
	// fmt.Println()
}


func (this *MedianFinder) FindMedian() float64 {
	if this.minhp.Len()>this.maxhp.Len(){
		return float64((*this.minhp)[0])
	}else{
		return float64((*this.minhp)[0]+(*this.maxhp)[0])/2
	}
}

