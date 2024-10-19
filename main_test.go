package main //ensure your tests are part of the same package, mine here is main, since my main.go contains a main package

import "testing" //import the golang testing package

//Tests NewFunc
func TestNewFunc(t *testing.T){
	NewFunc("hello world")
}

//calls MyFunc
func TestMyFunc(t *testing.T){
	MyFunc()
}