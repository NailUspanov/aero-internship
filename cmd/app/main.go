package main

import (
	"aero-internship/internal/app"
	"aero-internship/pkg/config"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/xlab/closer"
)

func main() {

	closer.Bind(func() {
		logrus.Print("Stop running...")
	})

	// инициализация env конфига
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// функция, которая вызовется при получении приложением сигнала об остановке

	cfg := config.NewConfig()
	application, err := app.NewApp(cfg)
	if err != nil {
		logrus.Fatalf("Can't create new app.App object: %v", err)
	}
	go application.Run()
	// ждем в основной горутине сигнал
	closer.Hold()
}
