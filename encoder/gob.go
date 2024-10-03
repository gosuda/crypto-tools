package encoder

import (
	"encoding/gob"
	"io"
)

type GobEncoder struct{}

var _ Encoder = (*GobEncoder)(nil)

func (e GobEncoder) Encode(input io.Reader, output io.Writer) error {
	encoder := gob.NewEncoder(output)

	var data interface{}
	dec := gob.NewDecoder(input)
	if err := dec.Decode(&data); err != nil {
		return err
	}

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func (e GobEncoder) Decode(input io.Reader, output io.Writer) error {
	decoder := gob.NewDecoder(input)

	var data interface{}
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	encoder := gob.NewEncoder(output)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
