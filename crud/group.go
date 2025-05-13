package crud

import (
	"context"
	"fmt"
	"ledger/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Group struct {
	Collection *mongo.Collection
}

func Exists(groupName string) (bool, error) {

	coll, err := db.Client.Database("ledgers").ListCollectionNames(context.Background(), bson.D{{Key: "name", Value: groupName}})

	if err != nil {
		return false, err
	}

	return len(coll) == 1, nil

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

func CreateGroup(groupName string) (Group, error) {

	exists, err := Exists(groupName)

	if err != nil {
		return Group{}, err
	}

	if exists {
		return Group{}, fmt.Errorf("group already exists!")
	}

	g := Group{Collection: db.Client.Database("ledgers").Collection(groupName)}
	return g, nil
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
