package api

import (
	"strings"
	"text/template"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"join": strings.Join,
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

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
	ViewGroupsHandler(r)
	GroupHandler(r)

	return r

}
