**Q1. What is an identifier?**
A1. An Identifier is used to "name" a variable or constant, to use the parlance of the term itself, it is used to identify a variable or constant within the program. 
In the Constant exercise I created 4 identifiers (3 Constants, 1 Variable)
Constants 'hotelName', 'longitude', 'latitude' and a variable 'occupancy'

**Q2. By which type of character should an identifier begin?**
A2. Identifiers should use camelCase, so should always start with a lowercase character.

**Q3. What is the type of a variable?**
A3. The type of a variable will detremine what values can be stored in that variable, such as type string will mean any text can be used to store in the variable, type of int would mean only integers can be stored in the variable
The variable in this exercise is of type integer, type = int.

**Q4. From exercise 2: What is the default type of longitude and latitude?**
A4. longitude and latitude variables would default to float types as they have been set to decimal values, type = float64.
I proved this by using the reflect function: "fmt.Println(reflect.TypeOf(longitude))"

**Q5. When you declare a type bool variable, what is its value?**
A5. If a bool variable is declared without specifying an initial value then it's initial value will be set to false.

**Q6 What are the two main characteristics of an untyped constant?**
A6. The 2 main characteristics of an untyped constant are:
1.  Has no defined type but will use a default type implied from the value (expression) that the const has been set to.
2.  It has no limits "?" I guess this means we can set any value we like as the compiler will decide what type to set it to, depending on the inital value we set the const to.