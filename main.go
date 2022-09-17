package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"math/rand"
	"minesweeper/pkg/compo"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	app.Route("/", compo.NewHomePage())
	app.RunWhenOnBrowser()

	handler := &app.Handler{
		Name:        "Minesweeper",
		Description: "Minesweeper game",
		Styles: []string{
			"/web/css/bootstrap.min.css",
			"/web/css/main.css",
		},
		Scripts: []string{
			"/web/js/bootstrap.min.js",
		},
		CacheableResources: []string{
			"/web/img/flag.svg",
			"/web/img/mine.svg",
		},
		//Resources: app.GitHubPages("minesweeper"),
	}
	http.Handle("/", handler)

	//err := app.GenerateStaticWebsite("", handler)
	//if err != nil {
	//	log.Fatal(err)
	//}

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
