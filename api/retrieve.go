package api

import (
	"ledger/crud"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RetrieveItemHandler(r *gin.Engine) {
	r.GET("/group/item", func(ctx *gin.Context) {
		id := strings.TrimSpace(ctx.Query("item_id"))
		groupName := ctx.Query("group")

		group, err := crud.AccessGroup(groupName)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		itemID, err := strconv.Atoi(id)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		expense, err := group.RetrieveItemByID(itemID)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"body":  nil,
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"body":  expense,
			"error": nil,
		})

	})
}
