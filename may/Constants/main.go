package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	var e int
	fmt.Println("e =", e)

	var f float32 = 58888888889999
	g := 999999999999999999999i
	h := 5.78
	var i float64 = 5.7
	var j complex64 = -3.14i

	fmt.Println("f =", f)
	fmt.Println("f Type =")
	fmt.Printf("%T", f)
	fmt.Println()
	fmt.Println("g Type =", reflect.TypeOf(g))
	fmt.Println("h Type =", reflect.TypeOf(h))
	fmt.Println("i Type =", reflect.TypeOf(i))
	fmt.Println("j Type =", reflect.TypeOf(j))

	// Using 8-bit unsigned int
	var X uint8 = 225
	fmt.Println(X, X-3)

	// Using 16-bit signed int
	var Y int16 = 32767
	var negY int16 = -32767
	fmt.Println("negY =", negY, ",", "Y =", Y)
	fmt.Fprintln(os.Stdout, []any{"Y+negY =", Y + negY, ",", "negY-Y =", negY - Y, ",", "-32767 - 32767 => -32768 - 32766"}...)
	fmt.Println("sort of flips to positive as can't get any lower than -32768 so is like +32768 - 32766 = 2")
	fmt.Println("Y+0 =", Y+0, ",", "Y+1 =", Y+1, ",", "Y+2 =", Y+2, ",", "Y-1 =", Y-1, ",", "Y-2 =", Y-2)
	fmt.Println("negY+0 =", negY+0, ",", "negY+1 =", negY+1, ",", "negY+2 =", negY+2, ",", "negY-1 =", negY-1, "negY-2 =", negY-2, negY-1)
	// 225 222
	// -32767 32765
}
