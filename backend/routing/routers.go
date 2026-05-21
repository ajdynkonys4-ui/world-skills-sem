package routing

import (
	"database/sql"
	"worldskills/backend/internal/modules"

	"github.com/gin-gonic/gin"
)

// Маршруты и роутинги
func Routers(serv *gin.Engine, conn *sql.DB) {
	serv.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"title": "website",
		})
	})
	serv.POST("/api/iin", func(ctx *gin.Context) {
		modules.GetIINData(ctx, conn) //Роут на взятие ИИН
	})
}
