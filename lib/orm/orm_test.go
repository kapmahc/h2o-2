package orm_test

import (
	"testing"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/h2o/lib/app"
	"github.com/kapmahc/h2o/lib/orm"
)

var (
	tablename = app.Random(8)
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestMigrationGenerate(t *testing.T) {
	mig := orm.Migration{
		Name:      tablename,
		Timestamp: time.Now(),
		Up: []string{
			"CREATE TABLE t1" + tablename + "(id SERIAL, val VARCHAR);",
			"CREATE TABLE t2" + tablename + "(id SERIAL, val VARCHAR);",
		},
		Down: []string{
			"DROP TABLE t1" + tablename + ";",
			"DROP TABLE t2" + tablename + ";",
		},
	}
	if err := mig.Write(); err != nil {
		t.Fatal(err)
	}
}

func TestMigrationsRead(t *testing.T) {
	items, err := orm.ReadMigrations()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(items)
}

// -------------------------

func TestSeedGenerate(t *testing.T) {
	mig := orm.Seed{
		Name:      tablename,
		Timestamp: time.Now(),
		Lines: []string{
			"INSERT INTO t1" + tablename + "(val) values('aaa');",
			"INSERT INTO t2" + tablename + "(val) values('aaa');",
		},
	}
	if err := mig.Write(); err != nil {
		t.Fatal(err)
	}
}

func TestSeedsRead(t *testing.T) {
	items, err := orm.ReadMigrations()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(items)
}

// -------------------------

func TestMapperGenerate(t *testing.T) {
	mig := orm.Mapper{
		Name: tablename,
		Lines: map[string]string{
			"insert.t1": "INSERT INTO t1" + tablename + "(val) values('aaa');",
			"insert.t2": "INSERT INTO t2" + tablename + "(val) values('aaa');",
		},
	}
	if err := mig.Write(); err != nil {
		t.Fatal(err)
	}
}

func TestMappersRead(t *testing.T) {
	items, err := orm.ReadMappers()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(items)
}
