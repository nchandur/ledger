package api

import (
	"fmt"
	"ledger/db"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GroupsHandler(r *gin.Engine) {
	r.GET("/groups/view", func(ctx *gin.Context) {
		database := db.Client.Database("ledgers")

		res, err := database.ListCollectionNames(ctx, bson.D{})

		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "index.html", gin.H{"message": []string{"Error fetching ledgers"}})
			return
		}

		fmt.Println(res)

		ctx.HTML(http.StatusOK, "index.html", gin.H{"message": res})
	})

	r.POST("/groups", func(ctx *gin.Context) {
		groupName := ctx.PostForm("groupName")
		participants := ctx.PostFormArray("participants")
		currency := ctx.PostForm("currency")

		participantStr := url.QueryEscape(strings.Join(participants, ","))

		redirectURL := fmt.Sprintf("/groups?groupName=%s&participants=%s&currency=%s",
			groupName,
			participantStr,
			url.QueryEscape(currency),
		)

		ctx.Redirect(http.StatusSeeOther, redirectURL)
	})

	r.GET("/groups", func(ctx *gin.Context) {
		groupName := ctx.Query("groupName")
		participantsRaw := ctx.Query("participants")
		currency := ctx.Query("currency")

		participants := strings.Split(participantsRaw, ",")

		ctx.HTML(http.StatusOK, "groups.html", gin.H{
			"groupName":    groupName,
			"participants": participants,
			"currency":     currency,
		})
	})

}
