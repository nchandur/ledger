package api

import (
	"fmt"
	"ledger/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBalancesHandler(r *gin.Engine) {
	r.GET("/group/balances", func(ctx *gin.Context) {
		name := ctx.Query("group")

		filter := bson.M{
			"group": name,
		}

		var res bson.M

		err := db.Client.Database("ledgers").Collection("balances").FindOne(ctx, filter).Decode(&res)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": fmt.Errorf("no group balances found"),
			})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"body":  res,
			"error": nil,
		})

	})
}
