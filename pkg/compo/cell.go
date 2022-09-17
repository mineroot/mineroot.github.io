package compo

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"minesweeper/pkg/game"
)

type cellCompo struct {
	app.Compo
	IsRevealed         bool
	IsMine             bool
	NeighborhoodsMines int
	IsFlag             bool
	IsExploded         bool
	pos                game.Pos
}

func newCellCompo(
	isRevealed bool,
	isMine bool,
	neighborhoodsMines int,
	isFlag bool,
	isExploded bool,
	pos game.Pos,
) *cellCompo {
	return &cellCompo{
		IsRevealed:         isRevealed,
		IsMine:             isMine,
		NeighborhoodsMines: neighborhoodsMines,
		IsFlag:             isFlag,
		IsExploded:         isExploded,
		pos:                pos,
	}
}

func (p *cellCompo) Render() app.UI {
	return app.Div().
		Class(p.class()).
		OnClick(p.onClick).
		OnContextMenu(p.onRightClick).
		Body(
			app.Div().
				Class("app-cell-content").
				Body(
					app.If(p.IsRevealed && p.IsMine,
						app.Img().Src("/web/img/mine.svg").Draggable(false),
					).ElseIf(p.IsRevealed && p.NeighborhoodsMines != 0,
						app.Text(p.NeighborhoodsMines),
					).ElseIf(!p.IsRevealed && p.IsFlag,
						app.Img().Src("/web/img/flag.svg").Draggable(false),
					),
				),
		)
}

func (p *cellCompo) class() string {
	class := app.AppendClass("", "app-cell")
	if !p.IsRevealed {
		return class
	}
	class = app.AppendClass(class, "app-cell-revealed")
	if p.IsExploded {
		return app.AppendClass(class, "app-cell-exploded")
	}
	if p.IsMine {
		return app.AppendClass(class, "app-cell-neighborhoods-0")
	}
	return app.AppendClass(class, fmt.Sprintf("app-cell-neighborhoods-%d", p.NeighborhoodsMines))
}

func (p *cellCompo) onClick(ctx app.Context, e app.Event) {
	if p.IsFlag {
		return
	}
	ctx.NewActionWithValue(cellRevealedAction, p.pos)
}

func (p *cellCompo) onRightClick(ctx app.Context, e app.Event) {
	e.PreventDefault()
	ctx.NewActionWithValue(setFlagAction, p.pos)
}
