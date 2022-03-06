package mongorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) UpdateByID(id primitive.ObjectID, document *T, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateByID(context.Background(), id, bson.M{"$set": document}, options...)
	return
}

func (c *C[T]) UpdateCustom(filter, data bson.M, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateOne(context.Background(), filter, bson.M{
		"$set": data,
	}, options...)
	return
}

func (c *C[T]) UpdateCustomMany(filter, data bson.M, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateMany(context.Background(), filter, bson.M{
		"$set": data,
	}, options...)
	return
}
