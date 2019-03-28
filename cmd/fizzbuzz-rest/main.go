package main

import (
	"github.com/melkor/fizzbuzz-rest/app"
)

func main() {
	app := app.New()
	app.Run(":8000")
}
