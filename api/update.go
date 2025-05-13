package api

import (
	"ledger/crud"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateItemByIDHandler(r *gin.Engine) {
	r.PUT("/group/item", func(ctx *gin.Context) {
		groupName := ctx.Query("group")

		group, err := crud.AccessGroup(groupName)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		id := ctx.Query("item_id")
		itemID, err := strconv.Atoi(id)

		var update map[string]any

		if err := ctx.ShouldBindJSON(&update); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error(), "body": nil})
			return
		}

		err = group.UpdateItemByID(itemID, update)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		err = crud.CalculateNetAndTransactions(groupName)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"body":  "updated item",
			"error": nil,
		})

	})
}
