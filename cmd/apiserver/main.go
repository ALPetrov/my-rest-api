package main

import (
	
	"fmt"
	"log"
	"github.com/ALPetrov/my-rest-api/internal/app/apiserver"
)


func main() {
	fmt.Println("\n Запускаем программу...")

	config, _ := apiserver.ParseBaseConfig("C:/Users/Admin/go/src/github.com/ALPetrov/my-rest-api/configs/apiserver.json")

	s := apiserver.New(config)           //передаем config в качестве аргумента
	if err := s.Start(); err != nil {    //вызываем функц Старт и запускаем его
		log.Fatal(err)                  // если что- то не так, выводит ошибку
	}
}
