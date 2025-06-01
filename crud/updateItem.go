package crud

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (g *Group) UpdateItem(item string, update map[string]any) error {
	filter := bson.D{
		{Key: "item", Value: item},
	}

	_, err := g.Collection.UpdateOne(
		context.TODO(),
		filter,
		bson.D{{Key: "$set", Value: update}},
	)

	if err != nil {
		return err
	}

	return nil

}
