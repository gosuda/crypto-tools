package encoder

import (
	"encoding/hex"
	"io"
)

type HexEncoder struct{}

var _ Encoder = (*HexEncoder)(nil)

func (e HexEncoder) Encode(input io.Reader, output io.Writer) error {
	data, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	encoded := hex.EncodeToString(data)

	_, err = output.Write([]byte(encoded))
	return err
}

func (e HexEncoder) Decode(input io.Reader, output io.Writer) error {
	data, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	decoded, err := hex.DecodeString(string(data))
	if err != nil {
		return err
	}

	_, err = output.Write(decoded)
	return err
}
