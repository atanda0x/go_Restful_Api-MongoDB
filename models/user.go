package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Name   string        `json:"name" bson:"name"`
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Age    int64         `json:"age" bson:"age"`
	Gender string        `json:"gender" bson:"gender"`
}
