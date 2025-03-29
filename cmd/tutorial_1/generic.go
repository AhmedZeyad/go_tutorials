package main

import "fmt"

func data() {
	
	var myISlice = []int{1, 2, 3}
	fmt.Printf("the sum is :%v\n",sumSlice(myISlice))
	var myFSlice = []float32{1, 2, 3}
	fmt.Printf("the sum is :%v\n ",sumSlice(myFSlice))
	var myF64Slice = []float64{1, 2, 3}
	fmt.Printf("the sum is :%v\n %T",sumSlice(myF64Slice),sumSlice(myF64Slice))
}
func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
	// generic in stract
}
	type Pcc[T Cpu|Gpu] struct{
		name string
		processor T
		ram int
		storage int   
	
	}
var  aiPc=Pcc[Gpu]{"colab",Gpu{"t4",16,16,32},16,512}