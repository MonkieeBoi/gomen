package gomen_test

import (
	. "github.com/MonkieeBoi/gomen"
	"testing"
)

func shapeEq(a [][2]int, b [][2]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func printShapeErr(t *testing.T, want [][2]int, got [][2]int) {
	if !shapeEq(want, got) {
		t.Errorf(`expected: %v, got: %v`, want, got)
	}
}

func TestOperationShape1(t *testing.T) {
	want := [][2]int{}
	op := NewOperation(M_X, R_SPAWN, 3, 10)
	printShapeErr(t, want, op.Shape())
}

func TestOperationShape2(t *testing.T) {
	want := [][2]int{}
	op := NewOperation(M__, R_REVERSE, 6, 20)
	printShapeErr(t, want, op.Shape())
}

func TestOperationShape3(t *testing.T) {
	want := [][2]int{{5, 3}, {4, 3}, {6, 3}, {7, 3}}
	op := NewOperation(M_I, R_SPAWN, 5, 3)
	printShapeErr(t, want, op.Shape())
}

func TestOperationShape4(t *testing.T) {
	want := [][2]int{{4, 8}, {4, 7}, {3, 8}, {3, 9}}
	op := NewOperation(M_S, R_LEFT, 4, 8)
	printShapeErr(t, want, op.Shape())
}

func TestOperationShape5(t *testing.T) {
	want := [][2]int{{2, 1}, {2, 2}, {1, 1}, {1, 0}}
	op := NewOperation(M_Z, R_LEFT, 2, 1)
	printShapeErr(t, want, op.Shape())
}

func TestOperationShape6(t *testing.T) {
	want := [][2]int{{4,10},{4,9},{4,11},{3,10}}
	op := NewOperation(M_T, R_LEFT, 4, 10)
	printShapeErr(t, want, op.Shape())
}

func TestOperationShape7(t *testing.T) {
	want := [][2]int{{4,7},{4,8},{3,7},{3,8}}
	op := NewOperation(M_O, R_LEFT, 4, 7)
	printShapeErr(t, want, op.Shape())
}

func TestOperationShape8(t *testing.T) {
	want := [][2]int{{8,16},{8,15},{8,17},{7,15}}
	op := NewOperation(M_J, R_LEFT, 8, 16)
	printShapeErr(t, want, op.Shape())
}

func TestOperationShape9(t *testing.T) {
	want := [][2]int{{6,9},{6,8},{6,10},{5,10}}
	op := NewOperation(M_L, R_LEFT, 6, 9)
	printShapeErr(t, want, op.Shape())
}

func TestOperationShape10(t *testing.T) {
	want := [][2]int{{4,2},{4,1},{4,3},{3,2}}
	op := NewOperation(M_T, R_LEFT, 4, 2)
	printShapeErr(t, want, op.Shape())
}
