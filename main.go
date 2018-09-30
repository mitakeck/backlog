package main

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("backlog", "backlog")

	login      = app.Command("login", "login")
	loginSpace = login.Arg("s", "space").Required().String()

	refresh = app.Command("refresh", "refresh access token")

	create        = app.Command("create", "create issue")
	createProject = create.Arg("p", "project").String()

	view = app.Command("view", "view issue detail")
)

func main() {
	cmdopt := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch cmdopt {
	case login.FullCommand():
		Login(*loginSpace)

	case refresh.FullCommand():
		fmt.Println("refresh")
		model, err := ReadConfig()
		if err != nil {
			fmt.Println(err)
		}
		pp.Println(model)

	case create.FullCommand():
		fmt.Println("create")

	case view.FullCommand():
		fmt.Println("view")
	}
}
