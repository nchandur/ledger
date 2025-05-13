package crud

import (
	"context"
	"fmt"
	"ledger/db"
	"ledger/models"
	"ledger/split"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CalculateNetAndTransactions(groupName string) error {
	group, err := AccessGroup(groupName)

	if err != nil {
		return err
	}

	filter := bson.M{
		"item_id": bson.M{"$ne": 0},
	}

	cursor, err := group.Collection.Find(context.TODO(), filter)

	if err != nil {
		return err
	}

	defer cursor.Close(context.TODO())

	people := make(models.People)

	for cursor.Next(context.TODO()) {
		var expense models.Expense

		if err := cursor.Decode(&expense); err != nil {
			return err
		}

		switch expense.SplitType {
		case "equal":
			err = split.SplitEqual(expense, people)
		case "manual":
			err = split.SplitManual(expense, people)
		case "percentage":
			err = split.SplitPercentages(expense, people)
		}

		if err != nil {
			return fmt.Errorf("item: %s, %v", expense.Item, err)
		}

	}

	if err = cursor.Err(); err != nil {
		return err
	}

	balances := db.Client.Database("ledgers").Collection("balances")

	filter = bson.M{"group": groupName}

	update := bson.M{
		"$set": bson.M{
			"group":      groupName,
			"updated_at": time.Now(),
			"balances":   people,
		},
	}

	updateOpts := options.Update().SetUpsert(true)

	_, err = balances.UpdateOne(context.TODO(), filter, update, updateOpts)

	if err != nil {
		return err
	}

	transac := split.SimplifyDebts(people)

	transactions := db.Client.Database("ledgers").Collection("transactions")

	update = bson.M{
		"$set": bson.M{
			"group":        groupName,
			"updated_at":   time.Now(),
			"transactions": transac,
		},
	}

	transacOpts := options.Update().SetUpsert(true)

	_, err = transactions.UpdateOne(context.TODO(), filter, update, transacOpts)

	if err != nil {
		return err
	}

	return nil
}
