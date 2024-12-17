package binary

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestReadBool(t *testing.T) {
	data := []byte{1, 0}
	reader := NewReader(bytes.NewReader(data), binary.NativeEndian)

	b, err := reader.ReadBool()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !b {
		t.Fatalf("expected true, got false")
	}

	b, err = reader.ReadBool()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if b {
		t.Fatalf("expected false, got true")
	}
}

func TestReadInt8(t *testing.T) {
	data := []byte{0x80}
	reader := NewReader(bytes.NewReader(data), binary.NativeEndian)

	v, err := reader.ReadInt8()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != -128 {
		t.Fatalf("expected -128, got %d", v)
	}
}

func TestReadUint8(t *testing.T) {
	data := []byte{0xFF}
	reader := NewReader(bytes.NewReader(data), binary.NativeEndian)

	v, err := reader.ReadUint8()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 255 {
		t.Fatalf("expected 255, got %d", v)
	}
}

func TestReadInt16(t *testing.T) {
	data := []byte{0x00, 0x80}
	reader := NewReader(bytes.NewReader(data), binary.BigEndian)

	v, err := reader.ReadInt16()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 128 {
		t.Fatalf("expected 128, got %d", v)
	}
}

func TestReadUint16(t *testing.T) {
	data := []byte{0x00, 0xFF}
	reader := NewReader(bytes.NewReader(data), binary.BigEndian)

	v, err := reader.ReadUint16()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 255 {
		t.Fatalf("expected 255, got %d", v)
	}
}

func TestReadInt32(t *testing.T) {
	data := []byte{0x00, 0x00, 0x00, 0x80}
	reader := NewReader(bytes.NewReader(data), binary.BigEndian)

	v, err := reader.ReadInt32()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 128 {
		t.Fatalf("expected 128, got %d", v)
	}
}

func TestReadNegInt32(t *testing.T) {
	data := []byte{0xff, 0xff, 0xff, 0xfe}
	reader := NewReader(bytes.NewReader(data), binary.BigEndian)

	v, err := reader.ReadInt32()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != -2 {
		t.Fatalf("expected 128, got %d", v)
	}
}

func TestReadUint32(t *testing.T) {
	data := []byte{0x00, 0x00, 0x00, 0xFF}
	reader := NewReader(bytes.NewReader(data), binary.BigEndian)

	v, err := reader.ReadUint32()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 255 {
		t.Fatalf("expected 255, got %d", v)
	}
}

func TestReadInt64(t *testing.T) {
	data := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80}
	reader := NewReader(bytes.NewReader(data), binary.BigEndian)

	v, err := reader.ReadInt64()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 128 {
		t.Fatalf("expected 128, got %d", v)
	}
}

func TestReadUint64(t *testing.T) {
	data := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF}
	reader := NewReader(bytes.NewReader(data), binary.BigEndian)

	v, err := reader.ReadUint64()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 255 {
		t.Fatalf("expected 255, got %d", v)
	}
}

func TestReadFloat32(t *testing.T) {
	data := []byte{0x3F, 0x80, 0x00, 0x00}
	reader := NewReader(bytes.NewReader(data), binary.BigEndian)

	v, err := reader.ReadFloat32()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 1.0 {
		t.Fatalf("expected 1.0, got %f", v)
	}
}

func TestReadFloat64(t *testing.T) {
	data := []byte{0x3F, 0xF0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	reader := NewReader(bytes.NewReader(data), binary.BigEndian)

	v, err := reader.ReadFloat64()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 1.0 {
		t.Fatalf("expected 1.0, got %f", v)
	}
}

func TestReadNullTerminatedString(t *testing.T) {
	data := []byte("Hello, World!\x00More data...")
	reader := NewReader(bytes.NewReader(data), binary.NativeEndian)

	str, err := reader.ReadNullTerminatedString()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if str != "Hello, World!" {
		t.Fatalf("expected 'Hello, World!', got '%s'", str)
	}
}
