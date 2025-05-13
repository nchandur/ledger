package api

import (
	"fmt"
	"ledger/crud"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteGroupHandler(r *gin.Engine) {
	r.DELETE("/group", func(ctx *gin.Context) {
		name := ctx.Query("group")

		group, err := crud.AccessGroup(name)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		err = group.Delete()

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"body":  fmt.Sprintf("%s deleted successfully", group.Collection.Name()),
			"error": nil,
		})

	})
}

func RemoveItemByIDHandler(r *gin.Engine) {
	r.DELETE("/group/item", func(ctx *gin.Context) {
		name := ctx.Query("group")

		group, err := crud.AccessGroup(name)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		id := ctx.Query("item_id")
		itemID, err := strconv.Atoi(id)

		err = group.RemoveItemByID(itemID)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"body":  "removed item from group",
			"error": nil,
		})

	})
}
