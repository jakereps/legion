package fastq

import (
	"bufio"
	"compress/gzip"
	"os"
)

// newScanner takes a filepath to a gzip'd file and returns a scanner
// set to split by character.
func newScanner(p string) (*bufio.Scanner, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	zr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(zr)
	if err != nil {
		return nil, err
	}
	s.Split(bufio.ScanRunes)
	return s, nil
}
