package legion

import (
	"errors"
	"fmt"
	"io"

	"github.com/jakereps/legion/fastq"
)

// Demuxer ...
type Reader interface {
	Read() (*fastq.MultiplexedSequence, error)
}

// Demux ...
type Demux struct {
	r Reader
}

func (d *Demux) Run() {
	s, err := d.r.Read()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return
		}
		panic(err)
	}
	fmt.Println(s)
}

func NewDemultiplexer(r Reader) (*Demux, error) {
	return &Demux{
		r: r,
	}, nil
}
