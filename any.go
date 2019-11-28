package gapi

import (
	"bytes"
	"errors"

	"github.com/vmihailenco/msgpack/v4"
)

func NewAny(i interface{}) *Any {
	bytes, err := msgpack.Marshal(i)
	if err != nil {
		return nil
	}
	return &Any{
		Value: bytes,
	}
}

func (any *Any) Parse(i interface{}) error {
	if any == nil {
		return errors.New("value is null")
	}
	return msgpack.Unmarshal(any.Value, i)
}

func (any *Any) ParseString() (string, error) {
	s := string("")
	err := any.Parse(&s)
	return s, err
}

func (any *Any) ParseInt() (int, error) {
	i := int(0)
	err := any.Parse(&i)
	return i, err
}

func (any *Any) ParseBool() (bool, error) {
	i := false
	err := any.Parse(&i)
	return i, err
}

func (any *Any) ParseFloat32() (float32, error) {
	i := float32(0)
	err := any.Parse(&i)
	return i, err
}

func (any *Any) Equal(any2 *Any) bool {
	return bytes.Equal(any.Value, any2.Value)
}
