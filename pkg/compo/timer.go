package compo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"time"
)

type timerCompo struct {
	app.Compo
	Time         time.Time
	ticker       *time.Ticker
	stopTickerCh chan struct{}
}

func (c *timerCompo) OnMount(ctx app.Context) {
	ctx.Handle(gameStartAction, c.startTicker)
	ctx.Handle(gameWinAction, c.onStopTicker)
	ctx.Handle(gameOverAction, c.onStopTicker)
}

func (c *timerCompo) Render() app.UI {
	elapsed := c.Time.Format("04:05")
	return app.Div().
		Class("fs-2").
		Text(elapsed)
}

func (c *timerCompo) reset() {
	c.stopTicker()
	c.Time = time.Time{}
	c.Update()
}

func (c *timerCompo) startTicker(ctx app.Context, _ app.Action) {
	ctx.Async(func() {
		c.ticker = time.NewTicker(time.Second)
		c.stopTickerCh = make(chan struct{})
		for {
			select {
			case <-c.ticker.C:
				ctx.Dispatch(func(app.Context) {
					c.Time = c.Time.Add(time.Second)
				})
			case <-c.stopTickerCh:
				c.ticker.Stop()
				c.stopTickerCh = nil
				return
			}
		}
	})
}

func (c *timerCompo) onStopTicker(app.Context, app.Action) {
	c.stopTicker()
}

func (c *timerCompo) stopTicker() {
	if c.stopTickerCh != nil {
		close(c.stopTickerCh)
	}
}
