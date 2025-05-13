package api

import (
	"fmt"
	"ledger/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTransactionsHandler(r *gin.Engine) {
	r.GET("/group/transactions", func(ctx *gin.Context) {
		name := ctx.Query("group")

		filter := bson.M{
			"group": name,
		}

		var res bson.M

		err := db.Client.Database("ledgers").Collection("transactions").FindOne(ctx, filter).Decode(&res)

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
				"error": fmt.Errorf("no group transactions found"),
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
