package apiserver

import (
	//"io"
	//"net/http"

	"github.com/sirupsen/logrus"
)

//1-Создаем структуру сервера и добавляем поля с конфигурацией, логом состояния подключения, конфигурацией роутера
type APIServer struct {
	config *Config
	logger *logrus.Logger


}
//2-Инициализируем сервер через функцию New, подготавливаем и конфигурируем поля
func New (config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}
//3-Запускает HTTP server, подключается к БД и т.д.
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {    // Запускаем функцию, если ошибка выводим error
		return err
	}
	s.logger.Info("Сервер запущен")
	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)  //Запрашивает состояние уровня ошибки
	if err != nil {
		return nil
	}
	s.logger.SetLevel(level)
	return nil
}

