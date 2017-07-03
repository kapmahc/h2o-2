package orm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

// Migration database migration
type Migration struct {
	Name      string
	Timestamp time.Time
	Up        []string
	Down      []string
}

func (p *Migration) Write() error {
	if err := p.touch("up"+ext, p.Up); err != nil {
		return err
	}
	if err := p.touch("down"+ext, p.Down); err != nil {
		return err
	}
	return nil
}

func (p *Migration) touch(name string, lines []string) error {
	root := filepath.Join(
		migrationRoot(),
		p.String(),
	)
	if err := os.MkdirAll(root, 0700); err != nil {
		return err
	}

	fn := filepath.Join(root, name)
	log.Infof("generate file %s", fn)
	fd, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	fd.WriteString(strings.Join(lines, sep))
	return nil
}

func (p Migration) String() string {
	return fmt.Sprintf("%s_%s", p.Timestamp.Format(timestamp), p.Name)
}

// ReadMigrations read migrations
func ReadMigrations() ([]Migration, error) {
	root := migrationRoot()
	var items []Migration

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if info.IsDir() && path != root {
			if len(name) <= len(timestamp)+1 {
				log.Errorf("bad migration name %s", path)
				return nil
			}

			var it Migration
			if it.Timestamp, err = time.Parse(timestamp, name[:len(timestamp)]); err != nil {
				return err
			}
			it.Name = name[len(timestamp)+1:]
			if it.Up, err = readScripts(filepath.Join(path, "up"+ext)); err != nil {
				return err
			}

			if it.Down, err = readScripts(filepath.Join(path, "down"+ext)); err != nil {
				return err
			}
			items = append(items, it)
			return nil
		}
		return nil
	})
	return items, err
}
