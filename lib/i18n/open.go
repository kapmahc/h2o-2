package i18n

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/kapmahc/h2o/lib/orm"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
)

var (
	_db      *sql.DB
	_locales map[string]map[string]string
)

// Set set locale
func Set(lang, code, message string) error {
	var id uint
	err := orm.One(
		_db,
		func(it *sql.Row) error {
			return it.Scan(&id)
		},
		"i18n.locales.select-id",
		code, message)

	if err == nil {
		_, _, err = orm.Do(_db, "i18n.locales.update", id)
	} else if err == sql.ErrNoRows {
		_, _, err = orm.Do(_db, "i18n.locales.insert", lang, code, message)
	}

	if err == nil {
		set(lang, code, message)
	}
	return err
}

// Open load locale records
func Open(db *sql.DB) error {
	// clear
	_locales = make(map[string]map[string]string)

	// load from filesystem
	if err := filepath.Walk("locales", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		names := strings.Split(info.Name(), ".")
		if info.IsDir() || len(names) != 3 || names[2] != ".yaml" {
			return nil
		}
		lang, err := language.Parse(names[1])
		if err != nil {
			return err
		}

		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		val := make(map[interface{}]interface{})
		if err = yaml.Unmarshal(buf, &val); err != nil {
			return err
		}

		return loopNode(names[0], val, func(code string, message string) error {
			// log.Debugf("%s.%s = %s", lang.String(), code, message)
			set(lang.String(), code, message)
			return nil
		})

	}); err != nil {
		return err
	}

	// check database
	for _, qry := range []string{
		"table",
		"index-code-lang",
		"index-lang",
		"index-code",
	} {
		if _, _, err := orm.Do(db, "i18n.locales.create-"+qry); err != nil {
			return err
		}
	}
	// load from database
	if err := orm.All(
		db,
		func(it *sql.Rows) error {
			var lang, code, message string
			if err := it.Scan(&lang, &code, &message); err != nil {
				return err
			}
			set(lang, code, message)
			return nil
		},
		"locales.select-all",
	); err != nil {
		return err
	}

	// done
	_db = db
	return nil
}

func loopNode(r string, m map[interface{}]interface{}, f func(string, string) error) error {
	for k, v := range m {
		ks, ok := k.(string)
		if ok {
			ks = r + "." + ks
			vs, ok := v.(string)
			if ok {
				if e := f(ks, vs); e != nil {
					return e
				}
			} else {
				vm, ok := v.(map[interface{}]interface{})
				if ok {
					if e := loopNode(ks, vm, f); e != nil {
						return e
					}
				}
			}
		}
	}
	return nil
}

func set(lang, code, message string) {
	if _, ok := _locales[lang]; !ok {
		_locales[lang] = make(map[string]string)
	}
	_locales[lang][code] = message
}

func get(lang, code string) (string, bool) {
	items, ok := _locales[lang]
	if ok {
		if msg, ok := items[code]; ok {
			return msg, true
		}
	}
	return lang + "." + code, false
}
