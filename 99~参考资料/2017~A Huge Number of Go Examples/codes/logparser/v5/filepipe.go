
package main

import (
	"os"
	"strings"

	"github.com/inancgumus/learngo/logparser/v5/pipe"
)

// fromFile generates a default pipeline.
// Detects the correct parser by the file extension.
// Uses a TextReport and groups by domain.
func fromFile(path string) (*pipe.Pipeline, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var src pipe.Iterator
	switch {
	case strings.HasSuffix(path, ".txt"):
		src = pipe.NewTextLog(f)
	case strings.HasSuffix(path, ".jsonl"):
		src = pipe.NewJSONLog(f)
	}

	return pipe.New(
		src,
		pipe.NewTextReport(os.Stdout),
		pipe.GroupBy(pipe.DomainGrouper),
	), nil
}
