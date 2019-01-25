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
	s.Split(bufio.ScanRunes)

	return &FastQ{
		scanner: s,
	}, nil
}

// Base is a type to hold the enumerated sequence bases
type Base byte

// Base pairs enumeration
const (
	A Base = iota
	C
	G
	T
)

// Nucleobase holds the pairing of a base and its quality score
type Nucleobase struct {
	Base    Base
	Quality uint8
}

// Sequence represents a full sequence read
type Sequence struct {
	ID   string
	Data []Nucleobase
}

func (f *FastQ) Read() (*Sequence, error) {
	s := &Sequence{}
	for f.scanner.Scan() {
		t := f.scanner.Text()
		if t == "\n" {
			break
		}
		s.ID += t
	}

	// n := Nucleobase{}

	return s, nil
}

func main() {
	sfpath := "/Users/jordenkreps/Downloads/sequences.fastq.gz"
	// bcpath := "/Users/jordenkreps/Downloads/barcodes.fastq.gz"
	// TODO: Read in sequence files (F/R S/P)

	seqs, err := NewFastQ(sfpath)
	if err != nil {
		log.Fatal(err)
		return
	}
	_ = seqs

	for {

		s, err := seqs.Read()
		if err != nil {
			log.Fatal(err)
			return
		}
		_ = s
		break
	}

	// TODO: Read in barcodes

	// bc, err := NewFastQ(bcpath)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// _ = bc
	// // TODO: Demux/Error correct?

	// // TODO: Write results

}
