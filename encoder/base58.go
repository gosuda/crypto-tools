package encoder

import (
	"io"

	"github.com/btcsuite/btcutil/base58"
)

type Base58Encoder struct{}

var _ Encoder = (*Base58Encoder)(nil)

func (e Base58Encoder) Encode(input io.Reader, output io.Writer) error {
	data, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	encoded := base58.Encode(data)

	_, err = output.Write([]byte(encoded))
	return err
}

func (e Base58Encoder) Decode(input io.Reader, output io.Writer) error {
	data, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	decoded := base58.Decode(string(data))

	_, err = output.Write(decoded)
	return err
}
