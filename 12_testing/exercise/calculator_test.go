// calculator_test.go
package math

import "testing"

var input1, input2 = 8, 3

func TestAdd(testAddVar *testing.T) {
	//AAA

	// Arrange
	//	input1 := 8
	//	input2 := 3
	expectedResult := 14

	// Act
	actualResult := Add(input1, input2)

	// Assert
	if actualResult != expectedResult {
		testAddVar.Errorf("%d + %d = %d; expected %d", input1, input2, actualResult, expectedResult)
		testAddVar.Fail()
	}
}

func TestSubtract(testSubVar *testing.T) {
	//AAA

	// Arrange
	//	input1 := 8
	//	input2 := 3
	expectedResult := 5

	// Act
	actualResult := Subtract(input1, input2)

	// Assert
	if actualResult != expectedResult {
		testSubVar.Fail()
	}
}

func TestMult(testMultVar *testing.T) {
	//AAA

	// Arrange
	//	input1 := 8
	//	input2 := 3
	expectedResult := 24

	// Act
	actualResult := Multiply(input1, input2)

	// Assert
	if actualResult != expectedResult {
		testMultVar.Fail()
	}
}

var tests = []struct {
	wantAdd, wantSub, wantMult, input1, input2 int
}{
	{0, 0, 0, 0, 0},
	{1, -1, 0, 0, 1},
	{2, 0, 1, 1, 1},
	{0, -2, -1, -1, 1},
	{-1, -1, 0, -1, 0},
	{2, 0, 1, 1, 1},
	{321, 321, 1, 320, 1},
	{44, 0, 484, 22, 21},
	{43, 1, 462, 22, 21},
}

func TestValues(t *testing.T) {
	for _, test := range tests {
		if got := Add(test.input1, test.input2); got != test.wantAdd {
			t.Errorf("sum of %d and %d should be %d not %d", test.input1, test.input2, test.wantAdd, got)
		}
		if got := Subtract(test.input1, test.input2); got != test.wantSub {
			t.Errorf("%d minus %d should be %d not %d", test.input1, test.input2, test.wantSub, got)
		}
		if got := Multiply(test.input1, test.input2); got != test.wantMult {
			t.Errorf("multiplication of %d and %d should be %d not %d", test.input1, test.input2, test.wantMult, got)
		}
	}
}
