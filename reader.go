package binary

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"
)

func NewReader(r io.Reader, byteOrder binary.ByteOrder) *Reader {
	return &Reader{
		r:         r,
		byteOrder: byteOrder,
	}
}

func NewReaderFromBytes(data []byte, byteOrder binary.ByteOrder) *Reader {
	return &Reader{
		r:         bytes.NewBuffer(data),
		byteOrder: byteOrder,
	}
}

type Reader struct {
	r         io.Reader
	byteOrder binary.ByteOrder
}

func (r *Reader) ByteOrder() binary.ByteOrder {
	return r.byteOrder
}

func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	s, ok := r.r.(io.Seeker)
	if !ok {
		return 0, errors.New("reader does not support seek")
	}

	return s.Seek(offset, whence)
}

func (r *Reader) Peek(n int) ([]byte, error) {
	b, ok := r.r.(*bufio.Reader)
	if !ok {
		return nil, errors.New("reader does not support peek")
	}

	return b.Peek(n)
}

func (r *Reader) Read(p []byte) (int, error) {
	return r.r.Read(p)
}

func (r *Reader) ReadAll() ([]byte, error) {
	return io.ReadAll(r.r)
}

func (r *Reader) ReadByte() (byte, error) {
	data, err := r.ReadBytes(1)
	if err != nil {
		return 0x00, err
	}

	return data[0], nil
}

func (r *Reader) ReadBytes(count int64) ([]byte, error) {
	buf := make([]byte, count)

	_, err := io.ReadFull(r.r, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *Reader) SkipBytes(count int64) error {
	_, err := io.CopyN(io.Discard, r.r, count)
	if err != nil {
		return err
	}

	return nil
}

func (r *Reader) ReadBool() (bool, error) {
	b, err := r.ReadBytes(1)
	if err != nil {
		return false, err
	}

	return b[0] != 0, nil
}

func (r *Reader) ReadInt8() (int8, error) {
	u8, err := r.ReadUint8()
	if err != nil {
		return 0, err
	}

	return int8(u8), nil
}

func (r *Reader) ReadUint8() (uint8, error) {
	b, err := r.ReadBytes(1)
	if err != nil {
		return 0, err
	}

	return b[0], nil
}

func (r *Reader) ReadInt16() (int16, error) {
	u16, err := r.ReadUint16()
	if err != nil {
		return 0, err
	}

	return int16(u16), nil
}

func (r *Reader) ReadUint16() (uint16, error) {
	b, err := r.ReadBytes(2)
	if err != nil {
		return 0, err
	}

	return r.byteOrder.Uint16(b), nil
}

func (r *Reader) ReadInt32() (int32, error) {
	u32, err := r.ReadUint32()
	if err != nil {
		return 0, err
	}

	return int32(u32), nil
}

func (r *Reader) ReadUint32() (uint32, error) {
	b, err := r.ReadBytes(4)
	if err != nil {
		return 0, err
	}

	return r.byteOrder.Uint32(b), nil
}

func (r *Reader) ReadInt64() (int64, error) {
	u64, err := r.ReadUint64()
	if err != nil {
		return 0, err
	}

	return int64(u64), nil
}

func (r *Reader) ReadUint64() (uint64, error) {
	b, err := r.ReadBytes(8)
	if err != nil {
		return 0, err
	}

	return r.byteOrder.Uint64(b), nil
}

func (r *Reader) ReadFloat32() (float32, error) {
	u32, err := r.ReadUint32()
	if err != nil {
		return 0, err
	}

	return math.Float32frombits(u32), nil
}

func (r *Reader) ReadFloat64() (float64, error) {
	u64, err := r.ReadUint64()
	if err != nil {
		return 0, err
	}

	return math.Float64frombits(u64), nil
}

func (r *Reader) ReadNullTerminatedString() (string, error) {
	var buf bytes.Buffer
	b := make([]byte, 1)
	for {
		n, err := r.r.Read(b)
		if err != nil {
			if err == io.EOF && buf.Len() > 0 {
				return "", errors.New("null byte not found before EOF")
			}
			return "", err
		}

		if n == 0 {
			continue
		}

		if b[0] == 0x00 {
			break
		}

		buf.WriteByte(b[0])
	}

	return buf.String(), nil
}
