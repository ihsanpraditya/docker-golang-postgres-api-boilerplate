package config

type Config struct {
    DB DatabaseConfig
}

func LoadConfig() *Config {
    return &Config{
        DB: LoadDatabaseConfig(),
    }
}
