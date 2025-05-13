package api

import (
	"fmt"
	"ledger/crud"
	"ledger/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGroupHandler(r *gin.Engine) {
	r.GET("/group/create", func(ctx *gin.Context) {
		name := ctx.Query("group")

		err := crud.CreateGroup(name)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"body":  fmt.Sprintf("group %s created", name),
			"error": nil,
		})

	})
}

func AddItemHandler(r *gin.Engine) {
	r.POST("/group/item/add", func(ctx *gin.Context) {
		groupName := ctx.Query("group")

		group, err := crud.AccessGroup(groupName)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
		}

		var expense models.Expense

		if err := ctx.BindJSON(&expense); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"body":  nil,
				"error": "invalid JSON: " + err.Error(),
			})
			return
		}

		err = group.AddItem(expense)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"body":  fmt.Sprintf("item %s added to group %s", expense.Item, groupName),
			"error": nil,
		})

	})
}
