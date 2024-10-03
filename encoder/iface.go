package encoder

import "io"

// Encoder interface for stream-based encoding and decoding
type Encoder interface {
	Encode(input io.Reader, output io.Writer) error
	Decode(input io.Reader, output io.Writer) error
}
