package api

import (
	"ledger/crud"
	"ledger/db"
	"ledger/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func LedgerHandler(r *gin.Engine) {
	r.GET("/ledger/view", func(ctx *gin.Context) {
		name := ctx.Query("name")

		collection := db.Client.Database("ledgers").Collection(name)

		cur, err := collection.Find(ctx, bson.D{})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
			return
		}

		expenses := []models.Expense{}

		for cur.Next(ctx) {
			var expense models.Expense

			err := cur.Decode(&expense)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
				return
			}

			expenses = append(expenses, expense)

		}

		ctx.JSON(http.StatusOK, gin.H{"error": nil, "body": gin.H{"expenses": expenses}})
	})

	r.POST("/ledger/add", func(ctx *gin.Context) {

		name := ctx.Query("name")

		group, err := crud.AccessGroup(name)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": nil})
			return
		}

		expense := models.Expense{}

		if err := ctx.ShouldBindJSON(&expense); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": nil})
			return
		}

		err = group.AddItem(expense)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
			return
		}

		err = crud.CalculateNetAndTransactions(name)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"error": nil, "body": "expense added"})

	})

	r.DELETE("/ledger/delete", func(ctx *gin.Context) {
		name := ctx.Query("name")
		item := ctx.Query("item")

		group, err := crud.AccessGroup(name)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": nil})
			return
		}

		err = group.RemoveItem(item)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
			return
		}

		err = crud.CalculateNetAndTransactions(name)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"error": nil, "body": "expense deleted"})

	})

	r.PATCH("/ledger/update", func(ctx *gin.Context) {
		name := ctx.Query("name")
		item := ctx.Query("item")

		group, err := crud.AccessGroup(name)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": nil})
			return
		}

		update := make(map[string]any)

		if err := ctx.ShouldBindJSON(&update); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": nil})
			return
		}

		err = group.UpdateItem(item, update)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
			return
		}

		err = crud.CalculateNetAndTransactions(name)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"error": nil, "body": "expense updated"})

	})

	r.GET("/ledger/balances", func(ctx *gin.Context) {
		name := ctx.Query("name")

		filter := bson.D{{Key: "group", Value: name}}

		collection := db.Client.Database("ledgers").Collection("balances")

		balance := models.Balance{Balances: make(map[string]float64)}

		err := collection.FindOne(ctx, filter).Decode(&balance)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"error": nil, "body": gin.H{"balances": balance}})

	})

	r.GET("/ledger/transactions", func(ctx *gin.Context) {
		name := ctx.Query("name")

		filter := bson.D{{Key: "group", Value: name}}

		collection := db.Client.Database("ledgers").Collection("transactions")

		transactions := bson.M{}

		err := collection.FindOne(ctx, filter).Decode(&transactions)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "body": nil})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"error": nil, "body": gin.H{"transactions": transactions}})

	})

}
