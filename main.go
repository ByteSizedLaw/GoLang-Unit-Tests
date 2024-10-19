package main

import "fmt"

var mySlice = []string{"Hey","There","Delilah","!"}

func main(){
	MyFunc()
}

//public function
func MyFunc(){
	var x = len(mySlice)
	for i := 0; i < 10; i++ {
		if i<x {
			fmt.Println(mySlice[i])
		}
		if i>x {
			fmt.Println("Nothing left to output")
		}
	}
}

//public function
func NewFunc(abc string){
	fmt.Println(abc)
	fmt.Println("hello ", 146, "300")
}