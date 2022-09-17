package game

import (
	"math/rand"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

type Field struct {
	difficulty Difficulty
	rows, cols int
	cells      [][]*Cell
	minesCount int
}

func NewField(difficulty Difficulty) *Field {
	f := &Field{difficulty: difficulty}
	f.initDifficulty()
	cells := make([][]*Cell, f.rows)
	for row := 0; row < f.rows; row++ {
		cells[row] = make([]*Cell, f.cols)
		for col := 0; col < f.cols; col++ {
			cells[row][col] = newCell(f, NewPos(row, col))
		}
	}
	f.cells = cells
	return f
}

func (f *Field) InitMines(firstCell *Cell) {
	i := 0
	firstCell.GetNeighborhoods()
	noMineCells := append(firstCell.GetNeighborhoods(), firstCell)
	// place `f.minesCount` mines
	for i < f.minesCount {
		// pick random cell
		row := rand.Intn(f.rows)
		col := rand.Intn(f.cols)
		cell := f.Cell(NewPos(row, col))
		// this cell already has mine - skip
		if cell.IsMine() {
			continue
		}
		isNoMineSell := false
		for _, noMineCell := range noMineCells {
			if noMineCell == cell {
				isNoMineSell = true
				break
			}
		}
		// this cell cannot have mine - skip
		if isNoMineSell {
			continue
		}
		// place mine
		cell.isMine = true
		i++
	}
	f.calcNeighborhoodsMines()
}

func (f *Field) Cell(pos Pos) *Cell {
	if pos.row >= f.rows || pos.row < 0 {
		return nil
	}
	if pos.col >= f.cols || pos.col < 0 {
		return nil
	}
	return f.cells[pos.row][pos.col]
}

func (f *Field) Cells() [][]*Cell {
	return f.cells
}

func (f *Field) Rows() int {
	return f.rows
}

func (f *Field) Cols() int {
	return f.cols
}

func (f *Field) IsEmpty() bool {
	return len(f.cells) == 0
}

func (f *Field) Reveal(cell *Cell) bool {
	if cell.isRevealed {
		return false
	}
	cell.reveal()
	if cell.NeighborhoodsMines() != 0 {
		return f.isWin()
	}
	neighborhoods := cell.GetNeighborhoods()
	for _, neighborhood := range neighborhoods {
		f.Reveal(neighborhood)
	}
	return f.isWin()
}

func (f *Field) RevealAll() {
	for row := 0; row < f.rows; row++ {
		for col := 0; col < f.cols; col++ {
			cell := f.Cell(NewPos(row, col))
			cell.reveal()
		}
	}
}

func (f *Field) MinesCount() int {
	return f.minesCount
}

func (f *Field) FlaggedCount() int {
	flaggedCount := 0
	for row := 0; row < f.rows; row++ {
		for col := 0; col < f.cols; col++ {
			cell := f.Cell(NewPos(row, col))
			if cell.isFlag {
				flaggedCount++
			}
		}
	}
	return flaggedCount
}

func (f *Field) initDifficulty() {
	switch f.difficulty {
	case Easy:
		f.rows, f.cols, f.minesCount = 10, 10, 16
	case Medium:
		f.rows, f.cols, f.minesCount = 16, 16, 40
	case Hard:
		f.rows, f.cols, f.minesCount = 16, 30, 99
	}
}

func (f *Field) calcNeighborhoodsMines() {
	for row := 0; row < f.rows; row++ {
		for col := 0; col < f.cols; col++ {
			cell := f.Cell(NewPos(row, col))
			neighborhoods := cell.GetNeighborhoods()
			for _, neighborhood := range neighborhoods {
				if neighborhood.isMine {
					cell.neighborhoodsMines++
				}
			}
		}
	}
}

func (f *Field) isWin() bool {
	for row := 0; row < f.rows; row++ {
		for col := 0; col < f.cols; col++ {
			cell := f.Cell(NewPos(row, col))
			if !cell.isRevealed && !cell.isMine {
				return false
			}
		}
	}
	return true
}
