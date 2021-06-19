package legion

// Demux ...
type Demux struct {
	Paired bool
}

// Demuxer ...
type Demuxer interface {
	Demux() (*Demux, error)
}
