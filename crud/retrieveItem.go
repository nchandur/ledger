package crud

import (
	"context"
	"ledger/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (g *Group) RetrieveItemByID(itemID int) (models.Expense, error) {
	filter := bson.D{{Key: "item_id", Value: itemID}}

	var res models.Expense

	err := g.Collection.FindOne(context.TODO(), filter).Decode(&res)

	return res, err

}
