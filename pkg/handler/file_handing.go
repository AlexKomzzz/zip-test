package handler

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver/v4"
	"github.com/sirupsen/logrus"
)

// получение файла от клиента
func (h *Handler) getFile(c *gin.Context) {

	// получение файла из запроса от клиента по ключевому слову "file"
	formFile, _, err := c.Request.FormFile("file")
	if err != nil {
		logrus.Println("ошибка при получении файла от клиента: ", err)
		errorServerResponse(c, err)
		return
	}
	defer formFile.Close()

	_, err = h.service.SaveFileToServe(formFile)
	if err != nil {
		logrus.Println("Handler/getFile()/ошибка при сохранении файла в системе: " + err.Error())
		errorServerResponse(c, err)
		return
	}

	// fs, err := archiver.FileSystem("./web/files/")
	files, err := archiver.FilesFromDisk(nil, map[string]string{
		"./web/files/123.pdf": "123.pdf",
	})
	if err != nil {
		errorServerResponse(c, err)
		return
	}

	out, err := os.Create("example.zip")
	if err != nil {
		errorServerResponse(c, err)
		return
	}
	defer out.Close()

	format := archiver.CompressedArchive{
		Compression: archiver.Gz{},
		Archival:    archiver.Zip{},
	}

	err = format.Archive(context.Background(), out, files)
	if err != nil {
		errorServerResponse(c, err)
		return
	}

	c.File("example.zip")
	// c.FileAttachment("example.zip", "new_file")
	logrus.Println("Файл отправлен")
}

// отправка данных клиенту
func (h *Handler) dataClient(c *gin.Context) {

}
