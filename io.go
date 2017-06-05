package protorec

import (
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/uluyol/binrec"
)

// This file contains helper io functions compatible with
// Java protobuf's writeDelimitedTo and mergeDelimitedFrom methods.

// WriteDelimitedTo writes a proto.Message to an io.Writer in a varint-delimited format.
//
// WriteDelimitedTo is analogous to writeDelimitedTo in protobuf-java.
func WriteDelimitedTo(w io.Writer, m proto.Message) error {
	data, err := proto.Marshal(m)

	if err != nil {
		return err
	}
	return binrec.WriteDelimitedTo(w, data)
}

type Reader interface {
	io.ByteReader
	io.Reader
}

// ReadDelimitedFrom reads a proto.Message from a Reader in a varint-delimited format.
// bufio.Reader may be used to construct a Reader.
//
// ReadDelimitedFrom is analogous to mergeDelimitedFrom in protobuf-java.
func ReadDelimitedFrom(r Reader, m proto.Message) error {
	buf, err := binrec.ReadDelimitedFrom(r)
	if err != nil {
		return err
	}
	return proto.Unmarshal(buf, m)
}
