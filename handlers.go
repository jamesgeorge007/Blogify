package main

import (
	"log"

	"github.com/Blogify/models"
	repo "github.com/Blogify/repositories"

	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

func index(ctx iris.Context) {
	posts, _ := repo.GetAll()

	ctx.ViewData("Posts", posts)
	ctx.View("index.html")
}

func loginUser(ctx iris.Context) {
	// Login User
}

func fetchPosts(ctx iris.Context) {
	// Fetch Posts
	posts, _ := repo.GetAll()
	ctx.JSON(posts)
}
func newPost(ctx iris.Context) {
	ctx.View("create_posts.html")
}
func createPosts(ctx iris.Context) {
	post := models.Post{}
	err := ctx.ReadForm(&post)
	if err != nil {
		log.Println(err.Error())
	}
	post.ID = bson.NewObjectId()
	repo.CreateOne(post)
	ctx.Redirect("/")
}

func updatePosts(ctx iris.Context) {
	var post models.Post
	_ = ctx.ReadJSON(&post)
	repo.Update(post)
	ctx.ViewData("Post", models.Post{
		ID:          post.ID,
		Name:        post.Name,
		Description: post.Description,
	})
	ctx.View("update_posts.html")
}

func deletePosts(ctx iris.Context) {
	var post models.Post
	_ = ctx.ReadJSON(&post)
	repo.Delete(post)
	ctx.ViewData("Post", models.Post{
		ID:          post.ID,
		Name:        post.Name,
		Description: post.Description,
	})
	ctx.View("delete_posts.html")
}

func logoutUser(ctx iris.Context) {
	// Logout User
}
