package site

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/kapmahc/h2o/web"
	log "github.com/sirupsen/logrus"
)

func (p *Plugin) openRender(theme string) (*template.Template, error) {

	funcs := template.FuncMap{
		"t": p.I18n.T,
		"tn": func(v interface{}) string {
			return reflect.TypeOf(v).String()
		},
		"even": func(i interface{}) bool {
			if i != nil {
				switch i.(type) {
				case int:
					return i.(int)%2 == 0
				case uint:
					return i.(uint)%2 == 0
				case int64:
					return i.(int64)%2 == 0
				case uint64:
					return i.(uint64)%2 == 0
				}
			}
			return false
		},
		"fmt": fmt.Sprintf,
		"eq": func(arg1, arg2 interface{}) bool {
			return arg1 == arg2
		},
		"str2htm": func(s string) template.HTML {
			return template.HTML(s)
		},
		"dtf": func(t interface{}) string {
			if t != nil {
				f := "Mon Jan _2 15:04:05 2006"
				switch t.(type) {
				case time.Time:
					return t.(time.Time).Format(f)
				case *time.Time:
					if t != (*time.Time)(nil) {
						return t.(*time.Time).Format(f)
					}
				}
			}
			return ""
		},
		"df": func(t interface{}) string {
			if t != nil {
				f := "Mon Jan _2 2006"
				switch t.(type) {
				case time.Time:
					return t.(time.Time).Format(f)
				case *time.Time:
					if t != (*time.Time)(nil) {
						return t.(*time.Time).Format(f)
					}
				}
			}
			return ""
		},
		"links": func(loc string) []web.Link {
			var items []web.Link
			if err := p.Db.Where("loc = ?", loc).Order("sort_order DESC").Find(&items).Error; err != nil {
				log.Error(err)
			}
			return items
		},
		"cards": func(loc string) []web.Card {
			var items []web.Card
			if err := p.Db.Where("loc = ?", loc).Order("sort_order DESC").Find(&items).Error; err != nil {
				log.Error(err)
			}
			return items
		},
		"in": func(o interface{}, args []interface{}) bool {
			for _, v := range args {
				if o == v {
					return true
				}
			}
			return false
		},
		"starts": func(s string, b string) bool {
			return strings.HasPrefix(s, b)
		},
	}

	var files []string
	filepath.Walk(path.Join("themes", theme, "views"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".html" {
			files = append(files, path)
		}
		return nil
	})

	return template.New("").Funcs(funcs).ParseFiles(files...)
}
