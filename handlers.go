package main

import (
	"Blogify/models"
	repo "Blogify/repositories"
	"log"

	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

func index(ctx iris.Context) {
	ctx.ViewData("message", "Welcome")
	ctx.View("index.html")
}

func loginUser(ctx iris.Context) {
	// Login User
}

func fetchPosts(ctx iris.Context) {
	// Fetch Posts

}

func createPosts(ctx iris.Context) {
	post := models.Post{}
	name := ctx.PostValue("name")
	description := ctx.PostValue("description")
	post.ID = bson.NewObjectId()
	post.Name = name
	post.Description = description
	log.Println("Post values: ", post.Description)
	repo.CreateOne(post)

	ctx.View("create_posts.html")
}

func updatePosts(ctx iris.Context) {
	// Handle Update Blog.
}

func deletePosts(ctx iris.Context) {
	// Handle Delete Blog.
}

func logoutUser(ctx iris.Context) {
	// Logout User
}
