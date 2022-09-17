package compo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"minesweeper/pkg/game"
)

type fieldCompo struct {
	app.Compo
	Field       *game.Field
	isGameOver  bool
	gameStarted bool
}

func (c *fieldCompo) OnMount(ctx app.Context) {
	ctx.Handle(cellRevealedAction, c.onCellReveal)
	ctx.Handle(setFlagAction, c.onSetFlag)
	ctx.Handle(gameWinAction, c.onGameWin)
	ctx.Handle(gameOverAction, c.onGameOver)
}

func (c *fieldCompo) Render() app.UI {
	return app.Div().Class("app-table").Body(
		app.Range(c.Field.Cells()).Slice(func(row int) app.UI {
			return app.Div().Class("app-row").Body(
				app.Range(c.Field.Cells()[row]).Slice(func(col int) app.UI {
					cell := c.Field.Cell(game.NewPos(row, col))
					return newCellCompo(
						cell.IsRevealed(),
						cell.IsMine(),
						cell.NeighborhoodsMines(),
						cell.IsFlag(),
						cell.Exploded(),
						cell.Position(),
					)
				}),
			)
		}),
	)
}

func (c *fieldCompo) reset(newField *game.Field) {
	c.Field = newField
	c.isGameOver = false
	c.gameStarted = false
	c.Update()
}

func (c *fieldCompo) onCellReveal(ctx app.Context, act app.Action) {
	if c.isGameOver {
		return
	}
	pos := act.Value.(game.Pos)
	cell := c.Field.Cell(pos)
	if !c.gameStarted {
		ctx.NewAction(gameStartAction)
		c.gameStarted = true
		c.Field.InitMines(cell)
	}
	if cell.IsMine() {
		ctx.NewActionWithValue(gameOverAction, cell)
		return
	}
	isWin := c.Field.Reveal(cell)
	if isWin {
		ctx.NewAction(gameWinAction)
	}
}

func (c *fieldCompo) onSetFlag(_ app.Context, act app.Action) {
	if c.isGameOver {
		return
	}
	pos := act.Value.(game.Pos)
	cell := c.Field.Cell(pos)
	if cell.IsRevealed() {
		return
	}
	cell.SwitchFlag()
}

func (c *fieldCompo) onGameOver(_ app.Context, act app.Action) {
	cell := act.Value.(*game.Cell)
	cell.Explode()
	c.Field.RevealAll()
	c.isGameOver = true
}

func (c *fieldCompo) onGameWin(app.Context, app.Action) {
	c.isGameOver = true
}
