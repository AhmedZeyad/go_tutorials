package main

import "fmt"

type Algo interface {
	encreption()
	decreption()
}
type Caesar struct {
	key int
	msg string
}

func (c Caesar) encreption() {
	fmt.Printf("caeser doen encreption the key is %v\n", c.key)
}
func (c Caesar) decreption() {
	fmt.Printf("caeser done decreption the key is %v\n", c.key)
}

type Vigenere struct {
	key  int
	sKey int
	msg  string
}

func (v Vigenere) encreption() {
	fmt.Printf("caeser doen encreption the key is %v %v", v.key, v.sKey)
}
func (v Vigenere) decreption() {
	fmt.Printf("visenar done decreption the key is %v %v", v.key, v.sKey)
}

func Encrepter(a Algo) {
	a.encreption()
}
func algoRuner() {
	var caeser Caesar = Caesar{12, "msg"}
	var vessiner Vigenere = Vigenere{22, 33, "msg2"}
	Encrepter(caeser)
	Encrepter(vessiner)
}
