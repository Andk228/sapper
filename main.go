package main

import (
	"fmt"
)

type fieldValues struct {
	bomb         string
	freeSpace    string
	unknownSpace string
}

func NewfieldsValues() *fieldValues {
	return &fieldValues{
		bomb:         "x",
		freeSpace:    " ",
		unknownSpace: "*",
	}
}

type Field struct {
	Rows int
	Cols int
	Data [][]string
}

func NewField() *Field {
	f := &Field{
		Rows: 10,
		Cols: 10,
		Data: make([][]string, 10),
	}

	defaultVals := NewfieldsValues()
	for i := 0; i < f.Rows; i++ {
		f.Data[i] = make([]string, f.Cols)
		for j := 0; j < f.Cols; j++ {
			f.Data[i][j] = defaultVals.unknownSpace
		}
	}
	return f
}

func (field *Field) printField() {
	fmt.Print("   ")
	for c := 0; c < field.Cols; c++ {
		fmt.Printf("%2d ", c)
	}
	fmt.Println()

	for i := 0; i < field.Rows; i++ {
		fmt.Printf("%2d ", i)
		for j := 0; j < field.Cols; j++ {
			fmt.Printf(" %s ", field.Data[i][j])
		}
		fmt.Println()
	}
}

func main() {
	field := NewField()
	field.printField()

}
