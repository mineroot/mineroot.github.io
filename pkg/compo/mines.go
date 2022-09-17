package compo

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type minesCounter interface {
	FlaggedCount() int
	MinesCount() int
}

type minesCompo struct {
	app.Compo
	counter minesCounter
}

func (c *minesCompo) OnMount(ctx app.Context) {
	ctx.Handle(setFlagAction, func(app.Context, app.Action) {})
}

func (c *minesCompo) Render() app.UI {
	flaggedCount := c.counter.FlaggedCount()
	minesCount := c.counter.MinesCount()
	if flaggedCount > minesCount {
		flaggedCount = minesCount
	}
	return app.Div().
		Class("fs-3").
		Body(
			app.Text(fmt.Sprintf("%d/%d", flaggedCount, minesCount)),
		)
}

func (c *minesCompo) reset(counter minesCounter) {
	c.counter = counter
	c.Update()
}
