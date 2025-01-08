package server

import (
	"github.com/Ilya-c4talyst/gin_study/handlers"

	"github.com/gin-gonic/gin"
)

func InitRotes() {
	// Инициализация  роута (по умолчанию)
	router := gin.Default()
	// Создание заметки
	router.POST("/note", handlers.CreateNoteHandler)
	// Удаление заметки
	router.DELETE("/note/:id", handlers.DeleteNoteHandler)
	// Получение заметки
	router.GET("/note/:id", handlers.GetNoteHandler)
	// Редактирование заметки
	router.PUT("/note/:id", handlers.UpdateNoteHandler)
	// Получение списка всех заметок
	router.GET("/notes", handlers.GetNotesHandler)

	router.Run(":9100")
}
