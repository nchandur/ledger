package crud

import (
	"context"
	"fmt"
	"ledger/db"
	"ledger/models"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Group struct {
	Collection *mongo.Collection
}

func Exists(groupName string) (bool, error) {

	res, err := db.Client.Database("ledgers").Collection("groups").CountDocuments(context.Background(), bson.D{{Key: "group_name", Value: strings.TrimSpace(groupName)}})

	if err != nil {
		return false, err
	}

	return res >= 1, nil

}

func AccessGroup(groupName string) (Group, error) {
	exists, err := Exists(groupName)

	if err != nil {
		return Group{}, err
	}

	if exists {
		return Group{Collection: db.Client.Database("ledgers").Collection(groupName)}, nil
	}

	return Group{}, fmt.Errorf("group does not exist")
}

func CreateGroup(groupName string, people []string, currency string) error {

	exists, err := Exists(groupName)

	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("group already exists!")
	}

	g := Group{Collection: db.Client.Database("ledgers").Collection("groups")}

	group := models.Group{TimeStamp: time.Now().In(time.Local), GroupName: groupName, People: people, Currency: currency}

	_, err = g.Collection.InsertOne(context.TODO(), group)

	if err != nil {
		return fmt.Errorf("error creating group %s: %v", group.GroupName, err)
	}
	return nil
}

func (g *Group) Delete() error {

	exists, err := Exists(g.Collection.Name())

	if err != nil {
		return err
	}

	if exists {
		if err := g.Collection.Drop(context.TODO()); err != nil {
			return fmt.Errorf("failed to drop collection: %v", err)
		}
	}

	fmt.Println("group deleted!")

	return nil
}
