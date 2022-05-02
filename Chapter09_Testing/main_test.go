/*
	when we run	'go test' then go searches for *_test.go files (files name ended with _test.go)
	so if you want to test a function of main.go file then test case written file name must be
	main_test.go to be detected by go. As Average and TestAverage both functions are inside
	'package main' so you don't need to import function before use. 
*/
package main

import "testing"

func TestAverage(t *testing.T){
	v := Average([]int32 {1,2,3})
	expected := float32(2)

	if v != expected{
		t.Error("expected ",expected," but got ",v)
	}
}