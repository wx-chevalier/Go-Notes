
package pipe

import (
	"io"
)

// readClose the reader if it's a io.Closer.
func readClose(r io.Reader) {
	if rc, ok := r.(io.Closer); ok {
		rc.Close()
	}
}
