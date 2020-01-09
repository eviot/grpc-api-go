package gapi

import (
	"bytes"
	"errors"
	"reflect"

	"github.com/vmihailenco/msgpack/v4"
)

// NewAny create Any with value, which maybe a bool, int, uint, float, string, array, slice and map
func NewAny(i interface{}) *Any {
	if i == nil {
		return &Any{Value: &Any_NilValue{true}}
	}
	switch val := i.(type) {
	case string:
		return &Any{Value: &Any_StrValue{val}}

	case bool:
		return &Any{Value: &Any_BoolValue{val}}

	case int:
		return &Any{Value: &Any_Int32Value{int32(val)}}

	case int8:
		return &Any{Value: &Any_Int32Value{int32(val)}}

	case int16:
		return &Any{Value: &Any_Int32Value{int32(val)}}

	case int32:
		return &Any{Value: &Any_Int32Value{int32(val)}}

	case int64:
		return &Any{Value: &Any_Int64Value{val}}

	case uint:
		return &Any{Value: &Any_Uint32Value{uint32(val)}}

	case uint8:
		return &Any{Value: &Any_Uint32Value{uint32(val)}}

	case uint16:
		return &Any{Value: &Any_Uint32Value{uint32(val)}}

	case uint32:
		return &Any{Value: &Any_Uint32Value{uint32(val)}}

	case uint64:
		return &Any{Value: &Any_Uint64Value{val}}

	case float32:
		return &Any{Value: &Any_FloatValue{val}}

	case float64:
		return &Any{Value: &Any_DoubleValue{val}}

	case []byte:
		return &Any{Value: &Any_BinaryValue{val}}
	}

	val, err := msgpack.Marshal(i)
	if err != nil {
		return nil
	}
	return &Any{
		Value: &Any_MsgpackValue{val},
	}
}

// Parse return value
func (any *Any) Parse() (interface{}, error) {
	if any == nil {
		return nil, errors.New("any is nil")
	}
	switch val := any.Value.(type) {
	case *Any_BoolValue:
		return val.BoolValue, nil

	case *Any_Int32Value:
		return val.Int32Value, nil

	case *Any_Int64Value:
		return val.Int64Value, nil

	case *Any_Uint32Value:
		return val.Uint32Value, nil

	case *Any_Uint64Value:
		return val.Uint64Value, nil

	case *Any_FloatValue:
		return val.FloatValue, nil

	case *Any_DoubleValue:
		return val.DoubleValue, nil

	case *Any_StrValue:
		return val.StrValue, nil

	case *Any_BinaryValue:
		return val.BinaryValue, nil

	case *Any_MsgpackValue:
		dec := msgpack.NewDecoder(bytes.NewBuffer(val.MsgpackValue))
		i, err := dec.DecodeInterface()
		return i, err
	}
	return nil, errors.New("undefined type")
}

// ParseTo parse any to value
func (any *Any) ParseTo(value interface{}) error {
	if any == nil {
		return errors.New("any is nil")
	}
	kind := reflect.TypeOf(value).Kind()
	if kind != reflect.Ptr || kind != reflect.Map {
		return errors.New("value must be pointer or map")
	}
	elemPtr := reflect.ValueOf(value)
	elem := elemPtr.Elem()
	elemKind := elem.Kind()
	switch anyVal := any.Value.(type) {
	case *Any_BoolValue:
		if elemKind != reflect.Bool {
			return errors.New("value type must be bool")
		}
		elem.SetBool(anyVal.BoolValue)
		return nil

	case *Any_Int32Value, *Any_Int64Value:
		if elemKind != reflect.Int &&
			elemKind != reflect.Int8 &&
			elemKind != reflect.Int16 &&
			elemKind != reflect.Int32 &&
			elemKind != reflect.Int64 {
			return errors.New("value type must be int")
		}
		if i32, ok := anyVal.(*Any_Int32Value); ok {
			elem.SetInt(int64(i32.Int32Value))
		} else {
			i64, _ := anyVal.(*Any_Int64Value)
			elem.SetInt(i64.Int64Value)
		}
		return nil

	case *Any_Uint32Value, *Any_Uint64Value:
		if elemKind != reflect.Uint &&
			elemKind != reflect.Uint8 &&
			elemKind != reflect.Uint16 &&
			elemKind != reflect.Uint32 &&
			elemKind != reflect.Uint64 {
			return errors.New("value type must be uint")
		}
		if i32, ok := anyVal.(*Any_Uint32Value); ok {
			elem.SetUint(uint64(i32.Uint32Value))
		} else {
			i64, _ := anyVal.(*Any_Uint64Value)
			elem.SetUint(i64.Uint64Value)
		}
		return nil

	case *Any_FloatValue, *Any_DoubleValue:
		if elemKind != reflect.Float32 &&
			elemKind != reflect.Float64 {
			return errors.New("value type must be float")
		}
		if i32, ok := anyVal.(*Any_FloatValue); ok {
			elem.SetFloat(float64(i32.FloatValue))
		} else {
			i64, _ := anyVal.(*Any_DoubleValue)
			elem.SetFloat(i64.DoubleValue)
		}
		return nil

	case *Any_StrValue:
		if elemKind != reflect.String {
			return errors.New("value type must be string")
		}
		elem.SetString(anyVal.StrValue)
		return nil

	case *Any_BinaryValue:
		if elemKind == reflect.Slice &&
			elem.Type().Elem().Kind() != reflect.Uint8 {
			return errors.New("value type must be []byte")
		}
		elem.SetBytes(anyVal.BinaryValue)
		return nil

	case *Any_MsgpackValue:
		err := msgpack.Unmarshal(anyVal.MsgpackValue, value)
		return err
	}
	return errors.New("undefined type")
}

