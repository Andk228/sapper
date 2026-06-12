package main

import (
	"fmt"
	"math/rand"
)

type statusValue struct {
	bomb         string
	freeSpace    string
	unknownSpace string
	flag         string
	countOfBombs Digits
}

type Digits struct {
	One   string
	Two   string
	Three string
	Four  string
	Five  string
	Six   string
	Seven string
	Eight string
}

func NewstatusValue() *statusValue {
	digits := NewDigits()
	return &statusValue{
		bomb:         "x",
		freeSpace:    " ",
		unknownSpace: "*",
		flag:         "^",
		countOfBombs: *digits,
	}
}

func NewDigits() *Digits {
	return &Digits{
		One:   "1",
		Two:   "2",
		Three: "3",
		Four:  "4",
		Five:  "5",
		Six:   "6",
		Seven: "7",
		Eight: "8",
	}
}

type fieldValues struct {
	value       string
	secretValue string
	bombsAround int
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
			f.Data[i][j].secretValue = defaultVals.freeSpace
		}
	}

	f.placeBombs(10, defaultVals)
	f.calculateHints(defaultVals)

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

func (field *Field) placeBombs(count int, values *statusValue) {
	placed := 0
	for placed < count {
		i := rand.Intn(field.Rows)
		j := rand.Intn(field.Cols)
		if field.Data[i][j].secretValue == values.bomb {
			continue
		}
		field.Data[i][j].secretValue = values.bomb
		placed++
	}
}

func (field *Field) calculateHints(values *statusValue) {
	for i := 0; i < field.Rows; i++ {
		for j := 0; j < field.Cols; j++ {
			if field.Data[i][j].secretValue == values.bomb {
				continue
			}
			count := field.countAdjacentBombs(i, j, values)
			field.Data[i][j].bombsAround = count
			if count == 0 {
				field.Data[i][j].secretValue = values.freeSpace
			} else {
				field.Data[i][j].secretValue = field.digitFor(count, values)
			}
		}
	}
}

func (field *Field) countAdjacentBombs(i, j int, values *statusValue) int {
	count := 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			i2 := i + di
			j2 := j + dj
			if field.inBounds(i2, j2) && field.Data[i2][j2].secretValue == values.bomb {
				count++
			}
		}
	}
	return count
}

func (field *Field) digitFor(count int, values *statusValue) string {
	switch count {
	case 1:
		return values.countOfBombs.One
	case 2:
		return values.countOfBombs.Two
	case 3:
		return values.countOfBombs.Three
	case 4:
		return values.countOfBombs.Four
	case 5:
		return values.countOfBombs.Five
	case 6:
		return values.countOfBombs.Six
	case 7:
		return values.countOfBombs.Seven
	case 8:
		return values.countOfBombs.Eight
	default:
		return values.freeSpace
	}
}

func (field *Field) inBounds(i, j int) bool {
	return i >= 0 && i < field.Rows && j >= 0 && j < field.Cols
}

func (field *Field) firstMove(i, j int, values *statusValue) {
	if field.Data[i][j].secretValue == values.bomb {
		field.Data[i][j].secretValue = values.freeSpace
		for {
			i2 := rand.Intn(field.Rows)
			j2 := rand.Intn(field.Cols)
			if field.Data[i2][j2].secretValue != values.bomb && (i2 != i || j2 != j) {
				field.Data[i2][j2].secretValue = values.bomb
				break
			}
		}
	}
	field.calculateHints(values)
}

func (field *Field) reveal(i, j int, values *statusValue) {
	if !field.inBounds(i, j) {
		return
	}
	if field.Data[i][j].value != values.unknownSpace {
		return
	}
	field.Data[i][j].value = field.Data[i][j].secretValue
	if field.Data[i][j].bombsAround == 0 {
		for di := -1; di <= 1; di++ {
			for dj := -1; dj <= 1; dj++ {
				if di == 0 && dj == 0 {
					continue
				}
				field.reveal(i+di, j+dj, values)
			}
		}
	}
}

func (field *Field) isWin(values *statusValue) bool {
	for i := 0; i < field.Rows; i++ {
		for j := 0; j < field.Cols; j++ {
			if field.Data[i][j].secretValue == values.bomb {
				continue
			}
			if field.Data[i][j].value == values.unknownSpace {
				return false
			}
		}
	}
	return true
}

func main() {
	field := NewField()
	statusValue := NewstatusValue()
	firstMoveDone := false

	fmt.Println("Игра сапер")
	fmt.Println("Открывай ячейки, пока не попадешься на бомбу")
	fmt.Println("Формат ввода x и y, где x это строка, а y столбец")
	for {
		field.printField("")
		var i, j int
		fmt.Print("Ваш ход: ")
		_, err := fmt.Scan(&i, &j)
		if err != nil || !field.inBounds(i, j) {
			fmt.Println("Введите числа от 0 до 9")
			continue
		}

		if !firstMoveDone {
			field.firstMove(i, j, statusValue)
			firstMoveDone = true
		}

		if field.Data[i][j].secretValue == statusValue.bomb {
			field.Data[i][j].value = statusValue.bomb
			field.printField("secret")
			fmt.Println("Вы проиграли")
			break
		}

		field.reveal(i, j, statusValue)

		if field.isWin(statusValue) {
			field.printField("")
			fmt.Println("Поздравляю! Вы выиграли")
			break
		}
	}
}
