package lago

import (
	"fmt"
)

type Matrix struct {
	values [][]float64
	rows   int
	cols   int
}

func (m *Matrix) Size() (int, int) { return m.rows, m.cols }

func (m *Matrix) Values() *[][]float64 { return &m.values }

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

func Zeros(rows, cols int) *Matrix {
	return Create(rows, cols, 0, false)
}

func Ones(rows, cols int) *Matrix {
	return Create(rows, cols, 1, false)
}

func Id(rows int) *Matrix {
	return Create(rows, rows, 1, true)
}

func (m *Matrix) mul_scalar(scalar float64) *Matrix {
	nm := Zeros(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			nm.values[i][j] = m.values[i][j] * scalar
		}
	}

	return nm
}

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
