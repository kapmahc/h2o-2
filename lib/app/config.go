package app

// Server server
type Server struct {
	Name  string `toml:"name"`
	Port  int    `toml:"port"`
	Theme string `toml:"theme"`
	Ssl   bool   `toml:"ssl"`
}

// Secrets secrets
type Secrets struct {
	Hmac   string `toml:"hmac"`
	Aes    string `toml:"aes"`
	Jwt    string `toml:"jwt"`
	Cookie string `toml:"cookie"`
	Csrf   string `toml:"csrf"`
}

// Config config
type Config struct {
	Server     Server          `toml:"server"`
	Plugins    map[string]bool `toml:"plugins"`
	Secrets    Secrets         `toml:"secrets"`
	PostgreSQL PostgreSQL      `toml:"postgresql"`
	RabbitMQ   RabbitMQ        `toml:"rabbitmq"`
	Redis      Redis           `toml:"redis"`
	SMTP       SMTP            `toml:"smtp"`
}
