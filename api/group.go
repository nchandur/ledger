package api

import (
	"ledger/crud"
	"ledger/db"
	"ledger/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GroupsHandler(r *gin.Engine) {
	r.POST("/groups/create", func(ctx *gin.Context) {
		groupRequest := struct {
			Name     string   `bson:"name" json:"name"`
			People   []string `bson:"people" json:"people"`
			Currency string   `bson:"currency" json:"currency"`
		}{}

		if err := ctx.ShouldBindJSON(&groupRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": nil})
			return
		}

		err := crud.CreateGroup(groupRequest.Name, groupRequest.People, groupRequest.Currency)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": nil})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"body": "group created", "error": nil})
	})

	r.GET("/groups/view", func(ctx *gin.Context) {
		collection := db.Client.Database("ledgers").Collection("groups")

		name := ctx.Query("name")

		filter := bson.D{{}}

		if name != "" {
			filter = bson.D{{Key: "group_name", Value: name}}
		}

		cur, err := collection.Find(ctx, filter)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": nil})
			return
		}

		defer cur.Close(ctx)

		groups := []models.Group{}

		for cur.Next(ctx) {
			var group models.Group

			err := cur.Decode(&group)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": nil})
				return
			}

			groups = append(groups, group)
		}

		ctx.JSON(http.StatusOK, gin.H{"error": nil, "body": gin.H{"groups": groups}})
	})

	r.DELETE("/groups/delete", func(ctx *gin.Context) {
		name := ctx.Query("name")

		if name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing name parameter", "body": nil})
			return
		}

		collection := db.Client.Database("ledgers").Collection("groups")

		res, err := collection.DeleteOne(ctx, bson.M{"group_name": name})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
			return
		}

		if res.DeletedCount == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no group found with given name", "body": nil})
			return
		}

		db := db.Client.Database(name)

		if err := db.Drop(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to drop ledger", "body": nil})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"error": nil, "body": "group dropped"})

	})

}
