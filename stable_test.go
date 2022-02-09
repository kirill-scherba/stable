package stable

import (
	"fmt"
	"testing"
)

func TestStructToTable(t *testing.T) {
	fmt.Print("Simple table test:\n\n")

	type table struct {
		ID     string
		Name   string
		Num    int
		Float  float64
		Status string
		IPv4   string
	}
	var a = []table{
		{"ququ-1", "ququ-1_name", 1, 3.14, "ready", "1.1.1.1"},
		{"ququ-2", "ququ-2_name", 2, 512.75, "not-ready", "22.2.2.2"},
	}

	s := new(Stable)

	// lens := structToListLens(a)
	// aligns := []int{}
	// fmt.Println(structToListTitle(a[0], lens, aligns))
	// fmt.Println(structToListLine(a[0], lens, aligns))

	// default
	fmt.Print(s.StructToTable(a) + "\n\n")

	// with column align and lines
	var totals table
	s.Lines().Aligns(0, 0, 1, 1, 1, 1).Totals(&totals, 0, 0, 1)
	fmt.Print(s.StructToTable(a) + "\n\n")
}
