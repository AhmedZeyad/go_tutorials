package main

import (
	"errors"
	"fmt"

	"unicode/utf8"
)

func main() {
	// arrays()
	// slices()
	// maps()
	// MyStructs()
// algoRuner()
data()

}










































func ifStatment() {
	x, y := 33, 2
	var resoult, err = indDivision(x, y)
	if err != nil {
		fmt.Printf(" you can't divide on 0 so The resoult is %v good job\n", resoult)
	} else if resoult == 1 {
		fmt.Println("mashalla good job\n")

	} else {

		fmt.Println(resoult, err)
	}
	switch {
	case err != nil:
		fmt.Println("you cna't doit pro go back\n")
	default:
		fmt.Printf("wellcom  pro the resoultis %v and every thing is good\n", resoult)
	}
}
func varKnowing() {
	fmt.Println("hello momo")
	var intNum int32 = 429496729
	fmt.Println(intNum)

	var floatNum float32
	fmt.Print(floatNum)
	var intNum32 int32 = 2
	var floatNum32 float32 = 3.333
	var resoult float32 = float32(intNum32) + floatNum32
	fmt.Println(resoult)
	var x int32 = 2
	var y int32 = 2
	fmt.Println(x / y)
	fmt.Println(x % y)
	var name string = "ahmed¥†"
	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println(utf8.RuneCountInString(name))
	var myRune rune = 'A'
	fmt.Println("\n\n\n\nnow line \n\n\n")
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)
	ahmed := 1
	var momo int
	fmt.Println(momo)
	fmt.Println(ahmed)
}
func indDivision(numerator int, denominator int) (int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("شكد عبقري دخيل ربك")
		return 0, err
	} else {

		return numerator / denominator, err
	}
}
func arrays() {
	var myArray [3]int32
	myArray[1] = 313
	fmt.Printf("salam alla ala abo salih %v\n", myArray[1])
	fmt.Printf("to show the memory locaion you cna ad &brfore var name like &myArray[0] :%v\n", &myArray[0])
	var mySecondArray [2]int = [2]int{2, 3}
	var mySecondArray2 = [2]int{2, 3}
	mySecondArray3 := [2]int{2, 3}
	println(mySecondArray[0])
	println(mySecondArray2[0])
	println(mySecondArray3[0])
}
func slices() {
	intSlice := []int{2, 4}
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", intSlice, len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 9)
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", intSlice, len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 99999)
	intSlice = append(intSlice, 9)
	intSlice = append(intSlice, 9)
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", intSlice, len(intSlice), cap(intSlice))
	var mySlice []int = make([]int, 0)
	var mySlice2 []int = []int{2, 3, 4}

	fmt.Println(mySlice)
	mySlice = append(mySlice, 1)
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", mySlice, len(mySlice), cap(mySlice))
	mySlice = append(mySlice, 2)
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", mySlice, len(mySlice), cap(mySlice))
	mySlice = append(mySlice, 3)
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", mySlice, len(mySlice), cap(mySlice))
	mySlice = append(mySlice, 4)
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", mySlice, len(mySlice), cap(mySlice))
	mySlice = append(mySlice, 5)
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", mySlice, len(mySlice), cap(mySlice))
	mySlice = append(mySlice, 6)
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", mySlice, len(mySlice), cap(mySlice))
	mySlice = append(mySlice, intSlice...)
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", mySlice, len(mySlice), cap(mySlice))
	mySlice2 = append(mySlice2, 3)
	fmt.Printf("%vslicslent is %v and the capasty is %v\n", mySlice2, len(mySlice2), cap(mySlice2))

}
func maps() {
	var strMap map[string]int = map[string]int{"ahmed": 22, "joje": 20}
	var value,ok = strMap["ali"]
	if ok{
		fmt.Printf("hi bay %v\n",value)
	}else{
		fmt.Printf("gg man %v\n",ok)
	}
	for name :=range strMap{
fmt.Printf("Name is %v\n",name)
	}
}