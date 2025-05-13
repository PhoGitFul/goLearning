Answer below to show your understanding:
How to denote the type of a variable holding a pointer to a Product?
A. Not entirely sure what the question means?
Using the *star* before the type, denotes the pointer to that type of variable.
So to denote all pointers to variables of type Product, syntax would be:
*Product

What is the zero value of a pointer type?
A. nil

What does “dereferencing” mean?
A. When you derefernce a pointer this is when you retrieve the data at the address the pointer is pointing to. 

Fill in the blanks. A ____ is internally a pointer to an ____.
A. A pointer is internally a pointer to an address.

True or false. When I want to modify a map in a function, my function needs to accept a pointer to the map as parameter. I also need to return the modified map.
A. False
Maps are similar to Pointers in that they are reference types.  So when passing the map to a function you are already passing the reference to the underlying data but not the actual data itself.
So there is no need for passing a pointer to the map to the function.
Also as maps are reference types when the map is modified within the function the underlying data is modified already so no need to return the map back to the call.


Create a Go program that includes a function to modify a variable's value using a pointer. The program should:
Declare a variable of type int.
Define a function that takes a pointer to an int and modifies its value.
Call this function and print the value before and after the function call.

var poInt int                               //Declare an integer poInt

func square(ptr *int) {                     //Function which accepts a pointer and then squares the value
	*ptr = *ptr * *ptr                      //at that pointer location
}

func main() {

	poInt = 3
	fmt.Println("initial poInt:", poInt)    //Print inital value of the integer before the call to the func

	square(&poInt)                          //Pass the pointer of the integer to the Function

	fmt.Println("squared poInt:", poInt)    //Print inital value of the integer before the call to the func
	fmt.Println("poInt pointer:", &poInt)   //Print address of the integer
}