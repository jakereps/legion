package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var fwd string
	var rev string
	var md string
	flag.StringVar(&fwd, "f", "", "path to forward fastq file")
	flag.StringVar(&rev, "r", "", "path to reverse fastq file")
	flag.StringVar(&md, "m", "", "path to metadata file")
	flag.Parse()

	if fwd == "" && rev == "" {
		fmt.Println("need at least one read file, none provided.")
		os.Exit(1)
	}

	var f *FastQ
	if fwd != "" {
		fq, err := NewFastQ(fwd)
		if err != nil {
			fmt.Println("failed to create fastq: " + err.Error())
			os.Exit(1)
		}
		f = fq
	}

	fmt.Printf("%+v", f)

	for {
		s, err := f.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("failed to read fastq: " + err.Error())
			os.Exit(1)
		}

		_ = s
	}

	// TODO: Read in barcodes

	// TODO: Demux/Error correct?

	// TODO: Write results

}
