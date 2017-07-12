package web

import (
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"golang.org/x/text/language"
)

var (
	// Version version
	Version string
	// BuildTime build time
	BuildTime string
	// Usage usage
	Usage string
	// Copyright copyright
	Copyright string
	// AuthorName author's name
	AuthorName string
	// AuthorEmail author's email
	AuthorEmail string
)

// Main main entry
func Main() error {

	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Version = fmt.Sprintf("%s (%s)", Version, BuildTime)
	app.Authors = []cli.Author{
		cli.Author{
			Name:  AuthorName,
			Email: AuthorEmail,
		},
	}
	app.Copyright = Copyright
	app.Usage = Usage
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{}

	for _, en := range plugins {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	return app.Run(os.Args)
}

func init() {
	viper.SetEnvPrefix("h2o")
	viper.BindEnv("env")
	viper.SetDefault("env", "development")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("redis", map[string]interface{}{
		"host": "localhost",
		"port": 6379,
		"db":   8,
	})

	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"host":     "localhost",
			"port":     5432,
			"user":     "postgres",
			"password": "",
			"dbname":   "h2o_dev",
			"sslmode":  "disable",
		},
		"pool": map[string]int{
			"max_open": 180,
			"max_idle": 6,
		},
	})

	viper.SetDefault("server", map[string]interface{}{
		"name":  "www.change-me.com",
		"port":  8080,
		"ssl":   true,
		"theme": "bootstrap",
	})

	viper.SetDefault("secrets", map[string]interface{}{
		"jwt":    RandomStr(32),
		"aes":    RandomStr(32),
		"hmac":   RandomStr(32),
		"csrf":   RandomStr(32),
		"cookie": RandomStr(32),
	})

	viper.SetDefault("elasticsearch", map[string]interface{}{
		"host": "localhost",
		"port": 9200,
	})

	viper.SetDefault("languages", []string{
		language.AmericanEnglish.String(),
		language.SimplifiedChinese.String(),
		language.TraditionalChinese.String(),
	})

	viper.SetDefault("uploader", map[string]string{
		"path":     path.Join("public", "attachments"),
		"endpoint": "/public/attachment",
	})
}
