package crud

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (g *Group) UpdateItemByID(itemID int, update map[string]any) error {
	filter := bson.D{
		{Key: "item_id", Value: itemID},
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
