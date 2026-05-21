package main

import (
	"log"
	"worldskills/backend/database"
	"worldskills/backend/routing"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	//Подключение к базе данных
	conn := database.Connection()
	_, err := conn.Exec("CREATE TABLE IF NOT EXISTS people_iin1 (id BIGINT GENERATED ALWAYS AS IDENTITY, iin_value INTEGER, created_at TIMESTAMP)")
	if err != nil {
		log.Println("Не удалось создать таблицу в базе:", err)
		return
	}
	defer conn.Close() //Закрываем соединение с базой данных
	server.LoadHTMLGlob("frontend/template/*")
	server.Static("/static", "./static")
	routing.Routers(server, conn) //Подключение апи
	server.Run(":10000")
}
