package compo

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type gameStatusCompo struct {
	app.Compo
	Status string
	class  string
}

func (c *gameStatusCompo) OnMount(ctx app.Context) {
	ctx.Handle(gameWinAction, c.onGameWin)
	ctx.Handle(gameOverAction, c.onGameOver)
}

func (c *gameStatusCompo) Render() app.UI {
	return app.Div().
		Class("fs-1", c.class).
		Body(
			app.Text(c.Status),
		)
}

func (c *gameStatusCompo) onGameWin(app.Context, app.Action) {
	c.Status = "You won!"
	c.class = "text-success"
}

func (c *gameStatusCompo) onGameOver(app.Context, app.Action) {
	c.Status = "Game over!"
	c.class = "text-danger"
}

func (c *gameStatusCompo) reset() {
	c.Status = ""
	c.class = ""
	c.Update()
}
