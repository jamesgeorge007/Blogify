package models

import "gopkg.in/mgo.v2/bson"

type Post struct {
	ID          bson.ObjectId
	Name        string `form:"name"`
	Description string `form:"description"`
}
