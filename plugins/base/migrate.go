package base

import (
	"database/sql"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
)

// CheckDb migrate database and register orm
func CheckDb() error {
	drv := beego.AppConfig.String("databasedrv")
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
	if err != nil {
		return err
	}
	if err := mig.Up(); err != nil {
		return err
	}

	orm.RegisterDataBase("default", drv, url)
	return nil
}
