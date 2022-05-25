package store

type Config struct {
	Store struct {
        DatabaseURL  string `json:"database_url"`
        } `json:"store"`
}

func NewConfig() *Config {
	return &Config{}
}