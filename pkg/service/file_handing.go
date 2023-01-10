package service

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
)

// название файла
const fileName = "test.txt"

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

// сохранение файла на сервере
func (service *FileService) SaveFileToServe(file multipart.File) (*os.File, error) {

	// создание пустого файла в файловой системе
	saveFile, err := os.OpenFile("./web/files/"+"123.pdf", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, errors.New("FileService/SaveFileToServe()/ ошибка при создании файла в системе: " + err.Error())
	}
	defer saveFile.Close()

	// копирование данных из запроса в созданный файл
	io.Copy(saveFile, file)
	return saveFile, nil
}

// открыть файл
/*
// открытие файла
f, err := excelize.OpenFile("./web/files/" + fileName)
if err != nil {
	logrus.Println("ошибка при открытии файла из системы: ", err)
	errorServerResponse(c, err)
	return
}
defer func() {
	// Close the spreadsheet.
	if err := f.Close(); err != nil {
		logrus.Println("ошибка при закрытии файла: ", err)

	}
}()


}*/
