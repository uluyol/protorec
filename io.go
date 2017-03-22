package protorec

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"
)

// This file contains helper io functions compatible with
// Java protobuf's writeDelimitedTo and mergeDelimitedFrom methods.

// WriteDelimitedTo writes a proto.Message to an io.Writer in a varint-delimited format.
//
// WriteDelimitedTo is analogous to writeDelimitedTo in protobuf-java.
func WriteDelimitedTo(w io.Writer, m proto.Message) error {
	data, err := proto.Marshal(m)

	var buf [binary.MaxVarintLen64]byte

	n := binary.PutUvarint(buf[:], uint64(len(data)))

	concat := make([]byte, len(data)+n)
	copy(concat, buf[:n])
	copy(concat[n:], data)

	_, err = w.Write(concat)
	return err
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
	dlen, err := binary.ReadUvarint(r)
	if err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return err
		}
		return fmt.Errorf("unable to read length: %v", err)
	}
	buf := make([]byte, dlen)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return err
	}
	return proto.Unmarshal(buf, m)
}
