1. An Identifier is used to "name" a variable or constant, to use the parlance of the term itself, it is used to identify a variable or constant within the program. 
In the Constant exercise I created 4 identifiers (3 Constants, 1 Variable)
Constants 'hotelName', 'longitude', 'latitude' and a variable 'occupancy'

2. Identifiers should use camelCase, so should always start with a lowercase character.

3. The variable in the exercise is of type integer, type = int.

4. longitude and latitude variables would default to float types as they have been set to decimal values, type = flaot64.
I proved this by using the reflect function: "fmt.Println(reflect.TypeOf(longitude))"

5. If a bool variable is declared without specifying an initial value then it's initial value will be set to false

6. The 2 main characteristics of an untyped constant are:
a.  Has no defined type but will use a default type implied from the value (expression) that the const has been set to.
b.  It has no limits "?" I guess this means we can set any value we like as the compiler will decide what type to set it to, depending on the inital value we set the const to.