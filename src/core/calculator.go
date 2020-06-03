package core

import "bufio"

type Calculator interface {
	// return tuple (total line, comment line and error) in the file
	GetLinesNumber(r *bufio.Reader) (uint32, uint32, error)
}
