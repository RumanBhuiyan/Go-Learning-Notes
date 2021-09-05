//  package is a container of many go files
//  package are 2 types 
// i) Executable (.exe file is generated on it) 
// ii) Reusable(helper functions inside package) 
// due to package main, main.exe is generated at build time
package main

// fmt is short for format 
// import fmt statemnt links that package to my package main
// so that i can use built in functions of that package
import "fmt"

func main(){

	fmt.Println("Hello Ruman")

}

// running go code in terminal 

// process-01 : go run 01_flow.go
// this will compile and execute go file

// process-02 : go build
// this will build my .exe file named 01_flow.exe then run 01_flow.exe
// and output will arrive

