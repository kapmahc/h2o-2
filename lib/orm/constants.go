package orm

import (
	"bufio"
	"os"
	"path"
	"strings"
)

const (
	ext       = ".sql"
	sep       = "\n"
	timestamp = "20060102150405"
)

func seedsRoot() string   { return path.Join("db", "seeds") }
func mappersRoot() string { return path.Join("db", "scripts") }

func migrationRoot() string { return path.Join("db", "migrations") }

func readScripts(file string) ([]string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	var lines []string
	for scanner.Scan() {
		it := strings.TrimSpace(scanner.Text())
		if it != "" {
			lines = append(lines, it)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
