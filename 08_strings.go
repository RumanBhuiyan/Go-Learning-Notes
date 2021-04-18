package main

import ("fmt" )

func main(){
	var country string = "Bangladesh" 

	fmt.Println("Length of the String : ",len(country))
	fmt.Printf("country name:  %s\n",country)

	// iterating over strings 
	for i := 0; i< len(country); i++ {
		fmt.Printf("%d character: %c\n",i+1,country[i])
	}

}