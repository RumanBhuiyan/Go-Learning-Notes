# MY Go Learning Notes
> ## Hello World
```go
      // differences between Println() and Printf()
      // Println() prints output then create newline but Printf() doesn't
      package main

      import "fmt"

      func main(){

        fmt.Println("Hello World")
        fmt.Printf("Good")
        fmt.Printf(" Morning")
      }
```
> ## Variables 
```go
      // Go supports both statically typed and dynamic typed variables
      package main

      import "fmt"

      func main(){
        // Statically Typed variable
        var name string = "go" 
        fmt.Println(name,"is an awesome language ")

        // dynamic typed variable
        var y = 5
        fmt.Printf("type of y  %T\n",y)
        fmt.Println("Value of y ",y)

        // multiple variable declarations in one line 
        var a,b,c = 5 ,4, 3
        fmt.Println("a = ",a)
        fmt.Println("b = ",b)
        fmt.Println("c = ",c)

        // declaring constant
        const PI float32 = 3.1416
        var radius float32 = 2.5
        fmt.Println("Area of the circle : ",PI*radius*radius)  

      }
```
> ## Taking User Input
```go
      package main

      import "fmt"

      func main(){
        // process: 01 reading int number 
        var first_number int 
        fmt.Printf("Enter the first number: ")
        fmt.Scanf("%d",&first_number)

        // process: 02 reading int number 
        var second_number int 
        fmt.Printf("Enter Second number: ")
        fmt.Scan(&second_number)

        fmt.Println("summation : ",first_number+second_number)

        // reading string
        var country_name string 
        fmt.Printf("Enter country name : ")
        fmt.Scan(&country_name)
        fmt.Println("Your country is ",country_name)
      }
```
> ## Conditional Statements
```go
      package main 

      import "fmt"

      func main(){
        // if - else if -else ladder 
        var number int 
        fmt.Printf("Enter the number: ")
        fmt.Scan(&number)

        if (number == 0){
          fmt.Println("Number is zero")
        }else if (number > 0){
          fmt.Println("Number is Positive")
        }else {
          fmt.Println("Number is Negative")
        }

        // switch statement 
        var another_number int = 2
        switch(another_number){
        case 0 :
          fmt.Println("zero")
          break 
        case 1 :
          fmt.Println("One")
          break 
        default :
            fmt.Println("Undefined")
        }
      }
```
> ## Looping 
```go
      package main 

      import "fmt"

      func main(){
        // for loop 
        for i := 0; i < 5; i++ {
          fmt.Printf("i = %d\n",i)
        }

        // while type for loop 
        var a , b = 1,5
        for a <= b {
          fmt.Println("a = ",a)
          a++ 
        }
      }
```
> ## Functions
```go
      package main

      import "fmt"

      func main(){

        var number1, number2 float64
        fmt.Printf("Enter the first number: ")
        fmt.Scan(&number1)
        fmt.Printf("Enter the second number: ")
        fmt.Scan(&number2)

        fmt.Println("Summation: ",sum(int(number1),int(number2)))

        var product , quotient float64 = divisionResult(number1,number2)

        fmt.Println("Product :",product)
        fmt.Println("Quotient : ",quotient)
      }

      // function of returning one statement
      func sum(a,b int ) int {
         return a+b 
      }
      // function of returning two statement
      func divisionResult(a,b float64) (float64, float64) {
        return a*b , a / b ;
      }
```
> ## Local Scope VS Global Scope
```go
      package main

      import "fmt"

      var PI float32 = 3.1416 
      var age int = 23

      func main(){
          var radius float32 ;
        fmt.Printf("Enter the radius: ")
        fmt.Scan(&radius)
        fmt.Println("Radius of the circle ",getArea(radius))

        // before creating same name variable of global scope
        fmt.Println("Global age: ",age)
        // after creating same name variable of global scope
        var age int =25
        fmt.Println("Local age: ",age)
      }

      func getArea(radius float32) float32 {
        // programmer didn't need to pass PI here for global scope
        return PI*radius*radius 
      }
```
> ## Strings
```go
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
```
> ## Arrays
```go
      package main

      import "fmt"

      func main(){
        // declaring arrays
         var numbers [5]int 
         var another_set =[5] int {1,2,3,4,5}

         printElements(another_set,5)

         // changing items of numbers array
         var i int 
         for i = 0 ; i < len(numbers) ; i++ {
           numbers[i] = (i+1)*5
         }  

         printElements(numbers,len(numbers))
      }

      // decalaring a void function
      func printElements(numbers [5]int,size int){
        var i int 
        for i = 0 ; i < size ; i++ {
            fmt.Println(i+1,"element: ",numbers[i])
        } 
      }

```
> ## Pointers 
```go
      package main

      import "fmt"

      func main(){
        var pointer *int 

        var keep int = 5;
        pointer = &keep  

        fmt.Println("Address: ",pointer,"value : ",*pointer)
      }
```
> ## Structures 
```go
      package main

      import "fmt"

      type Info struct{
         name string
         age  int  
         occupation string
      }

      func main(){

        var first_person Info ;

        // assigning values
        first_person.name = "Ruman"
        first_person.age  = 24
        first_person.occupation = "student"

        // accessing items 
         fmt.Println("name: ",first_person.name)
         fmt.Println("age: ",first_person.age)
         fmt.Println("occupation: ",first_person.occupation)

      }
```
> ## Slice 
```go
      // slice allow programmer to create array which is resizable
      package main

      import "fmt"

      func main(){
        // create an array which current size 3 but capacity 5
         var numbers = make([]int, 3, 5)

         numbers[0] = 5
         numbers[1] = 10

         fmt.Println("Length ",len(numbers))
         fmt.Println("capacity ",cap(numbers))

         // sub-array or slice 
         fmt.Println("first 2 elements ",numbers[0:2])

         // append function 
         numbers = append(numbers,3)

         fmt.Println(numbers)
      }
```
> ## Range 
```go
      package main

      import "fmt"

      func main(){
        var numbers  = [5]int{5,10,15,20,25}
          var i int 
        for i = range numbers {
          fmt.Println(i+1,"element : ",numbers[i])
        }// output : 
        // 1 element :  5
        // 2 element :  10
        // 3 element :  15
        // 4 element :  20
        // 5 element :  25
      }
```
> ## Map
```go
      package main

      import "fmt"

      func main(){
        // declaring maps 
        var keys map[string]int
        // creating map
          keys  = make(map[string]int)

        keys["apple"]= 1
        keys["mango"]= 2
        keys["orange"]= 3
        keys["jam"]= 4
        keys["grape"]= 5

        // before deleting 
        showMapItems(keys)
           fmt.Println("------------")
        // after deleting item 
          delete(keys,"grape")
          showMapItems(keys)
      }
      func showMapItems(myMap map[string]int){
          var i string ;
        for i = range myMap {
          fmt.Println(i," : ",myMap[i])
        }
      }
```
> ## Recursion
```go
      package main

      import "fmt"

      func main(){

        var number int
        fmt.Printf("Enter the number: ")
        fmt.Scan(&number)
        fmt.Println("factorial of the number is ",getFactorial(number))
      }

      func getFactorial(number int) int {
         if (number == 1)  {
           return 1 
        }
         return number * getFactorial(number -1 )
      }
```
> ## Handling Errors
```go
      package main

      import "fmt"
      import "errors"

      func main(){
        var result , error = getResult(5,0)

        if (error == nil){
          fmt.Println(result)
        }else {
          fmt.Println(error)
        }
      }

      func getResult(a,b int) (float32,error) {
          if ( b == 0){
          return 0,errors.New("can't divide by zero")
        }
        // type casting first 
        return float32(a)/float32(b) ,nil;
      }
```
> ## Interface
```go
      package main

      import "fmt"
      import "math" 

      /* define an interface */
      type Shape interface {
         area() float64
      }

      /* define a circle */
      type Circle struct {
         x,y,radius float64
      }

      /* define a rectangle */
      type Rectangle struct {
         width, height float64
      }

      /* define a method for circle (implementation of Shape.area())*/
      func(circle Circle) area() float64 {
         return math.Pi * circle.radius * circle.radius
      }

      /* define a method for rectangle (implementation of Shape.area())*/
      func(rect Rectangle) area() float64 {
         return rect.width * rect.height
      }

      /* define a method for shape */
      func getArea(shape Shape) float64 {
         return shape.area()
      }

      func main() {
         circle := Circle{x:0,y:0,radius:5}
         rectangle := Rectangle {width:10, height:5}

         fmt.Printf("Circle area: %f\n",getArea(circle))
         fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
      }
```
