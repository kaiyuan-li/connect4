package main

import (
	"fmt"
	"strings"
)

type Board interface {
	Print()
}

const ROWS = 6
const COLS = 7

type Connect4Board struct {
	state [ROWS][COLS]byte // 2D array with dimension of Rows x Columns
}

func (b *Connect4Board) Print() {
	printLine := func() {
		fmt.Println("  -" + strings.Repeat("----", COLS))
	}
	printLine()
	for ir, row := range b.state {
		fmt.Printf("%d ", ir)
		for _, e := range row {
			fmt.Printf("| %s ", string(e))
		}
		fmt.Println("|")
		printLine()
	}
	fmt.Print(" ")
	for c := 0; c < COLS; c++ {
		fmt.Printf("   %d", c)
	}
	fmt.Println()
}

func main() {
	s := [ROWS][COLS]byte{}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			s[i][j] = ' '
		}
	}
	b := Connect4Board{
		state: s,
	}
	b.Print()
}
