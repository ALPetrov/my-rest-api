package store_test

import (
	"testing"
	"github.com/ALPetrov/my-rest-api/internal/app/model"
	"github.com/ALPetrov/my-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)
func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	
	assert.NoError(t, err)  //Проверяет нет ли ошибки
	assert.NotNil(t, u)    //Проверяет чтобы юзер был не нулевой
}
func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@Example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(model.TestUser(t))
	
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

