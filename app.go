package main

import (
	"github.com/kataras/iris"
	// "github.com/iris-contrib/middleware/cors"
)

func main() {
	app := iris.New()

	/* Cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug: true,
	})

	app.Use(Cors) */

	app.RegisterView(iris.HTML("./views", ".html"))

	// Routes
	app.Get("/", index)
	app.Get("/login", loginUser)
	app.Get("/logout", logoutUser)
	app.Get("/fetchposts", fetchPosts)
	app.Get("/create", newPost)
	app.Post("/create", createPosts)
	app.Post("/update", updatePosts)
	app.Post("/delete", deletePosts)

	app.Run(iris.Addr(":8080"))
}
