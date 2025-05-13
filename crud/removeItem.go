package crud

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (g *Group) RemoveItem(itemID int, date time.Time) error {

	filter := bson.D{{Key: "item_id", Value: itemID}}

	count, err := g.Collection.DeleteMany(context.TODO(), filter)

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
