package matrix

import "fmt"

// Matrix is the simple implement of the matrix
type Matrix struct {
	lenx int
	leny int
	vals [][]float64
}

func (mat *Matrix) init(lx int, ly int) {
	mat.lenx = lx
	mat.leny = ly
	mat.vals = make([][]float64, lx)
	for ind := range mat.vals {
		mat.vals[ind] = make([]float64, ly)
	}
}

func (mat *Matrix) getRow(i int) Matrix {
	var nm Matrix
	nm.init(1, mat.leny)
	if 0 <= i && i < mat.lenx {
		for ind := range nm.vals {
			nm.vals[0][ind] = mat.vals[i][ind]
		}
		return nm
	}
	fmt.Println("Error index")
	return nm
}

func (mat *Matrix) getCol(j int) Matrix {
	var nm Matrix
	nm.init(mat.lenx, 1)
	if 0 <= j && j < mat.leny {
		for ind := range nm.vals[j] {
			nm.vals[ind][0] = mat.vals[ind][j]
		}
		return nm
	}
	fmt.Println("Error index")
	return nm
}

func (mat *Matrix) product(mm Matrix) Matrix {
	if mat.lenx != mm.lenx || mat.leny != mm.leny {
		fmt.Println("Error size for dot!")
		return *mat
	}
	var nm Matrix
	nm.init(mat.lenx, mat.leny)
	for i := 0; i < mat.lenx; i++ {
		for j := 0; j < mat.leny; j++ {
			nm.vals[i][j] = mat.vals[i][j] * mm.vals[i][j]
		}
	}
	return nm
}

// Vector is the vector inherit from Matrix
type Vector struct {
	Matrix
	leng float32
	norm float32
}

func (vec *Vector) getNorm() float32 {
	vec.init(3, 1)
	fmt.Println(vec.lenx)
	return 0.
}
