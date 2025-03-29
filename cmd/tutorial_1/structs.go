package main

import "fmt"
// func main(){}
type Pc struct{
	name string
	Cpu
	Gpu

	ram int
	storage int   

}
type Cpu struct{
	name string
	gen int
	core int

}
type Gpu struct{
	name string
	ram int
	core int
	speed float32

}
func (g Gpu) getData()string{
	return "from GPU"
}
func(c Cpu) getData()string{
	return"from CPU"
}
type Printing interface{}
type ProccessingUnit interface{
getData() string

}

func MyStructs() {
var ahmedMac Pc =Pc{"ahmed mac",Cpu{"M1",1,7},Gpu{"M1",1,8,60},8,256}
fmt.Printf("your gpu is %v has %v ram and%v core you pc storage is %v",ahmedMac.Gpu.name,ahmedMac.Gpu.ram,ahmedMac.Gpu.core,ahmedMac.storage)
}
