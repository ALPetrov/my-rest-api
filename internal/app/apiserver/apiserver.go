package apiserver

import (
	//"io"
	"io"
	"net/http"
	"encoding/json"
	"os"
	"github.com/ALPetrov/my-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//1-Создаем структуру сервера и добавляем поля с конфигурацией, логом состояния подключения, конфигурацией роутера
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store

}
//2-Инициализируем сервер через функцию New, подготавливаем и конфигурируем поля
func New (config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}
//3-Запускаем HTTP server, подключаемся к БД и т.д.
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {    // Запускает функцию, если ошибка выводит error
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Сервер запущен")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
// 3.1 Конфигурирует поле Logger
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)  //Запрашивает состояние уровня ошибки
	if err != nil {
		return nil
	}
	s.logger.SetLevel(level)
	return nil
}
//3.2 конфигурирует поле Router
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello",s.handleHello()) 
}
func (s *APIServer) handleHello() http.HandlerFunc {  // запускает http сервер, посылает запрос, получает ответ и выводит его на страницу
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "You are really cool man")
	}
} 
//Конфигурируем и открываем хранилище
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st  // Если всё удачно, записываем переменную в s.store

	return nil
}
//Функция считывает из файла и конвертирует формат json в string
func ParseBaseConfig(filename string) (*Config, error) {
	var config *Config
	configFile, err := os.Open(filename)

	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}