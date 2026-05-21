package modules

import (
	"database/sql"
	"log"
	"strconv"
	"time"
	"worldskills/backend/model"

	"github.com/gin-gonic/gin"
)

func GetIINData(ctx *gin.Context, conn *sql.DB) {
	var dataIIN model.IINModel
	err := ctx.ShouldBindJSON(&dataIIN)
	if err != nil {
		log.Println("Не удалось распарсить данные:", err)
		return
	}
	iin, _ := strconv.Atoi(dataIIN.Iin)
	//Пробрасывание данных с безопасным контекстом
	_, err = conn.ExecContext(ctx, "INSERT INTO people_iin1 (iin_value, created_at) VALUES ($1, $2)", iin, time.Now())
	if err != nil {
		log.Println("Не получилось добавить данные в базу данных:", err)
		ctx.JSON(500, gin.H{"status": "error", "details": err.Error()})
		return
	}
	if err == nil {
		log.Println("Получилось добавить данные в базу")
	}
}
