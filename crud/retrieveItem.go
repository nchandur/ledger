package crud

import (
	"context"
	"ledger/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (g *Group) RetrieveItem(item string) (models.Expense, error) {
	filter := bson.D{{Key: "item", Value: item}}

	var res models.Expense

	err := g.Collection.FindOne(context.TODO(), filter).Decode(&res)

	return res, err

}
