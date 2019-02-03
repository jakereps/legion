package main

import (
	"bufio"
	"reflect"
	"testing"
)

func TestNewFastQ(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		want    *FastQ
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFastQ(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFastQ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFastQ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastQ_Read(t *testing.T) {
	type fields struct {
		scanner *bufio.Scanner
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Sequence
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FastQ{
				scanner: tt.fields.scanner,
			}
			got, err := f.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("FastQ.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FastQ.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
