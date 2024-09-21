
package main

import (
	"bufio"
	"fmt"
	"io"
)

func textReader(r io.Reader) inputFn {
	return func(process processFn) error {
		var (
			l  = 1
			in = bufio.NewScanner(r)
		)

		for in.Scan() {
			r, err := fastParseFields(in.Bytes())
			// r, err := parseFields(in.Text())
			if err != nil {
				return fmt.Errorf("line %d: %v", l, err)
			}

			process(r)
			l++
		}

		if c, ok := r.(io.Closer); ok {
			c.Close()
		}
		return in.Err()
	}
}
