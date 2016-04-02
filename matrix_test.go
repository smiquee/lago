package lago_test

import (
	"github.com/smiquee/lago"
	"testing"
)

func TestOnes(t *testing.T) {
	m := lago.Ones(2, 2)
	rows, cols := m.Size()
	if rows != 2 || cols != 2 {
		t.FailNow()
	}

	val := m.Values()
	if (*val)[0][0] != 1 {
		t.FailNow()
	}
	if (*val)[0][1] != 1 {
		t.FailNow()
	}
	if (*val)[1][0] != 1 {
		t.FailNow()
	}
	if (*val)[1][1] != 1 {
		t.FailNow()
	}
}

func TestZeros(t *testing.T) {
	m := lago.Zeros(2, 3)
	rows, cols := m.Size()
	if rows != 2 || cols != 3 {
		t.FailNow()
	}

	val := m.Values()
	if (*val)[0][0] != 0 {
		t.FailNow()
	}
	if (*val)[0][1] != 0 {
		t.FailNow()
	}
	if (*val)[0][2] != 0 {
		t.FailNow()
	}
	if (*val)[1][0] != 0 {
		t.FailNow()
	}
	if (*val)[1][1] != 0 {
		t.FailNow()
	}
	if (*val)[1][2] != 0 {
		t.FailNow()
	}
}

func TestId(t *testing.T) {
	m := lago.Id(4)
	rows, cols := m.Size()
	if rows != 4 || cols != 4 {
		t.FailNow()
	}

	val := m.Values()
	if (*val)[0][0] != 1 {
		t.FailNow()
	}
	if (*val)[0][1] != 0 {
		t.FailNow()
	}
	if (*val)[0][2] != 0 {
		t.FailNow()
	}
	if (*val)[0][3] != 0 {
		t.FailNow()
	}
	if (*val)[1][0] != 0 {
		t.FailNow()
	}
	if (*val)[1][1] != 1 {
		t.FailNow()
	}
	if (*val)[2][2] != 1 {
		t.FailNow()
	}
	if (*val)[3][3] != 1 {
		t.FailNow()
	}
}