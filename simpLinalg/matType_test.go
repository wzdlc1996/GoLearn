package matrix

import "testing"

// TestMatrix is the test function for Matrix class
func TestMatrix(t *testing.T) {
	var mat Matrix
	mat.init(2, 2)
	var vec Vector
	vec.init(2, 2)
	if vec.lenx != 2 {
		t.Errorf("Vector.init makes a wrong x_dim!")
	}
	if mat.lenx != 2 {
		t.Errorf("Matrix.init makes a wrong x_dim!")
	}
}
