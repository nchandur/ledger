package crud

import (
	"context"
	"fmt"
	"ledger/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (g *Group) AddItem(expense models.Expense) error {

	exists, err := Exists(g.Collection.Name())

	if err != nil {
		return fmt.Errorf("error adding item %s: %v", expense.Item, err)
	}

	if !exists {
		return fmt.Errorf("group does not exist")
	}

	count, err := g.Collection.CountDocuments(context.TODO(), bson.D{})

	if err != nil {
		return fmt.Errorf("error adding item %s: %v", expense.Item, err)
	}

	expense.ItemID = int(count) + 1
	expense.TimeStamp = time.Now()
	_, err = g.Collection.InsertOne(context.TODO(), expense)

	if err != nil {
		return fmt.Errorf("error adding item %s: %v", expense.Item, err)
	}

	return nil
}

func (g *Group) AddItems(expenses []models.Expense) error {

	for _, e := range expenses {
		err := g.AddItem(e)

		if err != nil {
			return err
		}
	}

	return nil

}
