package store

import (
	"database/sql"
	_ "github.com/lib/pq"

)
//4-Создаем структуру Базы данных, инициализируем поля
type Store struct {
	config 			*Config
	db    			*sql.DB
	userRepository  *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//5-Создаем метод для инициализации БД, подключения к БД и отлавливания ошибок
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
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
// Метод через который пользователь взаимодействует с БД
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository    //Проверяем на ошибки и возвращаем данные
	}
	s.userRepository = &UserRepository{     //Если не создан, инициализируем и возвражаем данные
		store: s,
	}
	return s.userRepository
}