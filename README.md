
# Binary Reader

`go-binary` is a Go package that provides a convenient way to read various types of binary data from an `io.Reader`. It supports reading primitive data types, including boolean, integers, floats, and null-terminated strings with specified byte order.

## Installation

To install the package, use the following command:

```sh
go get github.com/zelenin/go-binary
```

## Usage

Here's a quick example of how to use the `go-binary` package:

```go
package main

import (
	"bytes"
	"encoding/binary"
	gobinary "github.com/zelenin/go-binary"

	"fmt"
	"log"
)

func main() {
	data := []byte{0x01, 0x00, 0x80, 0x00, 0xff, 0x00, 0x00, 0x00, 0x80}
	reader := gobinary.NewReader(bytes.NewReader(data), binary.BigEndian)

	boolVal, err := reader.ReadBool()
	if err != nil {
		log.Fatalf("ReadBool error: %v", err)
	}
	fmt.Println("ReadBool:", boolVal)

	int8Val, err := reader.ReadInt8()
	if err != nil {
		log.Fatalf("ReadInt8 error: %v", err)
	}
	fmt.Println("ReadInt8:", int8Val)

	uint8Val, err := reader.ReadUint8()
	if err != nil {
		log.Fatalf("ReadUint8 error: %v", err)
	}
	fmt.Println("ReadUint8:", uint8Val)

	int16Val, err := reader.ReadInt16()
	if err != nil {
		log.Fatalf("ReadInt16 error: %v", err)
	}
	fmt.Println("ReadInt16:", int16Val)

	int32Val, err := reader.ReadInt32()
	if err != nil {
		log.Fatalf("ReadInt32 error: %v", err)
	}
	fmt.Println("ReadInt32:", int32Val)
}
```

## API

### `NewReader(r io.Reader, byteOrder binary.ByteOrder) *Reader`
### `NewReaderFromBytes(data []byte, byteOrder binary.ByteOrder) *Reader`

Creates a new `Reader` instance.

### `ReadBool() (bool, error)`

Reads a single byte and returns `true` if it's non-zero, otherwise returns `false`.

### `ReadInt8() (int8, error)`

Reads a single byte and returns it as `int8`.

### `ReadUint8() (uint8, error)`

Reads a single byte and returns it as `uint8`.

### `ReadInt16(byteOrder binary.ByteOrder) (int16, error)`

Reads 2 bytes and returns them as `int16` with the specified byte order.

### `ReadUint16(byteOrder binary.ByteOrder) (uint16, error)`

Reads 2 bytes and returns them as `uint16` with the specified byte order.

### `ReadInt32(byteOrder binary.ByteOrder) (int32, error)`

Reads 4 bytes and returns them as `int32` with the specified byte order.

### `ReadUint32(byteOrder binary.ByteOrder) (uint32, error)`

Reads 4 bytes and returns them as `uint32` with the specified byte order.

### `ReadInt64(byteOrder binary.ByteOrder) (int64, error)`

Reads 8 bytes and returns them as `int64` with the specified byte order.

### `ReadUint64(byteOrder binary.ByteOrder) (uint64, error)`

Reads 8 bytes and returns them as `uint64` with the specified byte order.

### `ReadFloat32(byteOrder binary.ByteOrder) (float32, error)`

Reads 4 bytes and returns them as `float32` with the specified byte order.

### `ReadFloat64(byteOrder binary.ByteOrder) (float64, error)`

Reads 8 bytes and returns them as `float64` with the specified byte order.

### `ReadNullTerminatedString() (string, error)`

Reads bytes until a null byte (`0x00`) is encountered and returns the resulting string.

## Tests

To run the tests for this package, use the following command:

```sh
go test ./...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
