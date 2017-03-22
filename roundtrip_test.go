package protorec

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

//go:generate bash -c "protoc --go_out=. simple_test.proto && mv simple_test.pb.go pb_simple_test.go"

func TestRoundTrip(t *testing.T) {
	msgs := []TestMsg{
		{
			F1: "abasdfasdf312123",
			F2: 109,
			F3: ^uint64(0),
			F4: []byte{0, 12, 44},
			F5: 505.5,
		}, {
			F1: "",
			F4: nil,
		}, {
			F3: 99,
		}, {},
		{
			F2: 0,
		},
	}

	var buf bytes.Buffer
	for i := range msgs {
		if err := WriteDelimitedTo(&buf, &msgs[i]); err != nil {
			t.Fatalf("unable to write %d: %v", i, err)
		}
	}

	readMsgs := make([]TestMsg, len(msgs))
	for i := range msgs {
		if err := ReadDelimitedFrom(&buf, &readMsgs[i]); err != nil {
			t.Fatalf("unable to read %d: %v", i, err)
		}
	}

	if err := ReadDelimitedFrom(&buf, &TestMsg{}); err != io.EOF {
		t.Errorf("want EOF after reading all data, got %v", err)
	}

	for i := range msgs {
		if !reflect.DeepEqual(&msgs[i], &readMsgs[i]) {
			t.Errorf("case %d: written and read messages differ", i)
		}
	}
}
