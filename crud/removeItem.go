package crud

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (g *Group) RemoveItem(item string) error {

	filter := bson.D{{Key: "item", Value: item}}

	count, err := g.Collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	if count.DeletedCount == 0 {
		fmt.Printf("no item found\n")
		return nil
	}

	fmt.Printf("deleted item\n")
	return nil
}
