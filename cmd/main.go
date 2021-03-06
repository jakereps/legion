package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jakereps/legion"
	"github.com/jakereps/legion/fastq"
)

func main() {
	fwd := flag.String("f", "", "path to forward fastq file")
	rev := flag.String("r", "", "path to reverse fastq file")
	idx := flag.String("i", "", "path to index fastq file")
	md := flag.String("m", "", "path to metadata file")
	flag.Parse()

	_ = md

	var (
		r   legion.Reader
		err error
	)
	switch {
	case *fwd != "" && *rev == "":
		r, err = fastq.SingleEnd(*fwd, *idx)
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to create single end fastq: "+err.Error())
			os.Exit(1)
		}
		fmt.Println(r)
	case *fwd != "" && *rev != "":
		r, err = fastq.PairedEnd(*fwd, *rev, *idx)
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to create paired end fastq: "+err.Error())
			os.Exit(1)
		}
		fmt.Println(r)
	case *fwd == "" && *rev == "":
		fmt.Fprintln(os.Stderr, "need at least one read file, none provided.")
		os.Exit(1)
	default:
		fmt.Fprintln(os.Stderr, "not sure how you got here, congrats!")
		os.Exit(9001)
	}

	d, err := legion.NewDemultiplexer(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed creating demux")
		os.Exit(1)
	}
	d.Run()
}
