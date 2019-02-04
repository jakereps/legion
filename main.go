package main

import (
	"flag"
	"fmt"
	"os"
)

// Demux ...
type Demux struct {
	Paired bool
}

// Demuxer ...
type Demuxer interface {
	Demux() (*Demux, error)
}

// SingleEndFASTQ ...
type SingleEndFASTQ struct {
	Sequences *FASTQ
	Index     *Index
}

// PairedEndFASTQ ...
type PairedEndFASTQ struct {
	Forward *FASTQ
	Reverse *FASTQ
	Index   *Index
}

// Demux ...
func (s *SingleEndFASTQ) Demux() (*Demux, error) {
	return &Demux{
		Paired: false,
	}, nil
}

// Demux ...
func (p *PairedEndFASTQ) Demux() (*Demux, error) {
	return &Demux{
		Paired: true,
	}, nil
}

// SingleEnd ...
func SingleEnd(fwd, idx string) (*SingleEndFASTQ, error) {
	f, err := NewFASTQ(fwd)
	if err != nil {
		return nil, err
	}
	return &SingleEndFASTQ{Sequences: f}, nil
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
	return &PairedEndFASTQ{
		Forward: f,
		Reverse: r,
	}, nil
}

func main() {
	var fwd string
	var rev string
	var idx string
	var md string
	flag.StringVar(&fwd, "f", "", "path to forward fastq file")
	flag.StringVar(&rev, "r", "", "path to reverse fastq file")
	flag.StringVar(&idx, "i", "", "path to index fastq file")
	flag.StringVar(&md, "m", "", "path to metadata file")
	flag.Parse()

	switch {
	case fwd == "" && rev == "":
		fmt.Fprintln(os.Stderr, "need at least one read file, none provided.")
		os.Exit(1)
	case fwd != "" && rev == "":
		s, err := SingleEnd(fwd, idx)
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to create single end fastq: "+err.Error())
			os.Exit(1)
		}
		fmt.Println(s)
	case fwd != "" && rev != "":
		p, err := PairedEnd(fwd, rev, idx)
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to create paired end fastq: "+err.Error())
			os.Exit(1)
		}
		fmt.Println(p)
	default:
		fmt.Fprintln(os.Stderr, "not sure how you got here, congrats!")
		os.Exit(9001)
	}
}
