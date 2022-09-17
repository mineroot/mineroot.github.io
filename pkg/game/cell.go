package game

type Cell struct {
	Pos
	isRevealed         bool
	isMine             bool
	neighborhoodsMines int
	field              *Field
	isFlag             bool
	exploded           bool
}

func newCell(field *Field, pos Pos) *Cell {
	return &Cell{
		Pos:   pos,
		field: field,
	}
}

func (c *Cell) Position() Pos {
	return c.Pos
}

func (c *Cell) GetNeighborhoods() []*Cell {
	neighborhoodsDeltas := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	var neighborhoods []*Cell
	for _, delta := range neighborhoodsDeltas {
		deltaRow, deltaCol := delta[0], delta[1]
		deltaPos := NewPos(c.row+deltaRow, c.col+deltaCol)
		neighborhood := c.field.Cell(deltaPos)
		if neighborhood != nil {
			neighborhoods = append(neighborhoods, neighborhood)
		}
	}
	return neighborhoods
}

func (c *Cell) NeighborhoodsMines() int {
	return c.neighborhoodsMines
}

func (c *Cell) IsRevealed() bool {
	return c.isRevealed
}

func (c *Cell) IsMine() bool {
	return c.isMine
}

func (c *Cell) IsFlag() bool {
	return c.isFlag
}

func (c *Cell) SwitchFlag() {
	c.isFlag = !c.isFlag
}

func (c *Cell) Exploded() bool {
	return c.exploded
}

func (c *Cell) Explode() {
	c.exploded = true
}

func (c *Cell) reveal() {
	c.isRevealed = true
}
