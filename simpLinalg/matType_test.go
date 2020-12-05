package matrix

import "testing"

// TestMatrix is the test function for Matrix class
func TestMatrix(t *testing.T) {
	var mat Matrix
	mat.init(2, 2)
	if mat.lenx != 2 {
		t.Errorf("Matrix.init make a wrong x_dim!")
	}
}
