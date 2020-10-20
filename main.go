package main

import (
	"espadafajar/restku/db"
	"espadafajar/restku/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
