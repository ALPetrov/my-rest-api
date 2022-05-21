package main

import (
	"log"
	"github.com/ALPetrov/my-rest-api/internal/app/apiserver"

)
func main() {
	config := apiserver.NewConfig() //инициализируем NewConfig с параметрами по умолчанию 
	
	s := apiserver.New(config)   //передаем config в качестве аргумента
	if err := s.Start(); err != nil {  //вызываем функц Старт и запускаем его
		log.Fatal(err)                // если что- то не так, выводит ошибку  
	}


}