package crud

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (g *Group) UpdatePrice(itemID int, update bson.D) error {
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
