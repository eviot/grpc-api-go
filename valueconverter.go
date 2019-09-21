package gapi

import (
	"bytes"
	"encoding/binary"
	"math"
	"reflect"
)

func NewValue(i interface{}) *Value {
	val := &Value{}
	kind := reflect.TypeOf(i).Kind()
	iVal := reflect.ValueOf(i)
	switch kind {
	case reflect.String:
		val.SetString(iVal.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val.SetUint64(uint64(iVal.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val.SetUint64(iVal.Uint())
	case reflect.Float32, reflect.Float64:
		val.SetFloat64(iVal.Float())
	case reflect.Bool:
		val.SetBool(iVal.Bool())
	default:
		return nil
	}
	return val
}

// SetString sets string in value
func (val *Value) SetString(s string) {
	val.Type = Value_STRING
	val.Value = []byte(s)
}

// ToString convert Value to string
func (val *Value) ToString() string {
	if val.Type != Value_STRING {
		return ""
	}
	s := string(val.Value)
	return s
}

// SetUint64 convert uint64 to Value protobuf message
func (val *Value) SetUint64(i uint64) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	val.Type = Value_UINT
	val.Value = b
}

// ToUint64 convert Value to uint64
func (val *Value) ToUint64() uint64 {
	if val.Type != Value_UINT {
		return 0
	}
	i := binary.LittleEndian.Uint64(val.Value)
	return i
}

// ToInt convert Value to int
func (val *Value) ToInt() int {
	return int(val.ToUint64())
}

// ToInt convert Value to int32
func (val *Value) ToInt32() int32 {
	return int32(val.ToUint64())
}

// SetFloat64 convert float64 to Value protobuf message
func (val *Value) SetFloat64(f float64) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, math.Float64bits(f))
	val.Type = Value_FLOAT
	val.Value = b
}

// ToFloat64 convert Value to float64
func (val *Value) ToFloat64() float64 {
	if val.Type != Value_FLOAT {
		return 0
	}
	f := binary.LittleEndian.Uint64(val.Value)
	return math.Float64frombits(f)
}

// SetBool convert bool to Value protobuf message
func (val *Value) SetBool(b bool) {
	val.Type = Value_BOOLEAN
	if b {
		val.Value = []byte{1}
	} else {
		val.Value = []byte{0}
	}
}

// ToBool convert Value to bool
func (val *Value) ToBool() bool {
	if val.Type != Value_BOOLEAN {
		return false
	}
	if len(val.Value) >= 1 && val.Value[0] >= 1 {
		return true
	}
	return false
}

func (val *Value) Equal(val2 *Value) bool {
	if val.Type != val2.Type {
		return false
	}
	if !bytes.Equal(val.Value, val2.Value) {
		return false
	}
	return true
}
