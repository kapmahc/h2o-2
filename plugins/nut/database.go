package nut

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
)

type migrateLogger struct {
}

func (p *migrateLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (p *migrateLogger) Verbose() bool {
	return true
}

type databaseCheck struct {
	driver string
	source string
}

func (p *databaseCheck) Check() error {

	db, err := sql.Open(p.driver, p.source)
	if err != nil {
		return err
	}
	return db.Ping()
}

// CheckDb migrate database and register orm
func CheckDb() error {
	drv := beego.AppConfig.String("databasedriver")
	url := beego.AppConfig.String("databaseurl")

	db, err := sql.Open(drv, url)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	ins, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	mig, err := migrate.NewWithDatabaseInstance("file://"+filepath.Join("db", drv, "migrations"), drv, ins)
	mig.Log = &migrateLogger{}
	if err != nil {
		return err
	}
	if err := mig.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	toolbox.AddHealthCheck("database", &databaseCheck{driver: drv, source: url})
	orm.Debug = beego.BConfig.RunMode != beego.PROD
	orm.RegisterDataBase("default", drv, url)
	return nil
}
