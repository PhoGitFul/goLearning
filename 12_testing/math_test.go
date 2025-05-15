// math_test.go
package math

import "testing"

func TestOne(testVar *testing.T) {
	//AAA

	// Arrange
	input1 := 8
	input2 := 3
	expectedResult := 11

	// Act
	actualResult := Add(input1, input2)

	// Assert
	if actualResult != expectedResult {
		testVar.Fail()
	}
}
