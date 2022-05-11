package main

import (
	"log"
	"github.com\ALPetrov\my-rest-api\internal\app\apiserver"

)
func main() {
	config := apiserver.NewConfig() //инициализируем NewConfig с параметрами по умолчанию 
	
	s := apiserver.New(config)   //передаем config в качестве аргумента, вызываем функц Старт и запускает его 	записываем АпиСервер и функц Старт запускает его 
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}


}
