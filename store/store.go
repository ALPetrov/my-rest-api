package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)
//4-Создаем структуру Базы данных, инициализируем поля
type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}
//5-Создаем метод для инициализации БД, подключения к БД и отлавливания ошибок
func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {  //Пингуем нашу БД для проверки соединения
		return err
	}

	s.db = db   // Если ошибок нет, записываем переменную в s.db

	return nil
}
//5.1-Создаем метод для отключения от БД
func (s *Store) Close() {
	s.db.Close()

}