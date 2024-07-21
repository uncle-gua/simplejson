package simplejson

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strconv"
)

var ErrInvalidType = errors.New("invalid value type")

// Implements the json.Unmarshaler interface.
func (j *Json) UnmarshalJSON(p []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(p))
	dec.UseNumber()
	return dec.Decode(&j.data)
}

// NewFromReader returns a *Json by decoding from an io.Reader
func NewFromReader(r io.Reader) (*Json, error) {
	j := new(Json)
	dec := json.NewDecoder(r)
	dec.UseNumber()
	err := dec.Decode(&j.data)
	return j, err
}

// Float64 coerces into a float64
func (j *Json) Float64() (float64, error) {
	switch v := j.data.(type) {
	case json.Number:
		return v.Float64()
	case float32, float64:
		return reflect.ValueOf(v).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(v).Uint()), nil
	case string:
		return strconv.ParseFloat(v, 64)
	}
	return 0, ErrInvalidType
}

// Int coerces into an int
func (j *Json) Int() (int, error) {
	switch v := j.data.(type) {
	case json.Number:
		i, err := v.Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(v).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(v).Uint()), nil
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		return int(i), err
	}
	return 0, ErrInvalidType
}

// Int64 coerces into an int64
func (j *Json) Int64() (int64, error) {
	switch v := j.data.(type) {
	case json.Number:
		return v.Int64()
	case float32, float64:
		return int64(reflect.ValueOf(v).Float()), nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(v).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(v).Uint()), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	}
	return 0, ErrInvalidType
}

// Uint64 coerces into an uint64
func (j *Json) Uint64() (uint64, error) {
	switch v := j.data.(type) {
	case json.Number:
		return strconv.ParseUint(v.String(), 10, 64)
	case float32, float64:
		return uint64(reflect.ValueOf(v).Float()), nil
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(v).Uint(), nil
	case string:
		return strconv.ParseUint(v, 10, 64)
	}
	return 0, ErrInvalidType
}
