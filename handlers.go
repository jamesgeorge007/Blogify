package main

import (
"github.com/kataras/iris"
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
    		/*username := ctx.PostValue("Username")
    		ctx.ViewData("user", username)
    		ctx.View("welcome.html") */
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
