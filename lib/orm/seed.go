package orm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

// Seed database seed
type Seed struct {
	Name      string
	Timestamp time.Time
	Lines     []string
}

func (p *Seed) Write() error {
	root := seedsRoot()
	if err := os.MkdirAll(root, 0700); err != nil {
		return err
	}

	fn := filepath.Join(root, p.String())
	log.Infof("generate file %s", fn)
	fd, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	fd.WriteString(strings.Join(p.Lines, sep))
	return nil
}

func (p Seed) String() string {
	return fmt.Sprintf("%s_%s", p.Timestamp.Format(timestamp), p.Name) + ext
}

// ReadSeeds read database seeds
func ReadSeeds() ([]Seed, error) {
	root := seedsRoot()
	var items []Seed

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if !info.IsDir() && filepath.Ext(name) == ext {
			if len(name) <= len(timestamp)+len(ext)+1 {
				log.Errorf("bad seed name %s", path)
				return nil
			}

			var it Seed
			if it.Timestamp, err = time.Parse(timestamp, name[:len(timestamp)]); err != nil {
				return err
			}
			it.Name = name[len(timestamp)+1 : len(name)-len(ext)]
			if it.Lines, err = readScripts(path); err != nil {
				return err
			}

			items = append(items, it)
			return nil
		}
		return nil
	})
	return items, err
}
