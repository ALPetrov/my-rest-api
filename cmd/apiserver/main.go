package main

import (
	//	"log"
	"encoding/json"
	"fmt"
	"os"
//	"github.com/ALPetrov/my-rest-api/configs"
//	"github.com/ALPetrov/my-rest-api/internal/app/apiserver"
)
type Config struct {
	BindAddr string `json:"bind_addr"`
    LogLevel string `json:"log_level"`
    Store struct {
        Database  string `json:"database_url"`
        } `json:"store"`
}

func main() {
	fmt.Println("Запускаем программу...")
	config, _ := LoadBaseConfig("apiserver.json")
	fmt.Println(config)

	//config := apiserver.NewConfig() //инициализируем NewConfig с параметрами по умолчанию 
	


	// s := apiserver.New(config)   //передаем config в качестве аргумента
	// if err := s.Start(); err != nil {  //вызываем функц Старт и запускаем его
	// 	log.Fatal(err)                // если что- то не так, выводит ошибку  
	// }
}
func LoadBaseConfig(filename string) (Config, error) {
    var config Config
    configFile, err := os.Open(filename)

    defer configFile.Close()

    if err != nil {
        return config, err
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config, err
}