package stable

import (
	"fmt"
	"testing"
)

func TestStructToTable(t *testing.T) {
	fmt.Print("Simple table test:\n\n")

	var a = []struct {
		ID     string
		Name   string
		Num    int
		Status string
		IPv4   string
	}{
		{"ququ-1", "ququ-1_name", 1, "ready", "1.1.1.1"},
		{"ququ-2", "ququ-2_name", 2, "ready", "2.2.2.2"},
	}

	s := new(Stable)

	// lens := structToListLens(a)
	// aligns := []int{}
	// fmt.Println(structToListTitle(a[0], lens, aligns))
	// fmt.Println(structToListLine(a[0], lens, aligns))

	// default
	fmt.Print(s.StructToTable(a) + "\n\n")

	// with align
	fmt.Print(s.StructToTable(a, 0, 0, 1) + "\n\n")
}
