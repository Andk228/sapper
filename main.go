package main

import (
	"fmt"
)

type statusValue struct {
	bomb         string
	freeSpace    string
	unknownSpace string
}

func NewstatusValue() *statusValue {
	return &statusValue{
		bomb:         "x",
		freeSpace:    " ",
		unknownSpace: "*",
	}
}

type fieldValues struct {
	value       string
	secretValue string
}

type Field struct {
	Rows int
	Cols int
	Data [][]fieldValues
}

func NewField() *Field {
	f := &Field{
		Rows: 10,
		Cols: 10,
		Data: make([][]fieldValues, 10),
	}

	defaultVals := NewstatusValue()
	for i := 0; i < f.Rows; i++ {
		f.Data[i] = make([]fieldValues, f.Cols)
		for j := 0; j < f.Cols; j++ {
			f.Data[i][j].value = defaultVals.unknownSpace
			if i == j {
				f.Data[i][j].secretValue = defaultVals.bomb
				continue
			}
			f.Data[i][j].secretValue = defaultVals.freeSpace

		}
	}
	return f
}

func (field *Field) printField(mod string) {
	fmt.Print("   ")
	for c := 0; c < field.Cols; c++ {
		fmt.Printf("%2d ", c)
	}
	fmt.Println()

	for i := 0; i < field.Rows; i++ {
		fmt.Printf("%2d ", i)
		for j := 0; j < field.Cols; j++ {
			if mod == "secret" {
				fmt.Printf(" %s ", field.Data[i][j].secretValue)
				continue
			}
			fmt.Printf(" %s ", field.Data[i][j].value)
		}
		fmt.Println()
	}
}

func main() {
	mod := "secrett"
	field := NewField()
	statusValue := NewstatusValue()

	fmt.Println("Игра сапер")
	fmt.Println("Открывай ячейки, пока не попадешься на бомбу")
	fmt.Println("Формат ввода x и y, где x это строка, а y столбец")
	for {
		var i, j int
		field.printField(mod)
		fmt.Print("Ваш ход: ")
		_, err := fmt.Scan(&i, &j)
		if err != nil {
			fmt.Println("Введите числа от 0 до 9")
			continue
		}

		field.Data[i][j].value = field.Data[i][j].secretValue

		if field.Data[i][j].secretValue == statusValue.bomb {
			field.printField("")
			fmt.Println("Вы проиграли")
			break
		}

	}

}
