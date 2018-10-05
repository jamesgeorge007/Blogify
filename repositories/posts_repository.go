package repositories

import (
	conn "Blogify/config"
	"Blogify/models"

	"gopkg.in/mgo.v2/bson"
)

var new_post = models.Post{}

var dbconf = conn.Configurations{}
var db = dbconf.Connect()

const (
	COLLECTION = "posts"
)

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
