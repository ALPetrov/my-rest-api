package store

import "github.com/ALPetrov/my-rest-api/internal/app/model"


type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users(email, encrypted_password) VALUES (?, ?) SELECT LAST_INSERT_ID()",
		u.Email, 
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {   // Запрос должен возвращать ID нашего пользователя (уточнить синтаксис запроса у Дениса)
	return nil, err
}
	return u, nil
}	
// используем при авторизации пользователя по email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}