package main

import (
	"context"
	"github.com/sanprasirt/blog-application/global"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	global.DB.Collection("user").InsertOne(context.Background(), bson.M{"name": "test"})
}
