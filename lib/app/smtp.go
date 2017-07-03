package app

// SMTP smtp
type SMTP struct {
	Host     string   `toml:"host"`
	Port     int      `toml:"port"`
	Bcc      []string `toml:"bcc"`
	From     string   `toml:"from"`
	Password string   `toml:"password"`
}
