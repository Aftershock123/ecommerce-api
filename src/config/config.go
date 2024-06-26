// config/config.go
package config

type Config struct {
	DB     DBConfig
	Server ServerConfig
	// Add more configurations as needed
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type ServerConfig struct {
	Port string
}

func LoadConfig() Config {
	return Config{
		DB: DBConfig{
			Username: "root",
			Host:     "localhost",
			Port:     "3307",
			Database: "ecommerce",
		},
		Server: ServerConfig{
			Port: "8080",
		},
	}
}
