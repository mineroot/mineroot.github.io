package compo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"minesweeper/pkg/game"
)

type HomePage struct {
	app.Compo
	FieldCompo      *fieldCompo
	TimerCompo      *timerCompo
	MinesCompo      *minesCompo
	GameStatusCompo *gameStatusCompo
}

func NewHomePage() *HomePage {
	return &HomePage{}
}

func (c *HomePage) OnPreRender(ctx app.Context) {
	c.initPage(ctx)
}

func (c *HomePage) OnNav(ctx app.Context) {
	c.initPage(ctx)
}

func (c *HomePage) Render() app.UI {
	return app.Main().
		Class("container-fluid", "app-main").
		Body(
			app.Div().Class("row", "text-center", "mt-3").Body(
				app.Div().
					Class("col-4").
					Body(
						app.Button().
							Class("btn", "btn-info").
							Text("Easy").
							OnClick(c.startEasyGame),
					),
				app.Div().
					Class("col-4").
					Body(
						app.Button().
							Class("btn", "btn-warning").
							Text("Medium").
							OnClick(c.startMediumGame),
					),
				app.Div().
					Class("col-4").
					Body(
						app.Button().
							Class("btn", "btn-danger").
							Text("Hard").
							OnClick(c.startHardGame),
					),
			),
			app.If(c.FieldCompo != nil,
				app.Div().
					Class("row", "text-center", "mt-5", "mb-5").
					Body(
						c.FieldCompo,
					),
				app.Div().
					Class("row", "text-center").
					Body(
						c.GameStatusCompo,
						c.TimerCompo,
						c.MinesCompo,
					),
			),
		)
}

func (c *HomePage) startEasyGame(app.Context, app.Event) {
	field := game.NewField(game.Easy)
	c.newGame(field)
}
func (c *HomePage) startMediumGame(app.Context, app.Event) {
	field := game.NewField(game.Medium)
	c.newGame(field)
}
func (c *HomePage) startHardGame(app.Context, app.Event) {
	field := game.NewField(game.Hard)
	c.newGame(field)
}

func (c *HomePage) newGame(newField *game.Field) {
	if c.FieldCompo == nil {
		panic("FieldCompo is nil")
	}
	c.FieldCompo.reset(newField)
	c.MinesCompo.reset(newField)
	c.TimerCompo.reset()
	c.GameStatusCompo.reset()
}

func (c *HomePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Minesweeper")
	field := game.NewField(game.Easy)
	c.FieldCompo = &fieldCompo{Field: field}
	c.MinesCompo = &minesCompo{counter: field}
	c.TimerCompo = &timerCompo{}
	c.GameStatusCompo = &gameStatusCompo{}
}
