package encoder

import (
	"encoding/base64"
	"io"
)

type Base64Encoder struct{}

var _ Encoder = (*Base58Encoder)(nil)

func (e Base64Encoder) Encode(input io.Reader, output io.Writer) error {
	encoder := base64.NewEncoder(base64.StdEncoding, output)
	defer encoder.Close()

	_, err := io.Copy(encoder, input)
	return err
}

func (e Base64Encoder) Decode(input io.Reader, output io.Writer) error {
	decoder := base64.NewDecoder(base64.StdEncoding, input)

	_, err := io.Copy(output, decoder)
	return err
}
