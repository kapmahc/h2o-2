package app_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/kapmahc/h2o/lib/app"
	_ "github.com/kapmahc/h2o/plugins/site"
)

func TestConfig(t *testing.T) {
	file := "config.toml"
	if _, err := os.Stat(file); err == nil {
		os.Remove(file)
	}

	plugins := make(map[string]bool)
	app.Loop(func(p app.Plugin) error {
		plugins[reflect.TypeOf(p).Elem().PkgPath()] = true
		return nil
	})

	if err := app.Write(file, &app.Config{
		Secrets: app.Secrets{
			Hmac:   app.Random(32),
			Cookie: app.Random(32),
			Csrf:   app.Random(32),
			Jwt:    app.Random(32),
			Aes:    app.Random(32),
		},
		PostgreSQL: app.PostgreSQL{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "",
			DbName:   "h2o_dev",
			SslMode:  "disable",
		},
		RabbitMQ: app.RabbitMQ{
			Host:     "localhost",
			Port:     5672,
			User:     "guest",
			Password: "guest",
			Virtual:  "h2o-dev",
		},
		Redis: app.Redis{
			Host: "localhost",
			Port: 6379,
			Db:   8,
		},
		Server: app.Server{
			Name:  "localhost",
			Port:  3000,
			Theme: "bootstrap",
			Ssl:   false,
		},
		SMTP: app.SMTP{
			Host:     "smtp.gmail.com",
			Port:     465,
			From:     "no-reply@change-me.com",
			Password: "change-me",
			Bcc:      []string{"web-master@change-me.com"},
		},
		Plugins: plugins,
	}); err != nil {
		t.Fatal(err)
	}

	if cfg, err := app.Read(file); err == nil {
		t.Logf("%+v", cfg)
	} else {
		t.Fatal(err)
	}

}
