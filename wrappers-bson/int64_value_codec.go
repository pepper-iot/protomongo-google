package wrappers_bson

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var Int64ValueType = reflect.TypeOf((*wrapperspb.Int64Value)(nil))

// Int64ValueCodec is the Codec used for *wrapperspb.Int64Value values.
type Int64ValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.Int64Value.
func (c *Int64ValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != Int64ValueType {
		return bsoncodec.ValueEncoderError{
			Name:     "Int64ValueCodec.EncodeValue",
			Types:    []reflect.Type{Int64ValueType},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.Int64Value)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteInt64(val.Value)
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.Int64Value.
func (c *Int64ValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != Int64ValueType {
		return bsoncodec.ValueDecoderError{
			Name:     "Int64ValueCodec.DecodeValue",
			Types:    []reflect.Type{Int64ValueType},
			Received: v,
		}
	}
	var val *wrapperspb.Int64Value
	switch bsonTyp := vr.Type(); bsonTyp {
	case bson.TypeInt64:
		v, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		val = wrapperspb.Int64(v)
	case bson.TypeInt32:
		v, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		val = wrapperspb.Int64(int64(v))
	case bson.TypeString:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		val = wrapperspb.Int64(v)
	case bson.TypeNull:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		val = nil
	case bson.TypeUndefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		val = &wrapperspb.Int64Value{}
	default:
		return fmt.Errorf("cannot decode %v into a *wrapperspb.Int64Value", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewInt64ValueCodec returns a Int64ValueCodec.
func NewInt64ValueCodec() *Int64ValueCodec {
	return &Int64ValueCodec{}
}
