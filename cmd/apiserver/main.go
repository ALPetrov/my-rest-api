package main

import (
	"encoding/json"
	"log"
	"github.com/ALPetrov/my-rest-api/internal/app/apiserver"
	"os"
)
func main() {
	config, _ := ParseBaseConfig("C:/Users/Admin/go/src/github.com/ALPetrov/my-rest-api/configs/apiserver.json")
	
	s := apiserver.New(config)   //передаем config в качестве аргумента
	
	if err := s.Start(); err != nil {  //вызываем функц Старт и запускаем его
	
		log.Fatal(err)                // если что- то не так, выводит ошибку  
	}
}

func ParseBaseConfig(filename string) (*apiserver.Config, error) {
	var config *apiserver.Config

	configFile, err := os.Open(filename)

	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, err
}