package service

import (
	"mime/multipart"
	"os"
)

type FileHanding interface {
	// сохранение файла на сервере
	SaveFileToServe(file multipart.File) (*os.File, error)
}

type Service struct {
	FileHanding
}

func NewService() *Service {
	return &Service{
		FileHanding: NewFileService(),
	}
}
