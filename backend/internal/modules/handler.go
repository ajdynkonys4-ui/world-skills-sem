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
	//Добавления ИИН в код
	iin, err := strconv.Atoi(dataIIN.Iin)
	if err != nil {
		log.Println("Нельзя вводить символы")
		ctx.JSON(400, gin.H{"status": "error", "details": err.Error()})
		return
	}

	//Пробрасывание данных с безопасным контекстом
	_, err = conn.ExecContext(ctx, "INSERT INTO people_iin1 (iin_value, created_at) VALUES ($1, $2)", iin, time.Now())
	if err != nil {
		log.Println("Не получилось добавить данные в базу данных:", err)
		ctx.JSON(500, gin.H{"status": "error", "details": err.Error()})
		return
	}
	if err == nil {
		log.Println("Получилось добавить данные в базу")
		return
	}
}
func DataRead(ctx *gin.Context, conn *sql.DB) {
	rows, err := conn.QueryContext(ctx, "SELECT * FROM people_iin1")
	if err != nil {
		log.Println("Не удалось взять данные из базы:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var dataIIN model.IINModel
		if err := rows.Scan(&dataIIN.ID, &dataIIN.Iin, dataIIN.Created_at); err != nil {
			ctx.JSON(200, gin.H{"data": dataIIN})
		}
	}
}
