// routine = asynchronous programming
package main

import (
	"fmt"
	"net/http"
)

func main(){

	links :=  []string {
		"http://google.com",
		"http://facebook.com",
		"http://github.com",
	}

	//printLinkStatus(links)
	 //printLinkStatus2(links) ;

	 // lets create our channel . you can create only inside main function
	 channel := make( chan string)
	 // i am creating child routine and linking it to channel
	 go printLinkStatus2(links , channel)

	 // i am telling channel to wait for receiving message for all of links
	 // whenever you will get all responses then close connection
	 for i :=0; i< len(links) ; i++{
		fmt.Println(<-channel)
	 } 
}

// this function is executing synchronously which 
// blocks next task to be executed before it finishes thats why
// response of servers follows the order of array 
func printLinkStatus(allLinks []string)  {
	for _,link := range allLinks{
		_,error := http.Get(link)
		if(error != nil){
			fmt.Println("Error : ",error)
		}else {
			fmt.Println(link,"is up")
		}
	}
}

// Go routines are 2 type. i) Main routine ii) child routine
// if main routine is finished then it doesn't take care of child routine 
//  that why we create a channel which links up with main routine and all other
// child routines before executing program Main routine ask Channel is there anyting
// to be finished. if there is none only then program finishes
func  printLinkStatus2(allLinks []string,channel chan string)  {

	// for sleeping : time.Sleep(5*time.Second)

	for _,link := range allLinks{
		_,error := http.Get(link)
		
		if(error != nil){
			fmt.Println("Error : ",error)
			channel <- "Error"
		}else {
			fmt.Println(link,"is up")
			channel <- "Status Ok"
		}
	}
}



