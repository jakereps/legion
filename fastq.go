package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
)

// FASTQ is the sequence data.
type FASTQ struct {
	scanner *bufio.Scanner
}

// Index is the index file of the reads. Sometimes called Barcodes.
type Index struct {
	scanner *bufio.Scanner
}

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

// NewIndex initializes an Index for a given filepath.
func NewIndex(p string) (*Index, error) {
	s, err := newScanner(p)
	if err != nil {
		return nil, err
	}

	return &Index{
		scanner: s,
	}, nil
}

// NewFASTQ makes
func NewFASTQ(p string) (*FASTQ, error) {
	s, err := newScanner(p)
	if err != nil {
		return nil, err
	}

	return &FASTQ{
		scanner: s,
	}, nil
}

// Base is a type to hold the enumerated sequence bases
type Base byte

// Base pairs enumeration
const (
	N Base = iota
	A
	C
	G
	T
)

var baseFromChar = map[string]Base{
	"n": N,
	"N": N,
	"a": A,
	"A": A,
	"c": C,
	"C": C,
	"g": G,
	"G": G,
	"t": T,
	"T": T,
}

// Nucleobase holds the pairing of a base and its quality score
type Nucleobase struct {
	Base    Base
	Quality uint8
}

// Sequence represents a full sequence read
type Sequence struct {
	ID      string
	Divider string
	Data    []Nucleobase
}

func (f *FASTQ) Read() (*Sequence, error) {
	s := &Sequence{}

	// Line 1 begins with a '@' character and is followed by a sequence
	// identifier and an optional description (like a FASTA title line).
	for f.scanner.Scan() {
		t := f.scanner.Text()
		if t == "\n" {
			break
		}
		s.ID += t
	}

	// Line 2 is the raw sequence letters.
	for f.scanner.Scan() {
		t := f.scanner.Text()
		if t == "\n" {
			break
		}
		if v, ok := baseFromChar[t]; ok {
			s.Data = append(s.Data, Nucleobase{Base: v})
		}
	}

	// Line 3 begins with a '+' character and is optionally followed by the same
	// sequence identifier (and any description) again.
	for f.scanner.Scan() {
		t := f.scanner.Text()
		if t == "\n" {
			break
		}
		s.Divider += t
	}

	// Line 4 encodes the quality values for the sequence in Line 2, and must
	// contain the same number of symbols as letters in the sequence.
	for i := range s.Data {
		f.scanner.Scan()
		t := f.scanner.Text()
		// taking the 0 index on a string gets the byte value
		// NOTE: make the phred score configurable probably
		s.Data[i].Quality = t[0] - 33
	}

	// kill the newline handle EOF
	if !f.scanner.Scan() {
		return nil, io.EOF
	}

	return s, nil
}
