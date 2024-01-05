// To test this, copy the whole code to "sum_test.go" and overwrite it
package calc

// import "testing" pkg
import "testing"

// Unit test function must comply the following rules
//	- Test function name should start its name with the string "Test"
//	- Test function naming should be f"Test{functionName}"
//	  (The currently inspecting function name is "Sum", so "TestSum")
//  - Always gets *testing.T type
func TestSum(t *testing.T) {

	r := Sum(1, 2)
	if r != 3 {
		t.Errorf("The result is not %d", 3)
	}
}
