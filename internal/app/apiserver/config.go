package apiserver

// 1.1 Структура с полем IP адреса сервера и полем уровня состояния ошибки
type Config struct {
	BindAddr string  `json:"bind_addr"`
	LogLevel string   `json:"log_level"`
}
// Инициализируем Config c начальными параметрами по умолчанию
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}