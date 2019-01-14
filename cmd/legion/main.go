package main

import (
	"bufio"
	"compress/gzip"
	"log"
	"os"
)

// FastQ is
type FastQ struct {
	scanner *bufio.Scanner
}

// NewFastQ makes
func NewFastQ(p string) (*FastQ, error) {
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
	return &FastQ{
		scanner: s,
	}, nil
}

func main() {
	sfpath := "/Users/jordenkreps/Downloads/sequences.fastq.gz"
	bcpath := "/Users/jordenkreps/Downloads/barcodes.fastq.gz"
	// TODO: Read in sequence files (F/R S/P)

	seqs, err := NewFastQ(sfpath)
	if err != nil {
		log.Fatal(err)
		return
	}
	_ = seqs

	// TODO: Read in barcodes

	bc, err := NewFastQ(bcpath)
	if err != nil {
		log.Fatal(err)
		return
	}
	_ = bc
	// TODO: Demux/Error correct?

	// TODO: Write results

}
