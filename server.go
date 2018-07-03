package main

import "github.com/kataras/iris"

func main() {
    app := iris.New()

    app.RegisterView(iris.HTML("./views", ".html"))

    app.Get("/", func(ctx iris.Context) {
        ctx.ViewData("message", "Welcome")
        ctx.View("index.html") 
    })

    app.Get("/create", func(ctx iris.Context){
    		ctx.View("create_post.html")
    	})

    /* app.Post("/create", func(ctx iris.Context){
    	username := ctx.PostValue("Username")
    	ctx.ViewData("user", username)
    	ctx.View("welcome.html")
    }) */

    app.Run(iris.Addr(":8080"))
}