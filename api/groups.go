package api

import (
	"fmt"
	"ledger/crud"
	"ledger/db"
	"ledger/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateGroupHandler(r *gin.Engine) {
	r.POST("/", func(ctx *gin.Context) {
		groupName := ctx.PostForm("groupName")
		currency := ctx.PostForm("currency")
		participants := ctx.PostFormArray("participants")

		err := crud.CreateGroup(groupName, participants, currency)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to create group: %v", err.Error())})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Group created successfully!"})

	})
}

func ViewGroupsHandler(r *gin.Engine) {
	r.GET("/view", func(ctx *gin.Context) {
		cur, err := db.Client.Database("ledgers").Collection("groups").Find(ctx, bson.D{})

		if err != nil {
			ctx.HTML(http.StatusBadRequest, "index.html", gin.H{"error": err.Error()})
			return
		}

		groups := []models.Group{}

		for cur.Next(ctx) {
			var group models.Group

			err = cur.Decode(&group)

			if err != nil {
				ctx.HTML(http.StatusBadRequest, "index.html", gin.H{"error": err.Error()})
				return
			}

			groups = append(groups, group)

		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{"groups": groups})

	})

}

func GroupHandler(r *gin.Engine) {
	r.GET("/group", func(ctx *gin.Context) {

		groupName := strings.TrimSpace(ctx.Query("name"))

		collection := db.Client.Database("ledgers").Collection("groups")

		var group models.Group

		err := collection.FindOne(ctx, bson.D{{Key: "group_name", Value: groupName}}).Decode(&group)

		if err != nil {
			ctx.HTML(http.StatusBadRequest, "groups.html", gin.H{"error": err.Error()})
			return
		}

		collection = db.Client.Database("ledgers").Collection(groupName)

		cur, err := collection.Find(ctx, bson.D{})

		if err != nil {
			ctx.HTML(http.StatusNoContent, "groups.html", gin.H{"error": err.Error()})
			return
		}

		expenses := []models.Expense{}

		for cur.Next(ctx) {
			var expense models.Expense
			cur.Decode(&expense)
			expenses = append(expenses, expense)
		}

		ctx.HTML(http.StatusOK, "groups.html", gin.H{
			"name":     groupName,
			"expenses": expenses,
		})
	})

	r.POST("/group", func(ctx *gin.Context) {
		action := ctx.PostForm("_action")
		groupName := ctx.PostForm("name")

		group, err := crud.AccessGroup(groupName)

		fmt.Println(groupName, action)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		switch action {
		case "add":
			item := strings.TrimSpace(ctx.PostForm("item"))
			priceTxt := strings.TrimSpace(ctx.PostForm("price"))
			lent := strings.TrimSpace(ctx.PostForm("lent"))
			involved := ctx.PostFormArray("participants")
			splitType := ctx.PostForm("splitType")

			price, err := strconv.ParseFloat(priceTxt, 64)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			expense := models.Expense{Item: item, Price: price, Lent: lent, Involved: involved, SplitType: splitType}

			err = group.AddItem(expense)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

		case "delete":
			itemID := strings.TrimSpace(ctx.PostForm("id"))

			id, err := strconv.Atoi(itemID)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			err = group.RemoveItemByID(id)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/group?name=%s", groupName))
	})

}
