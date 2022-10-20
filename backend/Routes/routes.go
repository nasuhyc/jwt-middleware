package routes

import (
	controllers "jwt-middleware/backend/Controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api") // /api
	//User Group
	user := api.Group("/user")
	user.Get("/", controllers.UserIndex)
	user.Post("/create", controllers.UserStore)
	user.Post("/update/:id", controllers.UserUpdate)
	user.Get("/get/:id", controllers.UserGetId)
	user.Delete("/delete/:id", controllers.UserDestroy)
	//User Group End

	//Login
	api.Post("/login", controllers.Login)
	//User

}
