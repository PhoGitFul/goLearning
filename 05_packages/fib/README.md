// I did originally have it passing back values to calling function but as I was just printing them out I decided not to, but leaving in so you can see

package main

import ("fmt"
    "golang.org/x/text/language"
    "golang.org/x/text/message"
)

// Fib.go
// External function

func Fibs(value int) (fib, fibnum int) { 
	var a [51] int
	a [0] = 0
	a [1] = 1
		if value <  2 { 
			return a[value], 1
			}
			if value > 50 {
				fmt.Println("Cannot handle more than 50, so 50th fibonacci number is:")
	p := message.NewPrinter(language.English)
    p.Printf("%d\n", 12586269025)
	return 12586269025, 50
				}
	for i := 2; i <= value; i++ {
	    a[i] = a[i-2] + a[i-1]
	}
    fmt.Println("Fibonacci Number",value,"is :")
	p := message.NewPrinter(language.English)
    p.Printf("%d\n", a[value])
	return a[value], value
}

func main() {
	val, fibnum := Fibs(10)
    fmt.Println("Fib Num",fibnum,":")
    fmt.Println("Fibonacci Number",fibnum,"is :")
	p := message.NewPrinter(language.English)
    p.Printf("%d\n", val)
}


Also the main issue I had was the problem around an error I had "No required module provides package"
After a lot of searching it seemed to relate to the mod file, and I realised the initial mod file was related to the first folder we coded in our goLearning so it contained "module 01_helloworld"
As I was not coding in that folder, I changed the name to the directory above so "module goLearning" and the issue went away :-)
I did then have other issues as I did not precede my function call with the package name so I had the call from main as Fibs(3), when I needed to have it preceded by the package "fib" such as => fib.Fibs(3).

I was fiddling around with the name of the 05_Package folder and this also caused issues with seemingly duplicate declarations, so I made sure all used lower case 05_packages anywhere in my code, I found the import statement in main was referring to 
import "goLearning/05_Packages/fib", so I changed this to be lowercase and the warning disappeared.

Oh and found that the packages I imported and used in my fib package needed to be "got", so I simply ran go mod tidy (I had tried this many times earlier to fix my missing package issue above but this always did nothing as it was looking only in the 01_hellowworld directory) and this time as the mod file ahd been amended to refer to the root goLearning directory the mod tidy comand pulled in the relevant code:

mal@ip-10-192-240-10 goLearning % go mod tidy                                                            
go: finding module for package golang.org/x/text/message
go: finding module for package golang.org/x/text/language
go: found golang.org/x/text/language in golang.org/x/text v0.19.0
go: found golang.org/x/text/message in golang.org/x/text v0.19.0

go.mod now looks like this:

module goLearning

go 1.21.4

require golang.org/x/text v0.19.0
