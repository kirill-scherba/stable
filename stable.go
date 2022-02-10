// Copyright 2021 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Stable create Simple Table string from struct
package stable

import (
	"fmt"
	"reflect"
	"strings"
)

type Stable struct {
	lines     bool
	aligns    []int
	totals    []int
	totalSPtr interface{}
	cleanLine string
	formats   []string
}

// title return list title string
func (t Stable) title(in interface{}, lens []int) (str string) {
	val := reflect.Indirect(reflect.ValueOf(in))
	l := val.NumField()
	str = t.cleanLine
	for i := 0; i < l; i++ {
		s := fmt.Sprint(strings.ToUpper(val.Type().Field(i).Name))
		if len(t.aligns) > i && t.aligns[i] > 0 {
			str += fmt.Sprintf("%*s", lens[i], s)
		} else {
			str += fmt.Sprintf("%-*s", lens[i], s)
		}
		if i < l-1 {
			str += "   "
		}
	}
	return
}

// line return list body string
func (t Stable) line(in interface{}, lens []int) (str string) {
	val := reflect.Indirect(reflect.ValueOf(in))
	l := val.NumField()
	str = t.cleanLine
	for i := 0; i < l; i++ {
		s := fmt.Sprintf(t.getFormat(i), val.Field(i))
		if len(t.aligns) > i && t.aligns[i] > 0 {
			str += fmt.Sprintf("%*s", lens[i], s)
		} else {
			str += fmt.Sprintf("%-*s", lens[i], s)
		}
		if i < l-1 {
			str += "   "
		}
		t.totalAdd(i, val.Field(i))
	}
	return
}

// totalAdd add table line value to totals
func (t Stable) totalAdd(i int, v reflect.Value) {
	if i >= len(t.totals) || t.totals[i] == 0 {
		return
	}

	val := reflect.Indirect(reflect.ValueOf(t.totalSPtr))

	switch val.Field(i).Kind() {
	case reflect.Int, reflect.Int64:
		tot := val.Field(i).Int() + v.Int()
		val.Field(i).SetInt(tot)
	case reflect.Float64, reflect.Float32:
		tot := val.Field(i).Float() + v.Float()
		val.Field(i).SetFloat(tot)
	}
}

// total return total string
func (t Stable) total(lens []int) (str string) {
	val := reflect.Indirect(reflect.ValueOf(t.totalSPtr))
	l := len(t.totals)
	for i := range t.totals {
		if i == 0 {
			str = "\n" + t.cleanLine
		}
		if t.totals[i] == 0 {
			str += fmt.Sprintf("%*s", lens[i], "")
		} else {
			s := fmt.Sprintf(t.getFormat(i), val.Field(i))
			str += fmt.Sprintf("%*s", lens[i], s)
		}
		if i < l-1 {
			str += "   "
		}
	}
	return
}

// titleLens calculate lists title columns len
func (t Stable) titleLens(in interface{}) (lens []int) {

	// Calculate titles length
	val := reflect.Indirect(reflect.ValueOf(in))
	l := val.NumField()
	for i := 0; i < l; i++ {
		lens = append(lens, len(fmt.Sprint(strings.ToUpper(val.Type().Field(i).Name))))
	}

	return
}

// lineLens calculate lists line columns len
func (t Stable) lineLens(in interface{}, lens []int) []int {
	val := reflect.Indirect(reflect.ValueOf(in))
	l := val.NumField()
	for i := 0; i < l; i++ {
		ln := len(fmt.Sprintf(t.getFormat(i), val.Field(i)))
		if ln > lens[i] {
			lens[i] = ln
		}
	}
	return lens
}

// lens calculate lists columns len
func (t Stable) lens(in interface{}) (lens []int, sumLen int) {
	switch reflect.TypeOf(in).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(in)
		if s.Len() == 0 {
			return
		}
		for i := 0; i < s.Len(); i++ {
			if i == 0 {
				lens = t.titleLens(s.Index(i).Interface())
			}
			lens = t.lineLens(s.Index(i).Interface(), lens)
		}
		for i := range lens {
			sumLen += lens[i]
		}
		if len(lens) > 1 {
			sumLen += (len(lens) - 1) * 3
		}
	}
	return
}

// StructToTable convert structs slice to table string
func (t Stable) StructToTable(in interface{}) (str string) {
	switch reflect.TypeOf(in).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(in)
		if s.Len() == 0 {
			return
		}
		lens, sumLen := t.lens(in)
		line := func(l int, lineFeedLeft bool) (str string) {
			if t.lines {
				if lineFeedLeft {
					str = "\n"
				}
				str += t.cleanLine
				str += strings.Repeat("┈", l)
				if !lineFeedLeft {
					str += "\n"
				}
			}
			return
		}
		for i := 0; i < s.Len(); i++ {
			if i == 0 {
				str += line(sumLen, false)
				str += t.title(s.Index(i).Interface(), lens)
				str += line(sumLen, true)
			}
			str += "\n" + t.line(s.Index(i).Interface(), lens)
		}
		str += line(sumLen, true)
		str += t.total(lens)
	}
	return
}

// Lines add lines to table
func (t *Stable) Lines() *Stable {
	t.lines = true
	return t
}

// Aligns set align to colums. Align parameter is array with values of
// 0 - 'align left' or 1 - 'align right'
func (t *Stable) Aligns(aligns ...int) *Stable {
	t.aligns = aligns
	return t
}

// Totals add totals to table
func (t *Stable) Totals(sptr interface{}, totals ...int) *Stable {
	t.totals = totals
	t.totalSPtr = sptr
	return t
}

// CleanLine set clean line terminal command before all table rows
func (t *Stable) CleanLine(cl bool) *Stable {
	if cl {
		t.cleanLine = "\033[K"
	}
	return t
}

// Format sets columns format
func (t *Stable) Formats(formats ...string) *Stable {
	t.formats = formats
	return t
}

// getFormat retun column format
func (t Stable) getFormat(col int) (fmt string) {
	if col < len(t.formats) && t.formats[col] != "" {
		fmt = t.formats[col]
	} else {
		fmt = "%v"
	}
	return
}
