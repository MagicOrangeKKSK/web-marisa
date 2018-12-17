package Routes

import (
	"github.com/kataras/iris"
	"web-marisa/server/Controllers"
)

func Configure(app *iris.Application) {
	// test
	app.Post("/test", Controllers.Test)

	// Index
	app.Get("/", Controllers.GetIndexHandler)

	// Core
	app.Post("/Add", Controllers.Add)
}
