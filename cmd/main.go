package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexKomzzz/zip-test/pkg/handler"
	"github.com/AlexKomzzz/zip-test/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	port = ":9090"
)

func main() {

	service := service.NewService()
	handler := handler.NewHandler(service)

	server := handler.InitRoutes()
	/*if err != nil {
		logrus.Fatalf("Error init server: %s", err.Error())
		return
	}*/

	go func() {
		if err := server.Run(port); err != nil {
			logrus.Fatalf("Error run web serv")
			return
		}
	}()

	logrus.Print("Server Started")

	// остановка сервера
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Server Stopted")
	logrus.Print("Server Stopted")
	logrus.Print("Server Stopted")
	logrus.Print("Server Stopted")
	logrus.Print("Server Stopted")
	logrus.Print("Server Stopted")

}
