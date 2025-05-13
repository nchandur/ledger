package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "UPDATE", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	LandingHandler(r)
	CreateGroupHandler(r)
	DeleteGroupHandler(r)

	AddItemHandler(r)
	RetrieveItemHandler(r)
	UpdateItemByIDHandler(r)
	RemoveItemByIDHandler(r)

	GetBalancesHandler(r)
	GetTransactionsHandler(r)

	return r

}
