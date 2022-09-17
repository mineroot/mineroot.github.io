package game

type Pos struct {
	row, col int
}

func NewPos(row int, col int) Pos {
	return Pos{row: row, col: col}
}
