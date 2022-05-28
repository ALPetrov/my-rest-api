package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper() // Вызов метода Helper указывает что, что TestStore тестовый метод 
	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config) 
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}
	// Возвращаем функцию в которой в качестве теста склеиваем массив таблицы с запятой.
	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); 
			err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}
}