// IsString returns true if it is a string
func (any *Any) IsString() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_StrValue)
	return ok
}

// IsBool returns true if it is a bool
func (any *Any) IsBool() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_BoolValue)
	return ok
}

// IsInt returns true if it is a int
func (any *Any) IsInt() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_Int32Value)
	if !ok {
		_, ok = any.Value.(*Any_Int64Value)
	}
	return ok
}

// IsInt32 returns true if it is a int32
func (any *Any) IsInt32() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_Int32Value)
	return ok
}

// IsInt64 returns true if it is a int64
func (any *Any) IsInt64() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_Int64Value)
	return ok
}

// IsUint returns true if it is a uint
func (any *Any) IsUint() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_Uint32Value)
	if !ok {
		_, ok = any.Value.(*Any_Uint64Value)
	}
	return ok
}

// IsUint32 returns true if it is a uint32
func (any *Any) IsUint32() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_Uint32Value)
	return ok
}

// IsUint64 returns true if it is a uint64
func (any *Any) IsUint64() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_Uint64Value)
	return ok
}

// IsFloat returns true if it is a float
func (any *Any) IsFloat() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_FloatValue)
	if !ok {
		_, ok = any.Value.(*Any_DoubleValue)
	}
	return ok
}

// IsFloat32 returns true if it is a float32
func (any *Any) IsFloat32() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_FloatValue)
	return ok
}

// IsFloat64 returns true if it is a float64
func (any *Any) IsFloat64() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_DoubleValue)
	return ok
}

// IsBinary returns true if it is a []byte
func (any *Any) IsBinary() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_BinaryValue)
	return ok
}

// IsMsgpack returns true if it is a array, slice, map or other value
func (any *Any) IsMsgpack() bool {
	if any == nil {
		return false
	}
	_, ok := any.Value.(*Any_MsgpackValue)
	return ok
}

func (any *Any) ToString() string {
	if any == nil {
		return ""
	}
	v, ok := any.Value.(*Any_StrValue)
	if !ok {
		return ""
	}
	return v.StrValue
}

func (any *Any) ToBoolean() bool {
	if any == nil {
		return false
	}
	v, ok := any.Value.(*Any_BoolValue)
	if !ok {
		return false
	}
	return v.BoolValue
}

func (any *Any) ToInt32() int32 {
	if any == nil {
		return 0
	}
	v, ok := any.Value.(*Any_Int32Value)
	if !ok {
		return 0
	}
	return v.Int32Value
}

func (any *Any) ToInt64() int64 {
	if any == nil {
		return 0
	}
	v, ok := any.Value.(*Any_Int64Value)
	if !ok {
		return 0
	}
	return v.Int64Value
}

func (any *Any) ToUint32() uint32 {
	if any == nil {
		return 0
	}
	v, ok := any.Value.(*Any_Uint32Value)
	if !ok {
		return 0
	}
	return v.Uint32Value
}

func (any *Any) ToUint64() uint64 {
	if any == nil {
		return 0
	}
	v, ok := any.Value.(*Any_Uint64Value)
	if !ok {
		return 0
	}
	return v.Uint64Value
}

func (any *Any) ToFloat32() float32 {
	if any == nil {
		return 0
	}
	v, ok := any.Value.(*Any_FloatValue)
	if !ok {
		return 0
	}
	return v.FloatValue
}

func (any *Any) ToFloat64() float64 {
	if any == nil {
		return 0
	}
	v, ok := any.Value.(*Any_DoubleValue)
	if !ok {
		return 0
	}
	return v.DoubleValue
}

func (any *Any) ToBinary() []byte {
	if any == nil {
		return []byte{}
	}
	v, ok := any.Value.(*Any_BinaryValue)
	if !ok {
		return []byte{}
	}
	return v.BinaryValue
}

// Equal returns true if any equal to any2
func (any *Any) Equal(any2 *Any) bool {
	if any == nil || any2 == nil {
		return any == any2
	}
	return any.Value == any2.Value
}

func AnyMap(m map[string]interface{}) map[string]*Any {
	res := map[string]*Any{}
	for k, v := range m {
		res[k] = NewAny(v)
	}
	return res
}

func InterfaceMap(m map[string]*Any) map[string]interface{} {
	res := map[string]interface{}{}
	for k, v := range m {
		vv, err := v.Parse()
		if err != nil {
			continue
		}
		res[k] = vv
	}
	return res
}
