// To test this, copy the whole code to "sum_test.go" and overwrite it
package calc

import "testing"

type TestData struct {
	argument1 int
	argument2 int
	result    int
}

var testDataset = []TestData{
	{2, 6, 8},
	{-8, 3, -5},
	{0, 0, 0},
	{1, 2, 3},
	{8, 8, 16},
}

func TestSumBulk(t *testing.T) {
	for _, data := range testDataset {
		testingFunctionResult := Sum(data.argument1, data.argument2)
		if testingFunctionResult != data.result {
			t.Errorf("Expected: %d + %d = %d, but received %d\n",
				data.argument1, data.argument2, data.result, testingFunctionResult)
		}
	}
}
