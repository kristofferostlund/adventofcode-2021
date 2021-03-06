package fileutil

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/kristofferostlund/adventofcode-2021/pkg/numutil"
)

func FileFrom(relativePath string) (io.ReadCloser, error) {
	_, filename, _, _ := runtime.Caller(1)
	dir, _ := filepath.Abs(filepath.Dir(filename))
	absPath := filepath.Join(dir, relativePath)

	readCloser, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %w", absPath, err)
	}

	return readCloser, nil
}

func ScanLines(reader io.Reader, onLine func(index int, line string) error) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		if err := onLine(i, line); err != nil {
			return err
		}
		i++
	}

	return nil
}

func MapNonEmptyLines[V any](reader io.Reader, onLine func(line string) (V, error)) ([]V, error) {
	out := make([]V, 0)

	if err := ScanLines(reader, func(_ int, line string) error {
		if line == "" {
			return nil
		}

		v, err := onLine(line)
		if err != nil {
			return err
		}
		out = append(out, v)
		return nil
	}); err != nil {
		return nil, err
	}
	return out, nil
}

func ParseSingleLineOfInts(reader io.Reader) ([]int, error) {
	parsed, err := MapNonEmptyLines(reader, func(line string) ([]int, error) {
		return numutil.Atois(strings.Split(line, ","))
	})
	if err != nil {
		return nil, err
	}
	if len(parsed) != 1 {
		return nil, fmt.Errorf("unexpected row count %d", len(parsed))
	}
	return parsed[0], nil
}
