package fastq

// FastQ represents a FastQ stream
type FastQ struct{}

// Read will read len(p) bytes into p
func (f *FastQ) Read(p []byte) (int, error) {
	return 0, nil
}

// Write will write len(p) bytes into p
func (f *FastQ) Write(p []byte) (int, error) {
	return 0, nil
}
