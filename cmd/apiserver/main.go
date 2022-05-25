package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ALPetrov/my-rest-api/internal/app/apiserver"
)

func main() {
	fmt.Println("Запускаем программу...")
	config, _ := ParseBaseConfig("apiserver.json")

	// config := apiserver.NewConfig() //инициализируем NewConfig с параметрами по умолчанию 
	
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