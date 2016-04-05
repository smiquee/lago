// Package lago aims at being a very simple linear algebra library.
// It is a playground for discivering golang and playing with linear
// algebra with the idea to be used in a future machine learning playground
// project.
package lago

import (
	"fmt"
)

// The Matrix structure with the internal values (as a 2D array) and
// the rows and cols numbers describing the Matrix dimensions.
// The stored values are only of float64 type.
type Matrix struct {
	values [][]float64
	rows   int
	cols   int
}

// Method Size returns the rows and cols numbers of a given Matrix.
func (m *Matrix) Size() (int, int) { return m.rows, m.cols }

// Method Values returns the 2D array values of a given Matrix.
func (m *Matrix) Values() *[][]float64 { return &m.values }

// Method Get returns the value in x and y position in the given Matrix.
func (m *Matrix) Get(x, y int) float64 { return m.values[x][y] }

// Function Create returns a Matrix of size rows and cols, with an
// init value. Also, the id parameter indicates whether to create
// an identity matrix or not.
func Create(rows, cols int, init float64, id bool) *Matrix {
	if id {
		if rows != cols {
			panic("Can only create square identity matrix!")
		}
	}
	values := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		values[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			if id {
				if i == j {
					values[i][j] = 1
				}
			} else {
				values[i][j] = init
			}
		}
	}

	m := Matrix{values, rows, cols}

	return &m
}

// Function Zeros returns a Matrix of size rows and cols with
// an init value of 0.
func Zeros(rows, cols int) *Matrix {
	return Create(rows, cols, 0, false)
}

// Function Zeros returns a Matrix of size rows and cols with
// an init value of 1.
func Ones(rows, cols int) *Matrix {
	return Create(rows, cols, 1, false)
}

// Function that return an identity Matrix of size rows.
func Id(rows int) *Matrix {
	return Create(rows, rows, 1, true)
}

// Internal mul_scalar returns a new Matrix which is the result
// of the multiplication of the given Matrix by the given scalar.
func (m *Matrix) mul_scalar(scalar float64) *Matrix {
	nm := Zeros(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			nm.values[i][j] = m.values[i][j] * scalar
		}
	}

	return nm
}

// Internal mul_matrix returns a new Matrix which is the result
// of the multiplication of the given Matrix by another Matrix
// given as parameter.
func (m0 *Matrix) mul_matrix(m1 *Matrix) *Matrix {
	if m0.cols != m1.rows {
		panic("Dimensions mismatch!")
	}

	nm := Zeros(m0.rows, m1.cols)
	for i := 0; i < m0.rows; i++ {
		for j := 0; j < m1.cols; j++ {
			for k := 0; k < m0.cols; k++ {
				nm.values[i][j] += m0.values[i][k] * m1.values[k][j]
			}
		}
	}

	return nm
}

// Method Mul is a generic method that returns a new Matrix which is
// the result of the multiplication of the given Matrix with an interface.
// This interface can either be a scalar (float64 or int) or another Matrix.
func (m *Matrix) Mul(mul interface{}) *Matrix {
	switch mul.(type) {
	case float64:
		return m.mul_scalar(mul.(float64))
	case int:
		return m.mul_scalar(float64(mul.(int)))
	case *Matrix, Matrix:
		r, c := mul.(*Matrix).Size()
		if c == 1 && r == 1 {
			return m.mul_scalar(mul.(*Matrix).values[0][0])
		}
		return m.mul_matrix(mul.(*Matrix))
	}
	return &Matrix{}
}

// Method Transpose returns a new Matrix which is the transposed version
// of the given Matrix.
func (m *Matrix) Transpose() *Matrix {
	values := make([][]float64, m.cols)
	for i := 0; i < m.cols; i++ {
		values[i] = make([]float64, m.rows)
		for j := 0; j < m.rows; j++ {
			values[i][j] = m.values[j][i]
		}
	}

	nm := Matrix{values, m.cols, m.rows}

	return &nm
}

// Method Print does a user friendly print of the given Matrix.
func (m *Matrix) Print() {
	i := 0
	for ; i < m.rows; i++ {
		j := 0
		for ; j < m.cols-1; j++ {
			fmt.Print(m.values[i][j])
			fmt.Print(" ")
		}
		fmt.Println(m.values[i][j])
	}
}
