package handler

import (
	"github.com/AlexKomzzz/zip-test/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		service: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine { // Инициализация групп функций мультиплексора

	//gin.SetMode(gin.ReleaseMode) // Переключение сервера в режим Релиза из режима Отладка

	mux := gin.New()

	mux.LoadHTMLGlob("web/templates/*.html")
	mux.NoRoute(Response404) // При неверном URL вызывает ф-ю Response404

	mux.Static("/assets", "./web/assets")

	// Стартовая страница
	mux.StaticFile("/", "./web/templates/start_list.html")

	// работа с файлами
	files := mux.Group("/files")
	{
		// получение файла от клиента
		files.POST("/get-file", h.getFile)
		// отправка данных клиету
		files.GET("/get-data", h.dataClient)
	}

	return mux
}
