package fileutil

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
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

func MapLines[V any](reader io.Reader, onLine func(line string) (V, error)) ([]V, error) {
	out := make([]V, 0)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		v, err := onLine(line)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}

	return out, nil
}
