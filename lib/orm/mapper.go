package orm

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
)

// Mapper database scripts mapper
type Mapper struct {
	Name  string
	Lines map[string]string
}

func (p *Mapper) Write() error {
	root := mappersRoot()
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

	enc := toml.NewEncoder(fd)
	return enc.Encode(p.Lines)
}

func (p Mapper) String() string {
	return p.Name + ".toml"
}

// ReadMappers read database mappers
func ReadMappers() (map[string]string, error) {
	ext := ".toml"
	root := mappersRoot()
	items := make(map[string]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if !info.IsDir() && filepath.Ext(name) == ext {
			if len(name) <= len(ext)+1 {
				log.Errorf("bad mapper name %s", path)
				return nil
			}

			it := make(map[string]string)
			if _, err := toml.DecodeFile(path, &it); err != nil {
				return err
			}

			pre := name[:len(name)-len(ext)]
			for k, v := range it {
				items[pre+"."+k] = v
			}

			return nil
		}
		return nil
	})
	return items, err
}
