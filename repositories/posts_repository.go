package repositories

import (
	conn "github.com/Blogify/config"

	"github.com/Blogify/models"

	"gopkg.in/mgo.v2/bson"
)

var new_post = models.Post{}

var dbconf = conn.Configurations{}
var db = dbconf.Connect()

const (
	COLLECTION = "posts"
)

// Update a post
func Update(post models.Post) error {
	return db.C(COLLECTION).UpdateId(&post.ID, &post)
}

// Delete a post
func Delete(post models.Post) error {
	return db.C(COLLECTION).RemoveId(&post.ID)
}

//Creates a new post
func CreateOne(post models.Post) error {
	new_post.ID = bson.NewObjectId()
	new_post.Name = post.Name
	new_post.Description = post.Description
	err := db.C(COLLECTION).Insert(&new_post)

	return err
}

//Retrieve posts

func GetAll() ([]models.Post, error) {
	var posts []models.Post
	err := db.C(COLLECTION).Find(bson.M{}).All(&posts)

	return posts, err
}
