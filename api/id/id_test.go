package id

import (
	"encoding/binary"
	"math"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func BenchmarkGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id := New()
		_ = id
	}
}

func BenchmarkMarshalText(b *testing.B) {
	id := New()
	for i := 0; i < b.N; i++ {
		bytes, _ := id.MarshalText()
		_ = bytes
	}
}

func BenchmarkUnmarshalText(b *testing.B) {
	id := New()
	bytes, _ := id.MarshalText()
	for i := 0; i < b.N; i++ {
		var id Id
		_ = id.UnmarshalText(bytes)
		_ = id
	}
}

func BenchmarkValidateText(b *testing.B) {
	id := New()
	bytes, _ := id.MarshalText()
	for i := 0; i < b.N; i++ {
		ValidateText(bytes)
	}
}

func TestValidInValid(t *testing.T) {
	id := New()
	bytes, _ := id.MarshalText()
	if !ValidateText(bytes) {
		t.Fatal("valid id should pass")
	}
	bytes[5] = ' '
	if ValidateText(bytes) {
		t.Fatal("invalid id should not pass")
	}
}

func TestIdRaw(t *testing.T) {
	SetMachineIdHost(net.IP{127, 0, 0, 1}, 8080)

	ts := time.Now()
	ms := uint64(ts.Unix())*1000 + uint64(ts.Nanosecond()/int(time.Millisecond))
	count := uint32(math.MaxUint32)
	id := newID(ms, machineID, count)

	var buf [8]byte
	copy(buf[2:], id[:6])
	idTime := binary.BigEndian.Uint64(buf[:])
	if ms != idTime {
		t.Fatal("id time doesn't not match time given", ms, idTime)
	}

	copy(buf[2:], id[6:12])
	idMid := binary.BigEndian.Uint64(buf[:])
	if idMid != machineID {
		t.Fatal("machine id mismatch", idMid, machineID)
	}

	idCount := binary.BigEndian.Uint32(id[12:16])
	if idCount != count {
		t.Fatal("count mismatch", idCount, count)
	}
}

func TestDescending(t *testing.T) {
	id := "0123WXYZ"

	flip := EncodeDescending(id)

	if len(flip) != len(id) {
		t.Fatal("flipped string has different length:", len(flip), len(id))
	}

	for i := range flip {
		if flip[i] != id[len(id)-1-i] {
			t.Fatalf("flipped encoding not working. got: %v, want: %v", flip[i], id[len(id)-1-i])
		}
	}
}

func Test_reverseString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "reverse english string", args: args{input: "ABCDEFG1234"}, want: "4321GFEDCBA"},
		{name: "reverse korean string", args: args{input: "가나다라마바사1234"}, want: "4321사바마라다나가"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseString(tt.args.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
