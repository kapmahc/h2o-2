package app

import (
	"database/sql"
	"fmt"
)

// PostgreSQL postgresql
type PostgreSQL struct {
	Host     string `toml:"host"`
	DbName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Port     int    `toml:"port"`
	SslMode  string `toml:"sslmode"`
}

func (p *PostgreSQL) String() string {
	return fmt.Sprintf(
		"postgres://%s@%s:%d/%s?sslmode=%s",
		p.User,
		p.Host,
		p.Port,
		p.DbName,
		p.SslMode,
	)
}

// Open open
func (p *PostgreSQL) Open() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			p.Host,
			p.Port,
			p.User,
			p.DbName,
			p.Password,
			p.SslMode,
		),
	)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
