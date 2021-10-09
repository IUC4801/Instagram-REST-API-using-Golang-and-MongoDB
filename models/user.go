package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"pwd" bson:"pwd"`
}

type Post struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Caption string        `json:"caption" bson:"caption"`
	URL     string        `json:"url" bson:"url"`
	Time    string        `json:"time" bson:"time"`
}
