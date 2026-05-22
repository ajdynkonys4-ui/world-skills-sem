package modules

import (
	"database/sql"
	"log"
	"time"
	"worldskills/backend/model"
	"worldskills/backend/utils"

	"github.com/gin-gonic/gin"
)

func GetIINData(ctx *gin.Context, conn *sql.DB) {
	var dataIIN model.IINModel
	err := ctx.ShouldBindJSON(&dataIIN)
	if err != nil {
		log.Println("Не удалось распарсить данные:", err)
		return
	}
	//Валидация входящих данных
	validation := utils.Validation(dataIIN.Iin)
	//Пробрасывание данных с безопасным контекстом
	_, err = conn.ExecContext(ctx, "INSERT INTO people_iin7 (iin, status, created_at) VALUES($1, $2, $3)", &dataIIN.Iin, validation, time.Now())
	if err != nil {
		log.Println("Не получилось добавить данные в базу данных:", err)
		ctx.JSON(500, gin.H{"status": "error", "details": err.Error()})
		return
	}
	if err == nil {
		log.Println("Получилось добавить данные в базу")
		return
	}

	ctx.JSON(200, gin.H{"status": "success"})
}

func DataRead(ctx *gin.Context, conn *sql.DB) {
	rows, err := conn.QueryContext(ctx, "SELECT * FROM people_iin7")
	if err != nil {
		log.Println("Не удалось взять данные из базы:", err)
		return
	}
	defer rows.Close()
	var dataIIN model.IINModel
	var arrDataIIN []model.IINModel
	for rows.Next() {
		err := rows.Scan(&dataIIN.ID, &dataIIN.Iin, &dataIIN.Status, &dataIIN.Created_at)
		arrDataIIN = append(arrDataIIN, dataIIN)
		if err != nil {
			log.Println("Не удалось достать данные из базы:", err)
			ctx.JSON(500, gin.H{"status": "error", "details": err.Error()})
			return
		}
	}
	ctx.JSON(200, gin.H{"data": arrDataIIN})
}
