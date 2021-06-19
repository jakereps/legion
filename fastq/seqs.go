package fastq

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// File is the sequence data.
type File struct {
	scanner *bufio.Scanner
}

func (f *File) Read() (*Sequence, error) {

	// Line 1 begins with a '@' character and is followed by a sequence
	// identifier and an optional description (like a FASTA title line).
	var id string
	for f.scanner.Scan() {
		t := f.scanner.Text()
		if t == "\n" {
			break
		}
		id += t
	}

	// Line 2 is the raw sequence letters.
	var seqs []Nucleobase
	for f.scanner.Scan() {
		t := f.scanner.Text()
		if t == "\n" {
			break
		}
		if v, ok := baseFromChar[t]; ok {
			seqs = append(seqs, Nucleobase{Base: v})
		}
	}

	// Line 3 begins with a '+' character and is optionally followed by the same
	// sequence identifier (and any description) again.
	var div string
	for f.scanner.Scan() {
		t := f.scanner.Text()
		if t == "\n" {
			break
		}
		div += t
	}

	// Line 4 encodes the quality values for the sequence in Line 2, and must
	// contain the same number of symbols as letters in the sequence.
	for i := range seqs {
		f.scanner.Scan()
		t := f.scanner.Text()
		// taking the 0 index on a string gets the byte value
		// NOTE: make the phred score configurable probably
		seqs[i].Quality = t[0] - 33
	}

	// kill the newline handle EOF
	if !f.scanner.Scan() {
		return nil, io.EOF
	}

	return &Sequence{
		ID:      id,
		Divider: div,
		Data:    seqs,
	}, nil
}

// SingleEndFASTQ ...
type SingleEndFASTQ struct {
	Sequences *File
	Index     *File
}

// PairedEndFASTQ ...
type PairedEndFASTQ struct {
	Forward *File
	Reverse *File
	Index   *File
}

// Demux ...
func (s *SingleEndFASTQ) Demux() error {
	var (
		seq *Sequence
		bar *Sequence

		i int

		err error
	)
	for {
		i++
		seq, err = s.Sequences.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("failed reading seq #%d: %w", i, err)
		}
		bar, err = s.Index.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("failed reading index #%d: %w", i, err)
		}
		fmt.Println(seq.String(), bar.String())
	}
	return nil
}

// Demux ...
func (p *PairedEndFASTQ) Demux() error {
	return nil
}

// SingleEnd ...
func SingleEnd(fwd, idx string) (*SingleEndFASTQ, error) {
	f, err := NewFASTQ(fwd)
	if err != nil {
		return nil, err
	}
	i, err := NewFASTQ(idx)
	if err != nil {
		return nil, err
	}
	return &SingleEndFASTQ{
		Sequences: f,
		Index:     i,
	}, nil
}

// PairedEnd ...
func PairedEnd(fwd, rev, idx string) (*PairedEndFASTQ, error) {
	f, err := NewFASTQ(fwd)
	if err != nil {
		return nil, err
	}
	r, err := NewFASTQ(rev)
	if err != nil {
		return nil, err
	}
	i, err := NewFASTQ(idx)
	if err != nil {
		return nil, err
	}
	return &PairedEndFASTQ{
		Forward: f,
		Reverse: r,
		Index:   i,
	}, nil
}

// NewFASTQ makes
func NewFASTQ(p string) (*File, error) {
	s, err := newScanner(p)
	if err != nil {
		return nil, err
	}

	return &File{
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

var baseToChar = map[Base]string{
	N: "n",
	A: "a",
	C: "c",
	G: "g",
	T: "t",
}

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

func (s *Sequence) String() string {
	result := make([]string, len(s.Data))
	for i := range s.Data {
		result[i] = baseToChar[s.Data[i].Base]
	}
	return strings.Join(result, "")
}

func (s *Sequence) Quality() []uint8 {
	result := make([]uint8, len(s.Data))
	for i := range s.Data {
		result[i] = s.Data[i].Quality
	}
	return result
}
